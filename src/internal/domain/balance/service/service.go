package balance

import (
	"context"
	"errors"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

var (
	ErrBalanceNotFound = errors.New("Balance not found")
	ErrUserNotFound    = errors.New("User not found")
)

type BalanceService interface {
	// GetBalance returns balance
	// If balance not found, it will return ErrBalanceNotFound
	GetBalance(ctx context.Context, userID uuid.UUID) (aggregate.Balance, error)

	// UpdateBalance updates balance
	UpdateBalance(ctx context.Context, balanceID uuid.UUID, balance float64) error

	// CreateBalance creates new balance
	// If user not found, it will return ErrUserNotFound
	CreateBalance(ctx context.Context, userID uuid.UUID) (aggregate.Balance, error)
}
