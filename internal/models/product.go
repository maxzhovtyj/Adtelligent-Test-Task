package models

type Product struct {
	ID       int     `json:"id,omitempty"`
	Title    string  `json:"title"`
	SellerID int     `json:"sellerId"`
	Price    float64 `json:"price"` // float64 is not the best data type for money!!!
}
