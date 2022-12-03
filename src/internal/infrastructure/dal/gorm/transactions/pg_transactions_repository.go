package transactions

import (
	"context"

	"github.com/0B1t322/BWG-Test/internal/domain/transaction/repository"
	"github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm/models"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
	"github.com/samber/lo"

	"gorm.io/gorm"
)

type PGTransactionsRepository struct {
	db *gorm.DB
}

func NewPGTransactionsRepository(db *gorm.DB) *PGTransactionsRepository {
	return &PGTransactionsRepository{db: db}
}

// SaveTransaction saves transaction to the database
// If transaction already exists, it will be updated
func (p *PGTransactionsRepository) SaveTransaction(
	ctx context.Context,
	transaction aggregate.Transaction,
) error {
	// Try to get transaction from the database
	_, err := p.GetTransaction(ctx, transaction.ID)
	if err == nil { // Transaction already exists
		return p.updateTransaction(ctx, transaction)
	} else if err == repository.ErrTransactionNotFound { // Transaction not exists
		return p.createTransaction(ctx, transaction)
	}

	return err
}

// GetAllTransactionsThatNotExecuted return all transactions that not executed
func (p *PGTransactionsRepository) GetAllTransactionsThatNotExecuted(
	ctx context.Context,
) ([]aggregate.Transaction, error) {
	var get []models.Transaction
	{
		result := p.db.WithContext(ctx).
			Model(&models.Transaction{}).
			Where(models.TransactionFieldStatus.WithTable()+" = ?", aggregate.TransactionStatusCreated).
			Order(models.TransactionFieldCreatedAt + "ASC").
			Find(&get)

		if result.Error != nil {
			return nil, result.Error
		}
	}

	return lo.Map(
		get,
		func(model models.Transaction, _ int) aggregate.Transaction {
			return models.TransactionModelTo(&model)
		},
	), nil
}

// GetTransaction return transaction from the database
// If transaction not found, it will return ErrTransactionNotFound
func (p *PGTransactionsRepository) GetTransaction(
	ctx context.Context,
	id uuid.UUID,
) (aggregate.Transaction, error) {
	var get models.Transaction
	{
		result := p.db.WithContext(ctx).
			Model(&models.Transaction{}).
			Where(models.TransactionFieldID.WithTable()+" = ?", id).
			First(&get)

		if result.Error == gorm.ErrRecordNotFound {
			return aggregate.Transaction{}, repository.ErrTransactionNotFound
		} else if result.Error != nil {
			return aggregate.Transaction{}, result.Error
		}
	}

	return models.TransactionModelTo(&get), nil
}

// GetTransactionsForUser return all transactions that not executed for user
func (p *PGTransactionsRepository) GetNotExecutedTransactionsForUser(
	ctx context.Context,
	userID uuid.UUID,
) ([]aggregate.Transaction, error) {
	var get []models.Transaction
	{
		result := p.db.WithContext(ctx).
			Model(&models.Transaction{}).
			Where(models.TransactionFieldUserID.WithTable()+" = ?", userID).
			Where(models.TransactionFieldStatus.WithTable()+" = ?", aggregate.TransactionStatusCreated).
			Order(models.TransactionFieldCreatedAt + "ASC").
			Find(&get)

		if result.Error != nil {
			return nil, result.Error
		}
	}

	return lo.Map(
		get,
		func(model models.Transaction, _ int) aggregate.Transaction {
			return models.TransactionModelTo(&model)
		},
	), nil
}

// GetTransactionsHistoryForUser return all transactions that executed for user
func (p *PGTransactionsRepository) GetExecutedTransactionsHistoryForUser(
	ctx context.Context,
	userID uuid.UUID,
) ([]aggregate.Transaction, error) {
	var get []models.Transaction
	{
		result := p.db.WithContext(ctx).
			Model(&models.Transaction{}).
			Where(models.TransactionFieldUserID.WithTable()+" = ?", userID).
			Where(models.TransactionFieldStatus.WithTable()+" != ?", aggregate.TransactionStatusCreated).
			Order(models.TransactionFieldCreatedAt + "ASC").
			Find(&get)

		if result.Error != nil {
			return nil, result.Error
		}
	}

	return lo.Map(
		get,
		func(model models.Transaction, _ int) aggregate.Transaction {
			return models.TransactionModelTo(&model)
		},
	), nil
}

func (p *PGTransactionsRepository) createTransaction(
	ctx context.Context,
	transaction aggregate.Transaction,
) error {
	model := models.TransactionModelFrom(transaction)
	{
		result := p.db.WithContext(ctx).
			Model(&models.Transaction{}).
			Create(model)

		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (p *PGTransactionsRepository) updateTransaction(
	ctx context.Context,
	transaction aggregate.Transaction,
) error {
	model := models.TransactionModelFrom(transaction)
	{
		result := p.db.WithContext(ctx).
			Model(&models.Transaction{}).
			Where(models.TransactionFieldID.WithTable()+" = ?", transaction.ID).
			Updates(model)

		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
