package v2

import (
	"net/http"

	"github.com/achintya-7/go-template-server/internal/dto"
)

type RouteHandler struct {
}

func NewRouteHandler() *RouteHandler {
	return &RouteHandler{}
}

func (rh *RouteHandler) ItemHandler(r *http.Request) (*map[string]any, *dto.ErrorResponse) {
	id := r.PathValue("id")
	return &map[string]any{"message": "Item ID: " + id}, nil
}
