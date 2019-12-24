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

func (s *Storage) Find(ctx context.Context, salemapID uint64, salemap *Salemap) error {
	q := `
		SELECT map_id as id, prod_id as prodId, dat as since, sub_id as subId, man_code as salemanId, quantity, sale_dat as saleDate, note 
		FROM salemaps
		WHERE map_id = $1
	`

	err := s.db.QueryRowContext(ctx, q, salemapID).Scan(&salemap.ID, &salemap.ProdId, &salemap.Since, &salemap.SubId, &salemap.SalemanId, &salemap.Quantity, &salemap.SaleDate, &salemap.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(ctx context.Context, salemap Salemap) error {
	q := `
		INSERT INTO salemaps (map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := s.db.ExecContext(ctx, q, salemap.ID, salemap.ProdId, salemap.Since, salemap.SubId, salemap.SalemanId, salemap.Quantity, salemap.SaleDate, salemap.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, salemapID uint64) error {
	q := `
		DELETE FROM salemaps WHERE map_id = $1
	`

	_, err := s.db.ExecContext(ctx, q, salemapID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, salemapID uint64, salemap Salemap) error {
	q := `
		UPDATE salemaps SET prod_id = $2, dat = $3, sub_id = $4, man_code = $5, quantity = $6, sale_dat = $7, note = $8 WHERE map_id = $1
	`
	_, err := s.db.ExecContext(ctx, q, salemapID, salemap.ProdId, salemap.Since, salemap.SubId, salemap.SalemanId, salemap.Quantity, salemap.SaleDate, salemap.Note)

	if err != nil {
		return err
	}
	return nil
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
