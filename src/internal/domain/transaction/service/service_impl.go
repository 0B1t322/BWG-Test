package transaction

import (
	"context"

	transrepo "github.com/0B1t322/BWG-Test/internal/domain/transaction/repository"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

type BalanceService interface {
	// GetBalance return user balance
	// If balance not found, it will return ErrBalanceNotFound
	GetBalance(ctx context.Context, userID uuid.UUID) (aggregate.Balance, error)

	// UpdateBalance updates balance
	// If balance not found, it will return ErrBalanceNotFound
	UpdateBalance(ctx context.Context, balance aggregate.Balance) error
}

type TransactionServiceImpl struct {
	repo transrepo.TransactionRepository
	bs   BalanceService
}

func NewTransactionService(
	repo transrepo.TransactionRepository,
	bs BalanceService,
) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		repo: repo,
		bs:   bs,
	}
}

// CreateTransaction creates new transaction
func (t *TransactionServiceImpl) CreateTransaction(
	ctx context.Context,
	amount float64,
	operationType aggregate.OperationType,
	userID uuid.UUID,
) (aggregate.Transaction, error) {
	transaction := aggregate.NewTransaction(amount, operationType, userID)
	if err := t.repo.SaveTransaction(ctx, transaction); err != nil {
		return aggregate.Transaction{}, err
	}

	return transaction, nil
}

func (t *TransactionServiceImpl) ExecuteTransaction(
	ctx context.Context,
	transaction aggregate.Transaction,
) error {
	balance, err := t.bs.GetBalance(ctx, transaction.UserID)
	if err != nil {
		return err
	}

	if transaction.CanExecute(balance) {
		transaction.Execute(balance)
		transaction.SetSuccess()
	} else {
		transaction.SetDenied()
	}

	if transaction.IsSuccess() {
		if err := t.bs.UpdateBalance(ctx, balance); err != nil {
			return err
		}
	}

	if err := t.repo.SaveTransaction(ctx, transaction); err != nil {
		return err
	}

	return nil
}

func (t *TransactionServiceImpl) ExecuteTransactions(
	ctx context.Context,
	transactions []aggregate.Transaction,
) error {
	balance, err := t.bs.GetBalance(ctx, transactions[0].UserID)
	if err != nil {
		return err
	}

	for _, transaction := range transactions {
		if transaction.CanExecute(balance) {
			transaction.Execute(balance)
			transaction.SetSuccess()
		} else {
			transaction.SetDenied()
		}

		if err := t.repo.SaveTransaction(ctx, transaction); err != nil {
			return err
		}
	}

	if err := t.bs.UpdateBalance(ctx, balance); err != nil {
		return err
	}

	return nil
}

// GetTransaction return transaction
func (t *TransactionServiceImpl) GetTransaction(
	ctx context.Context,
	transactionID uuid.UUID,
) (aggregate.Transaction, error) {
	get, err := t.repo.GetTransaction(ctx, transactionID)
	if err == transrepo.ErrTransactionNotFound {
		return aggregate.Transaction{}, ErrTransactionNotFound
	} else if err != nil {
		return aggregate.Transaction{}, err
	}

	return get, nil
}

// GetAllTransactionsForUser return all transactions for user
func (t *TransactionServiceImpl) GetAllTransactionsForUser(
	ctx context.Context,
	userID uuid.UUID,
) ([]aggregate.Transaction, error) {
	ts, err := t.repo.GetExecutedTransactionsHistoryForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return ts, nil
}
