package middleware

import (
	"context"
	"go-rest-api/logger"
	"net/http"

	"github.com/google/uuid"
)
// Context key for the logger
type contextKey string

const LoggerKey contextKey = "logger"
// LoggerMiddleware injects request ID and user ID into the context
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := uuid.New().String()
		ctx = context.WithValue(ctx, logger.RequestIDKey, requestID)


		// logger.Log().Infof("Received %s request for %s", r.Method, r.URL.Path)

		// // Pass updated context to next handler
		// next.ServeHTTP(w, r.WithContext(ctx))

		logEntry := logger.WithContext(ctx)
		ctx = context.WithValue(ctx, LoggerKey, logEntry)

		// Log incoming request
		logEntry.Infof("Incoming request: %s %s", r.Method, r.URL.Path)

		// Pass updated context to next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
