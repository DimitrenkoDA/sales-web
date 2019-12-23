package products

import "time"

type Product struct {
	ID             uint64    `json:"id"`
	Name           string    `json:"name"`
	Pincode        uint64    `json:"pincode"`
	ProductionDate time.Time `json:"production_date"`
	PgID           uint64    `json:"pg_id"`
	StatusID       uint64    `json:"status_id"`
	Note           *string   `json:"note"`
}
