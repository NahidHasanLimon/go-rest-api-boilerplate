package main

import (
	"fmt"	
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	 "encoding/json"
)


var drivers [] Driver

type Driver struct {
	ID int  `json:"id"`
	Name string  `json:"name"`
	Phone  string`json:"phone"`
}

func getLastDriverID() int {
	if len(drivers) == 0 {
		return 0 
	}
	return drivers[len(drivers)-1].ID
}


func getDrivers(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w, "get drivers!")
	json.NewEncoder(w).Encode(drivers)

}
func getDriver(w http.ResponseWriter , r *http.Request){
	   params := mux.Vars(r)
	   var foundDriver *Driver

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
	   if(foundDriver == nil){
			http.Error(w, "Driver not found", http.StatusNotFound) // ✅ Fixed error message
			return
	   }

	   fmt.Printf("Fetched Driver: %+v\n", *foundDriver)

	   json.NewEncoder(w).Encode(foundDriver)

}
func addDriver(w http.ResponseWriter , r *http.Request){
	// params := mux.Vars(r)

	var driver  Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if(err != nil){
		http.Error(w, "Invalid payload", http.StatusBadRequest)
				return
	}

	fmt.Println("pqrma s is", driver)


		drivers = append(drivers, Driver{
			ID:    getLastDriverID()+ 1,
			Name:  driver.Name,
			Phone: driver.Phone,
		})
		fmt.Fprintf(w, "Driver added succesfully")
	}
	func deleteDriver(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
	
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid driver ID", http.StatusBadRequest)
			return
		}
	
		newDrivers := []Driver{}
	
		for _, driver := range drivers {
			if driver.ID != id {
				newDrivers = append(newDrivers, driver)
			}
		}
		fmt.Println("lentght of drivers", len(drivers))
	
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



func main()  {
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
    
	router.HandleFunc("/drivers",getDrivers).Methods("GET")
	router.HandleFunc("/drivers/{id}",getDriver).Methods("GET")

	router.HandleFunc("/drivers/{id}",deleteDriver).Methods("DELETE")
	
	router.HandleFunc("/drivers",addDriver).Methods("POST")
	

	err := http.ListenAndServe(port, router)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
