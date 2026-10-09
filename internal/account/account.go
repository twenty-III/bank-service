package account

import (
	"context"
	"time"
	"uuid"
)

type Account struct {
	ID        uuid.UUID
	ClientID  uuid.UUID
	Balance   float64
	CreatedAt time.Time
}

type AccountRepo interface {
	NewAccount(ctx context.Context, a *Account) error
	GetByID(ctx context.Context, id uuid.UUID) (*Account, error)
	GetByClientID(ctx context.Context, id uuid.UUID) ([]Account, error)
}
