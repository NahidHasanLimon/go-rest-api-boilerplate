package models
type Driver struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required,min=2,max=100"`
	Phone string `json:"phone"`
}