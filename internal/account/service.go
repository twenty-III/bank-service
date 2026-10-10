package account

import (
	"context"
	"fmt"
	"uuid"
)

type Service struct {
	repo AccountRepo
}

func NewService(repo AccountRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) NewAccount(ctx context.Context, clientID uuid.UUID) (*Account, error) {
	a := &Account{
		ID:       uuid.New(),
		ClientID: clientID,
	}

	err := s.repo.NewAccount(ctx, a)
	if err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	return a, err
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Account, error) {
	account, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by id: %w", err)
	}

	return account, nil
}

func (s *Service) GetByClientID(ctx context.Context, clientID uuid.UUID) ([]Account, error) {
	accounts, err := s.repo.GetByClientID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts by client id: %w", err)
	}

	return accounts, nil
}