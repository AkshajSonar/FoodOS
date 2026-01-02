package server

import (
	"foodos-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	r := gin.Default()

	// register all routes
	routes.RegisterRoutes(r)

	return &Server{
		router: r,
	}
}

func (s *Server) Run(addr string) {
	s.router.Run(addr)
}
