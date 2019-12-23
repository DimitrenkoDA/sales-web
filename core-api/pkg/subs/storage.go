package subs

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

func (s *Storage) List(ctx context.Context) ([]Sub, error) {
	q := `
		SELECT sub_id as id, sub_name as name, address, phone, note 
		FROM subs
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var subs []Sub

	for rows.Next() {
		var sub Sub

		err = rows.Scan(&sub.ID, &sub.Name, &sub.Address, &sub.Phone, &sub.Note)

		if err != nil {
			return nil, err
		}

		subs = append(subs, sub)
	}

	return subs, nil
}
