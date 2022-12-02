package balance

import (
	"context"
	"errors"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

var (
	ErrBalanceNotFound = errors.New("Balance not found")
)

type BalanceService interface {
	// GetBalance returns balance
	GetBalance(ctx context.Context, userID string) (aggregate.Balance, error)

	// UpdateBalance updates balance
	UpdateBalance(ctx context.Context, balanceID uuid.UUID, balance float64) error

	// CreateBalance creates new balance
	CreateBalance(ctx context.Context, userID uuid.UUID) (aggregate.Balance, error)
}
