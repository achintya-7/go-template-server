package v1

import (
	"net/http"

	v2 "github.com/achintya-7/go-template-server/internal/handlers/v2"
	"github.com/achintya-7/go-template-server/internal/middlewares"
	"github.com/achintya-7/go-template-server/util"
)

type Router struct {
	mux      *http.ServeMux
	handlers *v2.RouteHandler
}

func NewRouter(mux *http.ServeMux) *Router {
	return &Router{
		handlers: v2.NewRouteHandler(),
		mux:      mux,
	}
}

func (r *Router) SetupRoutes() middlewares.Middleware {
	// register routes
	r.mux.HandleFunc("/item/{id}", util.HttpHandlerWrapper(r.handlers.ItemHandler))

	return nil
}
