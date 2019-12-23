package pricelist

import "time"

type Pricelist struct {
	ProdID uint64    `json:"id"`
	Since  time.Time `json:"since"`
	Price  uint64    `json:"price"`
	Note   *string   `json:"note"`
}
