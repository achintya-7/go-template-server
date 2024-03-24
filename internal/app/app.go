package app

import (
	"net/http"

	v1 "github.com/achintya-7/go-template-server/internal/controller/v1"
	v2 "github.com/achintya-7/go-template-server/internal/controller/v2"
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

	// v1 routes
	v1ServeMux := http.NewServeMux()
	v1Router := v1.NewRouter(v1ServeMux)
	middlewareStack := v1Router.SetupRoutes()
	if middlewareStack != nil {
		router.Handle("/v1/", middlewareStack(http.StripPrefix("/v1", v1ServeMux)))
	} else {
		router.Handle("/v1/", http.StripPrefix("/v1", v1ServeMux))
	}

	// v2 routes
	v2ServeMux := http.NewServeMux()
	v2Router := v2.NewRouter(v2ServeMux)
	middlewareStack = v2Router.SetupRoutes()
	if middlewareStack != nil {
		router.Handle("/v2/", http.StripPrefix("/v2", middlewareStack(v2ServeMux)))
	} else {
		router.Handle("/v2/", http.StripPrefix("/v2", v2ServeMux))
	}

	s.router = router

}

func (s *Server) Start(port string) error {
	server := http.Server{
		Addr:    port,
		Handler: s.router,
	}

	return server.ListenAndServe()
}
