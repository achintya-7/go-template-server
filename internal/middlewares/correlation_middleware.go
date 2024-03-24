package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const CorrelationIDHeader = "X-Correlation-ID"

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get(CorrelationIDHeader)
		if correlationID == "" {
			log.Println("Correlation ID not found in request headers")
			correlationID = uuid.NewString()
		}

		w.Header().Set(CorrelationIDHeader, correlationID)

		ctx := context.WithValue(r.Context(), CorrelationIDHeader, correlationID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
