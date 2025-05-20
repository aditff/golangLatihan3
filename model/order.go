package model

type Order struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Destination string  `json:"destination"`
	Price       float64 `json:"price"`
}
