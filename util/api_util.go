package util

import (
	"net/http"

	"github.com/achintya-7/go-template-server/internal/dto"

	"github.com/gin-gonic/gin"
)

// HandlerFunction is the type of handler function callback.
type HandlerFunction[T any] func(*gin.Context) (*T, *dto.ErrorResponse)

// HandlerWrapper is a wrapper around all the handler methods.
// Only this method will handle the response creation (i.e., ctx.JSON).
// All the handler methods in this wrapper should return a pair of values:
// 1. Result computed (which can be of any type)
// 2. A pointer to ErrorResponse
func HandlerWrapper[T any](callback HandlerFunction[T]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer handlePanic(ctx)

		result, err := callback(ctx)
		if err != nil {
			sendErrorResponse(ctx, err)
			return
		}

		sendSuccessResponse(ctx, result)
	}
}

func handlePanic(ctx *gin.Context) {
	if r := recover(); r != nil {
		// logger to send error
		sendErrorResponse(ctx, &dto.ErrorResponse{
			Code:           http.StatusInternalServerError,
			Message:        "Internal Server Error",
			HttpStatusCode: http.StatusInternalServerError,
		})
	}
}

func sendErrorResponse(ctx *gin.Context, err *dto.ErrorResponse) {
	ctx.AbortWithStatusJSON(err.HttpStatusCode, dto.ApiResponse{
		Status: false,
		Data:   nil,
		Error: &dto.ApiError{
			Code:           err.Code,
			Message:        err.Message,
			HttpStatusCode: err.HttpStatusCode,
		},
	})
}

func sendSuccessResponse[T any](ctx *gin.Context, result *T) {
	if result == nil {
		sendNotFoundResponse(ctx)
		return
	}
	ctx.JSON(http.StatusOK, dto.ApiResponse{
		Status: true,
		Data:   result,
		Error:  nil,
	})
}

func sendNotFoundResponse(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ApiResponse{
		Status: false,
		Data:   nil,
		Error: &dto.ApiError{
			Code:           http.StatusNotFound,
			Message:        "Oops! No data found",
			HttpStatusCode: http.StatusNotFound,
		},
	})
}
