package models

type Customer struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}
