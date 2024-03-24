package v1

import (
	v1 "github.com/achintya-7/go-template-server/internal/handlers/v1"
	"github.com/achintya-7/go-template-server/pkg/util"
	"github.com/gin-gonic/gin"
)

type RouterInterface interface {
	SetupRoutes(route *gin.RouterGroup)
}

type Router struct {
	handlers v1.RouterInterface
}

func NewRouter(route *gin.RouterGroup) RouterInterface {

	router := &Router{
		handlers: v1.NewRouteHandler(),
	}

	router.SetupRoutes(route)

	return router

}

func (r *Router) SetupRoutes(route *gin.RouterGroup) {
	v1GroupPublic := route.Group("/v1")
	v1GroupPrivate := route.Group("/v1")

	// new private routes here
	v1GroupPrivate.GET("private_hello", util.HandleWrapper(r.handlers.PrivateHello))

	// new public routes here
	v1GroupPublic.GET("public_hello", util.HandleWrapper(r.handlers.PublicHello))
}
