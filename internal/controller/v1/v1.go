package v1

import (
	"net/http"

	v1 "github.com/achintya-7/go-template-server/internal/handlers/v1"
	"github.com/achintya-7/go-template-server/util"
)

type Router struct {
	handlers *v1.RouteHandler
}

func NewRouter() *Router {
	return &Router{
		handlers: v1.NewRouteHandler(),
	}
}

func (r *Router) SetupRoutes(router *http.ServeMux) {
	v1Router := http.NewServeMux()

	v1Router.HandleFunc("/item/{id}", util.HttpHandlerWrapper(r.handlers.ItemHandler))

	router.Handle("/v1/", http.StripPrefix("/v1", v1Router))
}
