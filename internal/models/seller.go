package models

type Seller struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"passowrd,omitempty"`
}
