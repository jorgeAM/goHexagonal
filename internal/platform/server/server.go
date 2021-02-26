package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jorgeAM/goHexagonal/internal/kit/command"
	"github.com/jorgeAM/goHexagonal/internal/platform/server/handler/courses"
	"github.com/jorgeAM/goHexagonal/internal/platform/server/handler/ping"
)

type Server struct {
	httpAdrr   string
	engine     *gin.Engine
	commandBus command.Bus
}

func NewServer(host string, port uint, commandBus command.Bus) *Server {
	s := &Server{
		httpAdrr:   fmt.Sprintf("%v:%d", host, port),
		engine:     gin.New(),
		commandBus: commandBus,
	}

	s.registerRoutes()

	return s
}

func (s *Server) Run() error {
	fmt.Println("Server is running on: ", s.httpAdrr)
	return s.engine.Run(s.httpAdrr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/ping", ping.PingHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
}
