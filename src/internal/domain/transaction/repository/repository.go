package repository

import (
	"context"
	"errors"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

var (
	ErrTransactionNotFound = errors.New("transaction not found")
)

type TransactionRepository interface {
	// SaveTransaction saves transaction to the database
	// If transaction already exists, it will be updated
	SaveTransaction(ctx context.Context, transaction aggregate.Transaction) error

	// GetTransaction return transaction from the database
	GetTransaction(ctx context.Context, id uuid.UUID) (aggregate.Transaction, error)

	// GetTransactionsForUser return all transactions that not executed for user
	GetNotExecutedTransactionsForUser(
		ctx context.Context,
		userID uuid.UUID,
	) ([]aggregate.Transaction, error)

	// GetTransactionsHistoryForUser return all transactions that executed for user
	GetExecutedTransactionsHistoryForUser(
		ctx context.Context,
		userID uuid.UUID,
	) ([]aggregate.Transaction, error)

	GetAllTransactionsThatNotExecuted(
		ctx context.Context,
	) ([]aggregate.Transaction, error)
}
