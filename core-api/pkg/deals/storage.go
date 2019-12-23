package deals

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

func (s *Storage) List(ctx context.Context) ([]Deal, error) {
	q := `
		SELECT n_deal as id, data_start as startedAt, data_finish as finishedAt, prim as note, sub_id 
		FROM deals
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var deals []Deal

	for rows.Next() {
		var deal Deal

		err = rows.Scan(&deal.ID, &deal.StartedAt, &deal.FinishedAt, &deal.Note, &deal.SubId)

		if err != nil {
			return nil, err
		}

		deals = append(deals, deal)
	}

	return deals, nil
}
