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
