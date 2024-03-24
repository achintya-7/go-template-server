package util

import (
	"net/http"

	"github.com/achintya-7/go-template-server/internal/dto"

	"github.com/gin-gonic/gin"
)

type HandlerFunction[T any] func(*gin.Context) (*T, *dto.ErrorResponse)

func HandleWrapper[T any](callback HandlerFunction[T]) func(ctx *gin.Context) {
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
			Message:        "Oops! No data found",
			HttpStatusCode: http.StatusNotFound,
		},
	})
}
