package v1

import (
	"net/http"

	v1 "github.com/achintya-7/go-template-server/internal/handlers/v1"
	"github.com/achintya-7/go-template-server/internal/middlewares"
	"github.com/achintya-7/go-template-server/util"
)

type Router struct {
	mux      *http.ServeMux
	handlers *v1.RouteHandler
}

func NewRouter(mux *http.ServeMux) *Router {
	return &Router{
		handlers: v1.NewRouteHandler(),
		mux:      mux,
	}
}

func (r *Router) SetupRoutes() middlewares.Middleware {

	// register middleware
	stack := middlewares.CreateMiddlewareStack(
		middlewares.LoggingMiddleware,
		middlewares.CorrelationMiddleware,
	)

	// register routes
	r.mux.HandleFunc("/item/{id}", util.HttpHandlerWrapper(r.handlers.ItemHandler))

	return stack
}
