package models

type Order struct {
	Id         int `json:"id,omitempty"`
	ProductId  int `json:"productId"`
	CustomerId int `json:"customerId"`
}
