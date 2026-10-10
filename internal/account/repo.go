package account

import (
	"context"
	"database/sql"
	"uuid"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) NewAccount(ctx context.Context, a *Account) error {
	query := `
		INSERT INTO account (id, client_id, balance)
		VALUES ($1, $2, $3)
		RETURNING created_at;
	`

	return r.db.QueryRowContext(ctx, query,
		a.ID,
		a.ClientID,
		a.Balance,
	).Scan(
		&a.CreatedAt,
	)
}

func (r *Repo) GetByID(ctx context.Context, id uuid.UUID) (*Account, error) {
	const query = `
		SELECT (id, client_id, balance, created_at)
		FROM account
		WHERE id = $1;
	`

	var a Account
	err := r.db.QueryRowContext(ctx, query,
		id,
	).Scan(
		&a.ID,
		&a.ClientID,
		&a.Balance,
		&a.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *Repo) GetByClientID(ctx context.Context, clientID uuid.UUID) ([]Account, error) {
	const query = `
		SELECT id, client_id, balance, created_at
		FROM account
		WHERE client_id = $1
		LIMIT 50;
	`

	rows, err := r.db.QueryContext(ctx, query,
		clientID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := make([]Account, 0)
	for rows.Next() {
		var a Account
		err = rows.Scan(&a.ID, &a.ClientID, &a.Balance, &a.CreatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
