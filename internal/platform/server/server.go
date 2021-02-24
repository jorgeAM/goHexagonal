package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jorgeAM/goHexagonal/internal/platform/server/handler/courses"
	"github.com/jorgeAM/goHexagonal/internal/platform/server/handler/ping"

	mooc "github.com/jorgeAM/goHexagonal/internal/platform"
)

type Server struct {
	httpAdrr         string
	engine           *gin.Engine
	courseRepository mooc.CourseRepository
}

func NewServer(host string, port uint, courseRepository mooc.CourseRepository) *Server {
	s := &Server{
		httpAdrr:         fmt.Sprintf("%v:%d", host, port),
		engine:           gin.New(),
		courseRepository: courseRepository,
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
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}
