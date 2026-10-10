package account

import (
	"context"
	"time"
	"uuid"
)

type Account struct {
	ID        uuid.UUID
	ClientID  uuid.UUID
	Balance   int64 // stored in paise
	CreatedAt time.Time
}

type AccountRepo interface {
	NewAccount(ctx context.Context, a *Account) error
	GetByID(ctx context.Context, id uuid.UUID) (*Account, error)
	GetByClientID(ctx context.Context, id uuid.UUID) ([]Account, error)
}
