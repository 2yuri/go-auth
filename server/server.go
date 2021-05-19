package server

import (
	"go-auth/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	host   string
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		host:   "localhost",
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	routes := routes.ConfigRoutes(s.server)
	routes.Run(":" + s.port)
}
