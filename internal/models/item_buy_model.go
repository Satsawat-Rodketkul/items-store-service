package models

type ItemBuyRequest struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Categoty string  `json:"categoty"`
}
