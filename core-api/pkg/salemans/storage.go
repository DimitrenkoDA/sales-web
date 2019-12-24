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

func (s *Storage) ShowTop5(ctx context.Context) ([]Top5, error) {
	q := `
		SELECT LIST.saleman_name as name, LIST.CASH, TOP.RANK 
		FROM
		(SELECT CX.saleman_name, SUM(CH.QUANTITY * CZ.PRICE) as CASH 
			FROM salemans CX, salemaps CH, pricelist CZ 
			WHERE CX.MAN_CODE = CH.MAN_CODE AND CH.PROD_ID = CZ.PROD_ID AND CH.DAT = CZ.DAT AND to_char(ch.sale_dat, 'yyyy') = to_char(current_date - 365, 'yyyy')
			GROUP BY CX.saleman_name 
			ORDER BY CASH DESC) AS LIST 
		INNER JOIN 
		(SELECT CASH, RANK FROM (SELECT CASH, (ROW_NUMBER() over(ORDER BY CASH desc)) AS RANK
		  FROM (
			SELECT CX.saleman_name, SUM(CH.QUANTITY * CZ.PRICE) as CASH 
			FROM salemans CX, salemaps CH, pricelist CZ 
			WHERE CX.MAN_CODE = CH.MAN_CODE AND CH.PROD_ID = CZ.PROD_ID AND CH.DAT = CZ.DAT AND to_char(ch.sale_dat, 'yyyy') = to_char(current_date - 365, 'yyyy')
			GROUP BY CX.saleman_name 
			ORDER BY CASH DESC
		  ) as JJ)AS LOL WHERE RANK <= 3) AS TOP 
		ON LIST.CASH = TOP.CASH
	`

	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	var top5s []Top5

	for rows.Next() {
		var top5 Top5

		err = rows.Scan(&top5.Name, &top5.Cash, &top5.Rank)

		if err != nil {
			return nil, err
		}

		top5s = append(top5s, top5)
	}

	return top5s, nil
}

func (s *Storage) UnsoldProduct(ctx context.Context, salemanName string, leftDate string, rightDate string) ([]Unsold, error) {
	q := `
		SELECT DISTINCT A.prod_id, A.prod_name
		FROM products A, salemans B, salemaps C
		WHERE A.prod_id = C.prod_id AND
		B.MAN_CODE = C.MAN_CODE AND
		UPPER(B.saleman_name) <> UPPER($1) AND
		C.SALE_DAT > to_date($2, 'MM.DD.YYYY') AND
		C.SALE_DAT < to_date($3, 'MM.DD.YYYY')
	`

	rows, err := s.db.QueryContext(ctx, q, salemanName, leftDate, rightDate)

	if err != nil {
		return nil, err
	}

	var unsolds []Unsold

	for rows.Next() {
		var unsold Unsold

		err = rows.Scan(&unsold.ID, &unsold.Name)

		if err != nil {
			return nil, err
		}

		unsolds = append(unsolds, unsold)
	}

	return unsolds, nil
}

func (s *Storage) Find(ctx context.Context, salemanID uint64, saleman *Saleman) error {
	q := `
		SELECT man_code as id, saleman_name as name, card_code as code, n_dealer as dealer_id, status_id, note, condition 
		FROM salemans
		WHERE man_code = $1
	`

	err := s.db.QueryRowContext(ctx, q, salemanID).Scan(&saleman.ID, &saleman.Name, &saleman.Code, &saleman.DealerID, &saleman.StatusID, &saleman.Note, &saleman.Condition)

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
	_, err := s.db.ExecContext(ctx, q, salemanID, saleman.Name, saleman.Code, saleman.DealerID, saleman.StatusID, saleman.Note, saleman.Condition)

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
