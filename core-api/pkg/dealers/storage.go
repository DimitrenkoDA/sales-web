package dealers

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

func (s *Storage) Find(ctx context.Context, dealerID uint64, dealer *Dealer) error {
	q := `
		SELECT n_dealer as id, dealer_name as name, address, phone, status_id, note
		FROM dealers
		WHERE n_dealer = $1
	`

	err := s.db.QueryRowContext(ctx, q, dealerID).Scan(&dealer.ID, &dealer.Name, &dealer.Address, &dealer.Phone, &dealer.StatusID, &dealer.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(ctx context.Context, dealer Dealer) error {
	q := `
		INSERT INTO dealers (n_dealer, dealer_name, address, phone, status_id, note)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := s.db.ExecContext(ctx, q, dealer.ID, dealer.Name, dealer.Address, dealer.Phone, dealer.StatusID, dealer.Note)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, dealerID uint64) error {
	q := `
		DELETE FROM dealers WHERE n_dealer = $1
	`

	_, err := s.db.ExecContext(ctx, q, dealerID)

	if err != nil {
		return err
	}

	return nil
}
func (s *Storage) Update(ctx context.Context, dealerID uint64, dealer Dealer) error {
	q := `
		UPDATE dealers SET dealer_name = $2, address = $3, phone = $4, status_id = $5, note = $6 WHERE n_dealer = $1
	`
	_, err := s.db.ExecContext(ctx, q, dealerID, dealer.Name, dealer.Address, dealer.Phone, dealer.StatusID, dealer.Note)

	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) List(ctx context.Context) ([]Dealer, error) {
	q := `
		SELECT n_dealer as id, dealer_name as name, address, phone, status_id, note 
		FROM dealers
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var dealers []Dealer

	for rows.Next() {
		var dealer Dealer

		err = rows.Scan(&dealer.ID, &dealer.Name, &dealer.Address, &dealer.Phone, &dealer.StatusID, &dealer.Note)

		if err != nil {
			return nil, err
		}

		dealers = append(dealers, dealer)
	}

	return dealers, nil
}
