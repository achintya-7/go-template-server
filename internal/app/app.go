package app

import (
	v1 "github.com/achintya-7/go-template-server/internal/controller/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ServerInterface interface {
	run(port string) error
	setupRoutes()
}

type Server struct {
	router *gin.Engine
}

func NewServer() ServerInterface {
	server := &Server{}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	router := gin.Default()

	baseRouter := router.Group("/service-name")

	// register all v1 routes
	v1Router := v1.NewRouter(baseRouter)
	v1Router.SetupRoutes(baseRouter)

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
	}))

	s.router = router
}

func (s *Server) run(port string) error {
	return s.router.Run(port)
}
