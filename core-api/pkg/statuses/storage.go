package statuses

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

func (s *Storage) Find(ctx context.Context, statusID uint64, status *Status) error {
	q := `
		SELECT status_id as id, status_group as statusGroup, status_name as name, note 
		FROM statuses
		WHERE status_id = $1
	`

	err := s.db.QueryRowContext(ctx, q, statusID).Scan(&status.ID, &status.StutusGroup, &status.Name, &status.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(ctx context.Context, status Status) error {
	q := `
		INSERT INTO statuses (status_id, status_group, status_name, note)
		VALUES ($1, $2, $3, $4)
	`

	_, err := s.db.ExecContext(ctx, q, status.ID, status.StutusGroup, status.Name, status.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, statusID uint64) error {
	q := `
		DELETE FROM statuses WHERE status_id = $1
	`

	_, err := s.db.ExecContext(ctx, q, statusID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, dealID uint64, status Status) error {
	q := `
		UPDATE statuses SET status_group = $2, status_name = $3, note = $4 WHERE status_id = $1
	`
	_, err := s.db.ExecContext(ctx, q, dealID, status.StutusGroup, status.Name, status.Note)

	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) List(ctx context.Context) ([]Status, error) {
	q := `
		SELECT status_id as id, status_group as statusGroup, status_name as name, note 
		FROM statuses
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var statuses []Status

	for rows.Next() {
		var status Status

		err = rows.Scan(&status.ID, &status.StutusGroup, &status.Name, &status.Note)

		if err != nil {
			return nil, err
		}

		statuses = append(statuses, status)
	}

	return statuses, nil
}
