package main

import (
	// "fmt"
	"go-rest-api/config"
	"go-rest-api/routes"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)
var logger *zap.Logger


func init() {
	
	config.LoadConfig()
	config.ConnectDB()
	var err error
	logger, _ = zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger")
	}
}
func StartServer(router *mux.Router) {
	
	port := ":" + config.AppConfig.ServerPort
	logger.Info("Server stareting on port "+port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
	logger.Info("Server stareted on port "+port)
}

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	StartServer(router)
}
