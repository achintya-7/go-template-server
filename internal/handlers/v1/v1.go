package v1

import (
	"github.com/achintya-7/go-template-server/internal/dto"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
}

func NewRouteHandler() *RouteHandler {
	return &RouteHandler{}
}

func (*RouteHandler) PrivateHello(context *gin.Context) (*gin.H, *dto.ErrorResponse) {
	return &gin.H{"message": "Hello World from a private API"}, nil
}

func (*RouteHandler) PublicHello(context *gin.Context) (*gin.H, *dto.ErrorResponse) {
	return &gin.H{"message": "Hello World from a public API"}, nil
}
