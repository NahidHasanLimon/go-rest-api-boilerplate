package main

import (
	// "fmt"
	"go-rest-api/config"
	"go-rest-api/logger"
	"go-rest-api/routes"
	"go-rest-api/middleware"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	
)





func init() {
	// config.init()
	config.LoadConfig()
	config.ConnectDB()
	
}
func StartServer(router *mux.Router) {
	
	port := ":" + config.AppConfig.ServerPort
	logger.Log().Info("Server starting on port "+port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		logger.Log().Fatal("Failed to start server",err)
	}
	config.Log().Info("Server stareted on port "+port)
}

func main() {
	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)
	routes.RegisterRoutes(router)
	StartServer(router)
}
