package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jpastorm/hexagonalgoexample-cdtv/internal/platform/server/handler/course"
	"github.com/jpastorm/hexagonalgoexample-cdtv/internal/platform/server/handler/health"
	"github.com/jpastorm/hexagonalgoexample-cdtv/kit/command"
	"log"
)

type Server struct {
	httpAddr string
	engine *gin.Engine

	commandBus command.Bus
}

func New(host string, port uint, commandBus command.Bus) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine : gin.New(),
		commandBus: commandBus,
	}

	srv.registerRoutes()
	return srv
}

func (s Server) Run() error {
	log.Println("Server running", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s Server) registerRoutes()  {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("courses", course.CreateHandler(s.commandBus))
}