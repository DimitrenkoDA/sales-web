package salemaps

import (
	"context"
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) List(ctx context.Context) ([]Salemap, error) {
	q := `
		SELECT map_id as id, prod_id as prodId, dat as since, sub_id as subId, man_code as salemanId, quantity, sale_dat as saleDate, note 
		FROM salemaps
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var salemaps []Salemap

	for rows.Next() {
		var salemap Salemap

		err = rows.Scan(&salemap.ID, &salemap.ProdId, &salemap.Since, &salemap.SubId, &salemap.SalemanId, &salemap.Quantity, &salemap.SaleDate, &salemap.Note)

		if err != nil {
			return nil, err
		}

		salemaps = append(salemaps, salemap)
	}

	return salemaps, nil
}
