package routes
import(
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"myproject/handlers"

)
func Driverroutes(router *mux.Router){
	router.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request on /debug")
		fmt.Fprintf(w, "Debug endpoint hit!")
	})
	router.HandleFunc("/drivers", handlers.GetDrivers).Methods("GET")
	router.HandleFunc("/drivers/{id}", handlers.GetDriver).Methods("GET")
	router.HandleFunc("/drivers/{id}", handlers.DeleteDriver).Methods("DELETE")
	router.HandleFunc("/drivers", handlers.AddDriver).Methods("POST")
	router.HandleFunc("/drivers/{id}", handlers.UpdateDriver).Methods("PUT")
}