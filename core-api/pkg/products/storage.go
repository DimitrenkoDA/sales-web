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

func (s *Storage) Find(ctx context.Context, prodID uint64, product *Product) error {
	q := `
		SELECT prod_id as id, prod_name as name, pincode, data_iz as production_Date, pg_id, status_id, note 
		FROM products
		WHERE prod_id = $1
	`

	err := s.db.QueryRowContext(ctx, q, prodID).Scan(&product.ID, &product.Name, &product.Pincode, &product.ProductionDate, &product.PgID, &product.StatusID, &product.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(ctx context.Context, product Product) error {
	q := `
		INSERT INTO products (prod_id, prod_name, pincode, data_iz, pg_id, status_id, note)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := s.db.ExecContext(ctx, q, product.ID, product.Name, product.Pincode, product.ProductionDate, product.PgID, product.StatusID, product.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, prodID uint64) error {
	q := `
		DELETE FROM products WHERE prod_id = $1
	`

	_, err := s.db.ExecContext(ctx, q, prodID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, prodID uint64, product Product) error {
	q := `
		UPDATE deals SET prod_name = $2, pincode = $3, data_iz = $4, pg_id = $5, status_id = $6, note = $7 WHERE prod_id = $1
	`
	_, err := s.db.ExecContext(ctx, q, prodID, product.Name, product.Pincode, product.ProductionDate, product.PgID, product.StatusID, product.Note)

	if err != nil {
		return err
	}
	return nil
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
