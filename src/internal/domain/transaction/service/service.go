package transaction

import (
	"context"
	"errors"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

var (
	ErrTransactionNotFound = errors.New("transaction not found")
)

type TransactionService interface {
	// CreateTransaction creates new transaction
	CreateTransaction(
		ctx context.Context,
		amount float64,
		operationType aggregate.OperationType,
		userID uuid.UUID,
	) (aggregate.Transaction, error)

	ExecuteTransaction(ctx context.Context, transaction aggregate.Transaction) error

	ExecuteTransactions(ctx context.Context, transactions []aggregate.Transaction) error

	// GetTransaction return transaction
	GetTransaction(ctx context.Context, transactionID uuid.UUID) (aggregate.Transaction, error)

	// GetAllTransactionsForUser return all transactions that was executed for user
	GetAllTransactionsForUser(
		ctx context.Context,
		userID uuid.UUID,
	) ([]aggregate.Transaction, error)
}
