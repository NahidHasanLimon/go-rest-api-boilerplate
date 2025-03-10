package middleware

import (
	"context"
	"net/http"
	"github.com/google/uuid"
	"go-rest-api/config"
)

// LoggerMiddleware injects request ID and user ID into the context
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Generate request ID (always unique)
		requestID := uuid.New().String()
		ctx = context.WithValue(ctx, "request_id", requestID)
		// Extract user ID (if authenticated)
		userID := r.Header.Get("X-User-ID") // Assuming user ID is passed in header
		if userID != "" {
			ctx = context.WithValue(ctx, "user_id", userID)
		}
		config.Log().Infof("Received %s request for %s", r.Method, r.URL.Path)


		// Pass updated context to next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}