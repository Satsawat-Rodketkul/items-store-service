package models

import "time"

type TransactionMessage struct {
	TransactionId string
	UserId        string
	Title         string
	Price         float64
	Categoty      string
	CreateDate    time.Time
}
