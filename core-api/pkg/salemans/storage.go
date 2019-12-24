package salemans

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

func (s *Storage) Find(ctx context.Context, salemanID uint64, saleman *Saleman) error {
	q := `
		SELECT man_code as id, saleman_name as name, card_code as code, n_dealer as dealer_id, status_id, note, condition 
		FROM salemans
		WHERE man_code = $1
	`

	err := s.db.QueryRow(q, salemanID).Scan(&saleman.ID, &saleman.Name, &saleman.Code, &saleman.DealerID, &saleman.StatusID, &saleman.Note, &saleman.Condition)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(ctx context.Context, saleman Saleman) error {
	q := `
		INSERT INTO salemans (man_code, saleman_name, card_code, n_dealer, status_id, note, condition)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := s.db.ExecContext(ctx, q, saleman.ID, saleman.Name, saleman.Code, saleman.DealerID, saleman.StatusID, saleman.Note, saleman.Condition)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, salemanID uint64) error {
	q := `
		DELETE FROM salemans WHERE man_code = $1
	`

	_, err := s.db.ExecContext(ctx, q, salemanID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, salemanID uint64, saleman Saleman) error {
	q := `
		UPDATE salemans SET saleman_name = $2, card_code = $3, n_dealer = $4, status_id = $5, note = $6, condition = $7 WHERE man_code = $1
	`
	_, err := s.db.ExecContext(ctx, q, saleman.ID, saleman.Name, saleman.Code, saleman.DealerID, saleman.StatusID, saleman.Note, saleman.Condition)

	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) List(ctx context.Context) ([]Saleman, error) {
	q := `
		SELECT man_code as id, saleman_name as name, card_code as code, n_dealer as dealer_id, status_id, note, condition 
		FROM salemans
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var salemans []Saleman

	for rows.Next() {
		var saleman Saleman

		err = rows.Scan(&saleman.ID, &saleman.Name, &saleman.Code, &saleman.DealerID, &saleman.StatusID, &saleman.Note, &saleman.Condition)

		if err != nil {
			return nil, err
		}

		salemans = append(salemans, saleman)
	}

	return salemans, nil
}
