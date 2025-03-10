package logger
import (
	"os"
	"github.com/sirupsen/logrus"
)
var logger *logrus.Logger

func init(){
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout) 
	logger = logrus.New()

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

