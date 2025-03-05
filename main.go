package main

import (
	"fmt"
	"myproject/config"
	"myproject/routes"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	port := ":8070"

	config.LoadConfig()

	config.ConnectDB()

	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	err := http.ListenAndServe(port, router)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
