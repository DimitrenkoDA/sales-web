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

func (s *Storage) Find(ctx context.Context, dealID uint64, deal *Deal) error {
	q := `
		SELECT n_deal as id, data_start as startedAt, data_finish as finishedAt, prim as note, sub_id 
		FROM deals
		WHERE n_deal = $1
	`

	err := s.db.QueryRowContext(ctx, q, dealID).Scan(&deal.ID, &deal.StartedAt, &deal.FinishedAt, &deal.Note, &deal.SubId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(ctx context.Context, deal Deal) error {
	q := `
		INSERT INTO deals (n_deal, data_start, data_finish, prim, sub_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := s.db.ExecContext(ctx, q, deal.ID, deal.StartedAt, deal.FinishedAt, deal.Note, deal.SubId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, dealID uint64) error {
	q := `
		DELETE FROM deals WHERE n_deal = $1
	`

	_, err := s.db.ExecContext(ctx, q, dealID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, dealID uint64, deal Deal) error {
	q := `
		UPDATE deals SET data_start = $2, data_finish = $3, prim = $4, sub_id = $5 WHERE n_deal = $1
	`
	_, err := s.db.ExecContext(ctx, q, dealID, deal.StartedAt, deal.FinishedAt, deal.Note, deal.SubId)

	if err != nil {
		return err
	}
	return nil
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
