package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	r := gin.Default()

	return &Server{
		router: r,
	}
}

func (s *Server) Run(addr string) {
	s.router.Run(addr)
}
