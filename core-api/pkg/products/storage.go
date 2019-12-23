package products

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

func (s *Storage) List(ctx context.Context) ([]Product, error) {
	q := `
		SELECT prod_id as id, prod_name as name, pincode, data_iz as production_Date, pg_id, status_id, note 
		FROM products
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var product Product

		err = rows.Scan(&product.ID, &product.Name, &product.Pincode, &product.ProductionDate, &product.PgID, &product.StatusID, &product.Note)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
