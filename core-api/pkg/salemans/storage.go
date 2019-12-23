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
