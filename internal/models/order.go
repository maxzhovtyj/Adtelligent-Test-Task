package models

type Order struct {
	ID         int `json:"id,omitempty"`
	ProductID  int `json:"productId"`
	CustomerID int `json:"customerId"`
}
