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
