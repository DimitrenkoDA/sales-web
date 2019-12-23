package salemaps

import "time"

type Salemap struct {
	ID        uint64    `json:"id"`
	ProdId    string    `json:"prod_id"`
	Since     time.Time `json:"since"`
	SubId     string    `json:"sub_id"`
	SalemanId uint64    `json:"saleman_id"`
	Quantity  uint64    `json:"quantity"`
	SaleDate  time.Time `json:"sale_date"`
	Note      *string   `json:"note"`
}
