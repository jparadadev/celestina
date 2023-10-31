package server

import (
	"celestina/internal/platform/server/handler/forward"
	"celestina/internal/platform/server/handler/health"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr      string
	engine        *gin.Engine
	subscriptions map[string][]string
}

func New(host string, port int, subscriptions map[string][]string) Server {
	srv := Server{
		engine:        gin.New(),
		httpAddr:      fmt.Sprintf("%s:%d", host, port),
		subscriptions: subscriptions,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())

	forwardCtr := forward.New(s.subscriptions)
	s.engine.POST("/forward/:eventid", forwardCtr.PostHandler())
}
