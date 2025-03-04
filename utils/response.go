package utils
import(
	"net/http"
	"encoding/json"
)

type APIResponse struct {
	Status    string    `json:"status"`
	Message  string `json:"message"`
	Data  interface{} `json:"data,omitempty"`
} 

func SendSuccess(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}
func SendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(APIResponse{
		Status:  "error",
		Message: message,
	})
}