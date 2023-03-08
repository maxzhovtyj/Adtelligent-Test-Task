package models

type Product struct {
	Id       int     `json:"id,omitempty"`
	Title    string  `json:"title"`
	SellerId int     `json:"sellerId"`
	Price    float64 `json:"price"` // float64 is not the best data type for money!!!
}
