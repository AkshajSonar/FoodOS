package server

import (


	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"foodos-backend/internal/routes"

)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
	AllowAllOrigins: true,
	AllowMethods: []string{
		"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS",
	},
	AllowHeaders: []string{
		"Content-Type", "Authorization",
	},
}))


	routes.RegisterRoutes(r)

	return &Server{
		router: r,
	}
}

func (s *Server) Run(addr string) {
	s.router.Run(addr)
}
