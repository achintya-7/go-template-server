package util

import (
	"encoding/json"
	"net/http"

	"github.com/achintya-7/go-template-server/internal/dto"
)

type HttpHandlerFunction[T any] func(*http.Request) (*T, *dto.ErrorResponse)

func HttpHandlerWrapper[T any](callback HttpHandlerFunction[T]) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer handlePanicV2(w)

		result, err := callback(r)
		if err != nil {
			sendErrorResponseV2(w, err)
			return
		}
		sendSuccessResponseV2(w, result)
	}
}

func handlePanicV2(w http.ResponseWriter) {
	if r := recover(); r != nil {
		sendErrorResponseV2(w, &dto.ErrorResponse{
			Message:        "Internal Server Error",
			HttpStatusCode: http.StatusInternalServerError,
		})
	}
}

func sendErrorResponseV2(w http.ResponseWriter, err *dto.ErrorResponse) {
	w.WriteHeader(err.HttpStatusCode)
	w.Write([]byte(err.Message))
}

func sendSuccessResponseV2[T any](w http.ResponseWriter, result *T) {
	if result == nil {
		sendNotFoundResponseV2(w)
		return
	}

	data := dto.ApiResponse{
		Status: true,
		Data:   result,
		Error:  nil,
	}

	resultBytes, err := json.Marshal(data)
	if err != nil {
		sendErrorResponseV2(w, &dto.ErrorResponse{
			Message:        "Internal Server Error",
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resultBytes))
}

func sendNotFoundResponseV2(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}
