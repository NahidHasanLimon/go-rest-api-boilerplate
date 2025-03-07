package main

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/routes"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func init() {
	config.LoadConfig()
	config.ConnectDB()
}
func StartServer(router *mux.Router) {
	port := ":" + config.AppConfig.ServerPort
	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	StartServer(router)
}
