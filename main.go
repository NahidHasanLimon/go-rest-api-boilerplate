package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"myproject/models"
)

var drivers []models.Driver



func getLastDriverID() int {
	if len(drivers) == 0 {
		return 0
	}
	return drivers[len(drivers)-1].ID
}

func getDrivers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get drivers!")
	json.NewEncoder(w).Encode(drivers)

}
func getDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var foundDriver *models.Driver

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i := range drivers {
		if drivers[i].ID == id {
			foundDriver = &drivers[i] // ✅ Points to the actual slice element
			break
		}
	}
	if foundDriver == nil {
		http.Error(w, "Driver not found", http.StatusNotFound) // ✅ Fixed error message
		return
	}

	json.NewEncoder(w).Encode(foundDriver)

}
func addDriver(w http.ResponseWriter, r *http.Request) {
	var driver models.Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	fmt.Println("pqrma s is", driver)

	drivers = append(drivers, models.Driver{
		ID:    getLastDriverID() + 1,
		Name:  driver.Name,
		Phone: driver.Phone,
	})
	fmt.Fprintf(w, "Driver added succesfully")
}

func updateDriver(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var inputtedDriverInfo models.Driver
	err = json.NewDecoder(r.Body).Decode(&inputtedDriverInfo)
	if err != nil {
	 http.Error(w, err.Error(), http.StatusBadRequest)
	 return
	}
	
	fmt.Println("inputed ", inputtedDriverInfo)


	for i, driver := range drivers {
		if driver.ID == id {
			drivers[i] =  inputtedDriverInfo;
			json.NewEncoder(w).Encode(
				map[string]string{
					"message": "Driver updated successfully",
				})
			break
		}
	}
	

		http.Error(w, "Driver Not found", http.StatusNotFound)
}

func deleteDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	newDrivers := []models.Driver{}

	for _, driver := range drivers {
		if driver.ID != id {
			newDrivers = append(newDrivers, driver)
		}
	}

	if len(newDrivers) == len(drivers) {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}

	drivers = newDrivers

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Driver deleted successfully",
		})
}

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

	router.HandleFunc("/drivers", getDrivers).Methods("GET")
	router.HandleFunc("/drivers/{id}", getDriver).Methods("GET")
	router.HandleFunc("/drivers/{id}", deleteDriver).Methods("DELETE")
	router.HandleFunc("/drivers", addDriver).Methods("POST")
	router.HandleFunc("/drivers/{id}", updateDriver).Methods("PUT")

	err := http.ListenAndServe(port, router)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
