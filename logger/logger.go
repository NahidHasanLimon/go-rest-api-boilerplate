package logger
import (
	"os"
	"context"
	"github.com/sirupsen/logrus"
	"time"
)
var logger *logrus.Logger

type contextKey string

const RequestIDKey contextKey = "request_id"
func init(){
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout) 
	logger = logrus.New()

	// Custom JSON Formatter
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339, // ISO 8601 format
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
	})

	file,err := os.OpenFile("logs.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil { 
		logger.SetOutput(file)
	}else {
		logger.Warn("Failed to log to file, using default stderr")
	}
}
func Log() *logrus.Logger {
	return logger
}

func WithContext(ctx context.Context) *logrus.Entry {
	fields := logrus.Fields{}

	// Extract request ID from context
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		fields["request_id"] = reqID
	}

	return logger.WithFields(fields)
}

