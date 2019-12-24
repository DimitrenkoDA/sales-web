package product_groups

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

func (s *Storage) Find(ctx context.Context, PgID uint64, productGroup *ProductGroup) error {
	q := `
		SELECT  pg_id as id, pg_name as name, note 
		FROM product_groups
		WHERE pg_id = $1
	`

	err := s.db.QueryRowContext(ctx, q, PgID).Scan(&productGroup.ID, &productGroup.Name, &productGroup.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(ctx context.Context, productGroup ProductGroup) error {
	q := `
		INSERT INTO product_group (pg_id, pg_name, note)
		VALUES ($1, $2, $3)
	`

	_, err := s.db.ExecContext(ctx, q, productGroup.ID, productGroup.Name, productGroup.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, PgID uint64) error {
	q := `
		DELETE FROM product_groups WHERE pg_id = $1
	`

	_, err := s.db.ExecContext(ctx, q, PgID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, PgID uint64, productGroup ProductGroup) error {
	q := `
		UPDATE product_groups SET pg_name = $2, note = $3 WHERE pg_id = $1
	`
	_, err := s.db.ExecContext(ctx, q, PgID, productGroup.Name, productGroup.Note)

	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) List(ctx context.Context) ([]ProductGroup, error) {
	q := `
		SELECT  pg_id as id, pg_name as name, note 
		FROM product_groups
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var productGroups []ProductGroup

	for rows.Next() {
		var productGroup ProductGroup

		err = rows.Scan(&productGroup.ID, &productGroup.Name, &productGroup.Note)

		if err != nil {
			return nil, err
		}

		productGroups = append(productGroups, productGroup)
	}

	return productGroups, nil
}
