package pricelist

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

func (s *Storage) List(ctx context.Context) ([]Pricelist, error) {
	q := `
		SELECT  prod_id, dat as since, price, note 
		FROM pricelist
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var pricelists []Pricelist

	for rows.Next() {
		var pricelist Pricelist

		err = rows.Scan(&pricelist.ProdID, &pricelist.Since, &pricelist.Price, &pricelist.Note)

		if err != nil {
			return nil, err
		}

		pricelists = append(pricelists, pricelist)
	}

	return pricelists, nil
}
