package client

import (
	"bank/internal/account"
	"context"
	"fmt"
	"uuid"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo            ClientRepo
	accountProvider AccountProvider
}

func NewService(repo ClientRepo, accountProvider AccountProvider) *Service {
	return &Service{
		repo:            repo,
		accountProvider: accountProvider,
	}
}

func (s *Service) NewClient(ctx context.Context, name, email, password string) (*Client, error) {
	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to create new client: %w", err)
	}

	c := &Client{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(passwordHash),
	}

	err = s.repo.NewClient(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("failed to create new client: %w", err)
	}

	return c, nil
}

func (s *Service) GetAllAccounts(ctx context.Context, clientID uuid.UUID) ([]account.Account, error) {
	return s.accountProvider.GetByClientID(ctx, clientID)
}

func (s *Service) NewAccount(ctx context.Context, clientID uuid.UUID) (*account.Account, error) {
	return s.accountProvider.NewAccount(ctx, clientID)
}

func hashPassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}
