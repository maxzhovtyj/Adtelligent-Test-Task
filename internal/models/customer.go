package models

type Customer struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}
