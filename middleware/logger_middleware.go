package middleware

import (
	"context"
	"github.com/google/uuid"
	"go-rest-api/logger"
	"net/http"
)

// LoggerMiddleware injects request ID and user ID into the context
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := uuid.New().String()
		ctx = context.WithValue(ctx, logger.RequestIDKey, requestID)

		logEntry := logger.WithContext(ctx)
		logEntry.Infof("Incoming request: %s %s", r.Method, r.URL.Path)
	

		// Pass updated context to next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
