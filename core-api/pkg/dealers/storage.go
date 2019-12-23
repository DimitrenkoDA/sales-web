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
