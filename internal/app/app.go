package app

import (
	"net/http"

	v1 "github.com/achintya-7/go-template-server/internal/controller/v1"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	server := &Server{}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {

	router := http.NewServeMux()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	v1Router := v1.NewRouter()
	v1Router.SetupRoutes(router)

	s.router = router

}

func (s *Server) Start(port string) error {
	server := http.Server{
		Addr:    port,
		Handler: s.router,
	}

	return server.ListenAndServe()
}
