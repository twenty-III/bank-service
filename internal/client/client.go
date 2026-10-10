package client

import (
	"bank/internal/account"
	"context"
	"time"
	"uuid"
)

type Client struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type ClientRepo interface {
	NewClient(ctx context.Context, c *Client) error
}

type AccountProvider interface {
	NewAccount(ctx context.Context, clientID uuid.UUID) (*account.Account, error)
	GetByClientID(ctx context.Context, clientID uuid.UUID) ([]account.Account, error)
}
