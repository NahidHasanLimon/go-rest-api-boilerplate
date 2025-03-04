package main

import (
	"fmt"
	"log"
	"myproject/config"
	"myproject/handlers"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)


func main() {
	port := ":8070"
	config.ConnectDB()
	router := mux.NewRouter()
	


	router.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request on /debug")
		fmt.Fprintf(w, "Debug endpoint hit!")
	})

	router.HandleFunc("/drivers", handlers.GetDrivers).Methods("GET")
	router.HandleFunc("/drivers/{id}", handlers.GetDriver).Methods("GET")
	router.HandleFunc("/drivers/{id}", handlers.DeleteDriver).Methods("DELETE")
	router.HandleFunc("/drivers", handlers.AddDriver).Methods("POST")
	router.HandleFunc("/drivers/{id}", handlers.UpdateDriver).Methods("PUT")

	err := http.ListenAndServe(port, router)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
