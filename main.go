package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"myproject/handlers"
)


func main() {
	port := ":8070"
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running on port %s", port)
	})
	// drivers = append(drivers, 100,200,3000,440404)

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
