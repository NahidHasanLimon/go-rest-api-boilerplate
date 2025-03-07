package routes

import (
	"go-rest-api/handlers"

	"github.com/gorilla/mux"
)

func Driverroutes(router *mux.Router) {
	router.HandleFunc("/drivers", handlers.GetDrivers).Methods("GET")
	router.HandleFunc("/drivers/{id}", handlers.GetDriver).Methods("GET")
	router.HandleFunc("/drivers/{id}", handlers.DeleteDriver).Methods("DELETE")
	router.HandleFunc("/drivers", handlers.AddDriver).Methods("POST")
	router.HandleFunc("/drivers/{id}", handlers.UpdateDriver).Methods("PUT")
}
