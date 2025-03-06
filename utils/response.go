package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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
func SendError(w http.ResponseWriter, err interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	var message string
	switch v := err.(type) {

	case string:
		message = v

	case error:
		if _, ok := v.(validator.ValidationErrors); ok {
			message = CustomizeValidationError(v)
		} else {
			message = v.Error()
		}
	default:
		message = "An unknown error occurred"
	}

	json.NewEncoder(w).Encode(APIResponse{
		Status:  "error",
		Message: message,
	})
}
func CustomizeValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string
		for _, fieldError := range validationErrors {
			field := fieldError.Field() // Field name
			switch fieldError.Tag() {
			case "required":
				messages = append(messages, field+" is required")
			case "min":
				messages = append(messages, field+" must be at least "+fieldError.Param()+" characters long")
			case "max":
				messages = append(messages, field+" cannot exceed "+fieldError.Param()+" characters")
			default:
				messages = append(messages, field+" is invalid")
			}
		}
		return "Validation failed: " + messages[0] // Return first error (or join all)
	}
	return "Validation error"
}
