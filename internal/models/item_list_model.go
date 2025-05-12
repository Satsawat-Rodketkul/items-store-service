package models

type ItemListResponse struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Categoty    string  `json:"categoty"`
	Image       string  `json:"image"`
}
