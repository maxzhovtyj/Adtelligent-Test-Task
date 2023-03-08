package models

type Seller struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}
