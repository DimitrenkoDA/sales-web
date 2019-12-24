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

func (s *Storage) Find(ctx context.Context, subID uint64, sub *Sub) error {
	q := `
		SELECT sub_id as id, sub_name as name, address, phone, note 
		FROM subs
		WHERE sub_id = $1
	`

	err := s.db.QueryRow(q, subID).Scan(&sub.ID, &sub.Name, &sub.Address, &sub.Phone, &sub.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(ctx context.Context, sub Sub) error {
	q := `
		INSERT INTO subs (sub_id, sub_name, address, phone, note)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := s.db.ExecContext(ctx, q, sub.ID, sub.Name, sub.Address, sub.Phone, sub.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, subID uint64) error {
	q := `
		DELETE FROM subs WHERE sub_id = $1
	`

	_, err := s.db.ExecContext(ctx, q, subID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, subID uint64, sub Sub) error {
	q := `
		UPDATE subs SET sub_name = $2, address = $3, phone = $4, note = $5 WHERE sub_id = $1
	`
	_, err := s.db.ExecContext(ctx, q, sub.ID, sub.Name, sub.Address, sub.Phone, sub.Note)

	if err != nil {
		return err
	}
	return nil
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
