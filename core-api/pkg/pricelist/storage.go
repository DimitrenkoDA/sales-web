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

func (s *Storage) Find(ctx context.Context, prodID uint64, pricelist *Pricelist) error {
	q := `
		SELECT  prod_id, dat as since, price, note 
		FROM pricelist
		WHERE prod_id = $1
	`

	err := s.db.QueryRowContext(ctx, q, prodID).Scan(&pricelist.ProdID, &pricelist.Since, &pricelist.Price, &pricelist.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(ctx context.Context, pricelist Pricelist) error {
	q := `
		INSERT INTO deals (prod_id, dat, price, note)
		VALUES ($1, $2, $3, $4)
	`

	_, err := s.db.ExecContext(ctx, q, pricelist.ProdID, pricelist.Since, pricelist.Price, pricelist.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, prodID uint64) error {
	q := `
		DELETE FROM pricelist WHERE prod_id = $1
	`

	_, err := s.db.ExecContext(ctx, q, prodID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, prodID uint64, pricelist Pricelist) error {
	q := `
		UPDATE pricelist SET dat = $2, price = $3, note = $4 WHERE prod_id = $1
	`
	_, err := s.db.ExecContext(ctx, q, prodID, pricelist.Since, pricelist.Price, pricelist.Note)

	if err != nil {
		return err
	}
	return nil
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
