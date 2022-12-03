package balance

import (
	"context"

	balance "github.com/0B1t322/BWG-Test/internal/domain/balance/repository"
	"github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm/models"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PGBalanceRepository struct {
	db *gorm.DB
}

func NewPGBalanceRepository(db *gorm.DB) *PGBalanceRepository {
	return &PGBalanceRepository{db: db}
}

// GetBalance returns balance
// If balance not found, it will return ErrBalanceNotFound
func (b *PGBalanceRepository) GetBalance(
	ctx context.Context,
	id uuid.UUID,
) (aggregate.Balance, error) {
	var get models.Balance
	{
		result := b.db.WithContext(ctx).
			Model(&models.Balance{}).
			Where(models.BalanceFieldID.String()+" = ?", id).
			First(&get)
		if result.Error == gorm.ErrRecordNotFound {
			return aggregate.Balance{}, balance.ErrBalanceNotFound
		} else if result.Error != nil {
			return aggregate.Balance{}, result.Error
		}
	}

	return models.BalanceModelTo(&get), nil
}

// GetBalanceForUser returns balance for user
func (b *PGBalanceRepository) GetBalanceForUser(
	ctx context.Context,
	userID uuid.UUID,
) (aggregate.Balance, error) {
	var get models.Balance
	{
		result := b.db.WithContext(ctx).
			Model(&models.Balance{}).
			Where(models.BalanceFieldUserID.String()+" = ?", userID).
			First(&get)

		if result.Error == gorm.ErrRecordNotFound {
			return aggregate.Balance{}, balance.ErrBalanceNotFound
		} else if result.Error != nil {
			return aggregate.Balance{}, result.Error
		}
	}

	return models.BalanceModelTo(&get), nil
}

// CreateBalance creates new balance for user
func (b *PGBalanceRepository) CreateBalance(ctx context.Context, cb aggregate.Balance) error {
	_, err := b.GetBalance(ctx, cb.ID)
	if err == nil {
		return balance.ErrBalanceExists
	} else if err == balance.ErrBalanceNotFound {
		// Pass
	} else if err != nil {
		return err
	}

	result := b.db.WithContext(ctx).
		Model(&models.Balance{}).
		Create(models.BalanceModelFrom(cb))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateBalance updates balance for user
// If balance not found, it will return ErrBalanceNotFound
func (b *PGBalanceRepository) UpdateBalance(ctx context.Context, balance aggregate.Balance) error {
	_, err := b.GetBalance(ctx, balance.ID)
	if err != nil {
		return err
	}

	model := models.BalanceModelFrom(balance)

	result := b.db.WithContext(ctx).
		Model(&models.Balance{}).
		Where(models.BalanceFieldID.String()+" = ?", balance.ID).
		Updates(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
