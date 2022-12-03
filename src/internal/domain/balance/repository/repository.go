package balance

import (
	"context"
	"errors"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

var (
	ErrBalanceNotFound = errors.New("Balance not found")
	ErrBalanceExists   = errors.New("Balance with this user already exists")
)

type BalanceRepository interface {
	// GetBalance returns balance
	// If balance not found, it will return ErrBalanceNotFound
	GetBalance(ctx context.Context, id uuid.UUID) (aggregate.Balance, error)

	// GetBalanceForUser returns balance for user
	GetBalanceForUser(ctx context.Context, userID uuid.UUID) (aggregate.Balance, error)

	// CreateBalance creates new balance for user
	// If balance with this user already exists, it will return ErrBalanceExists
	CreateBalance(ctx context.Context, balance aggregate.Balance) error

	// UpdateBalance updates balance for user
	// If balance not found, it will return ErrBalanceNotFound
	UpdateBalance(ctx context.Context, balance aggregate.Balance) error
}
