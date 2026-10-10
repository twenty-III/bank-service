package client

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

func (r *Repo) NewClient(ctx context.Context, c *Client) error {
	const query = `
		INSERT INTO client (id, name, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at;
	`

	return r.db.QueryRowContext(ctx, query,
		c.ID,
		c.Name,
		c.Email,
		c.Password,
	).Scan(
		&c.CreatedAt,
	)
}

func (r *Repo) GetByID(ctx context.Context, id uuid.UUID) (*Client, error) {
	return nil, nil
}

func (r *Repo) GetByEmail(ctx context.Context, email string) (*Client, error) {
	return nil, nil
}
