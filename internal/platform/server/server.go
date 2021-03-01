package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jorgeAM/goHexagonal/internal/kit/command"
	"github.com/jorgeAM/goHexagonal/internal/platform/server/handler/courses"
	"github.com/jorgeAM/goHexagonal/internal/platform/server/handler/ping"
	"github.com/jorgeAM/goHexagonal/internal/platform/server/middleware/logging"
	"github.com/jorgeAM/goHexagonal/internal/platform/server/middleware/recovery"
)

type Server struct {
	httpAdrr        string
	engine          *gin.Engine
	shutdownTimeout time.Duration
	commandBus      command.Bus
}

func NewServer(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus) (context.Context, *Server) {
	s := &Server{
		httpAdrr:        fmt.Sprintf("%v:%d", host, port),
		engine:          gin.New(),
		shutdownTimeout: shutdownTimeout,
		commandBus:      commandBus,
	}

	s.registerRoutes()

	return serverContext(ctx), s
}

func (s *Server) Run(ctx context.Context) error {
	fmt.Println("Server is running on: ", s.httpAdrr)

	srv := &http.Server{
		Addr:    s.httpAdrr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(ctx, s.shutdownTimeout)

	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func (s *Server) registerRoutes() {
	s.engine.Use(logging.Middlewares(), recovery.Middleware())

	s.engine.GET("/ping", ping.PingHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
}

func serverContext(ctx context.Context) context.Context {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-ch
		cancel()
	}()

	return ctx
}
