package v1

import (
	"github.com/achintya-7/go-template-server/internal/dto"
	"github.com/gin-gonic/gin"
)

type RouterInterface interface {
	PrivateHello(context *gin.Context) (*gin.H, *dto.ErrorResponse)
	PublicHello(context *gin.Context) (*gin.H, *dto.ErrorResponse)
}

type RouteHandler struct {
}

func NewRouteHandler() RouterInterface {
	return &RouteHandler{}
}

func (r *RouteHandler) PrivateHello(context *gin.Context) (*gin.H, *dto.ErrorResponse) {
	return &gin.H{"message": "Hello World from a private API"}, nil
}

func (r *RouteHandler) PublicHello(context *gin.Context) (*gin.H, *dto.ErrorResponse) {
	return &gin.H{"message": "Hello World from a public API"}, nil
}
