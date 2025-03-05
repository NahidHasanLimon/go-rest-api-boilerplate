package main

import (
	"fmt"
	"myproject/config"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"myproject/routes"
)


func main() {
	port := ":8070"
	config.ConnectDB()
	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	err := http.ListenAndServe(port, router)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
