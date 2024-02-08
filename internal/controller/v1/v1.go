package v1

import (
	v1 "github.com/achintya-7/go-template-server/internal/handlers/v1"
	"github.com/achintya-7/go-template-server/util"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handlers *v1.RouteHandler
}

func NewRouter() *Router {
	return &Router{
		handlers: v1.NewRouteHandler(),
	}
}

func (r *Router) SetupRoutes(route *gin.RouterGroup) {
	v1GroupPublic := route.Group("/v1")
	v1GroupPrivate := route.Group("/v1")

	// new private routes here
	v1GroupPrivate.GET("private_hello", util.HandlerWrapper(r.handlers.PrivateHello))

	// new public routes here
	v1GroupPublic.GET("public_hello", util.HandlerWrapper(r.handlers.PublicHello))
}
