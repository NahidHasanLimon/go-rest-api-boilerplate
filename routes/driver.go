package routes
import(
	"github.com/gorilla/mux"
	"myproject/handlers"

)
func Driverroutes(router *mux.Router){
	router.HandleFunc("/drivers", handlers.GetDrivers).Methods("GET")
	router.HandleFunc("/drivers/{id}", handlers.GetDriver).Methods("GET")
	router.HandleFunc("/drivers/{id}", handlers.DeleteDriver).Methods("DELETE")
	router.HandleFunc("/drivers", handlers.AddDriver).Methods("POST")
	router.HandleFunc("/drivers/{id}", handlers.UpdateDriver).Methods("PUT")
}