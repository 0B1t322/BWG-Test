package balance

import (
	"context"

	balancerepo "github.com/0B1t322/BWG-Test/internal/domain/balance/repository"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

type UserGetter interface {
	// If user not found, it will return ErrUserNotFound
	GetUser(ctx context.Context, userID string) (aggregate.User, error)
}

type BalanceServiceImpl struct {
	repo balancerepo.BalanceRepository
	ug   UserGetter
}

func NewBalanceService(repo balancerepo.BalanceRepository, ug UserGetter) *BalanceServiceImpl {
	return &BalanceServiceImpl{
		repo: repo,
		ug:   ug,
	}
}

// GetBalance returns balance
func (b *BalanceServiceImpl) GetBalance(
	ctx context.Context,
	userID string,
) (aggregate.Balance, error) {
	balance, err := b.repo.GetBalanceForUser(ctx, userID)
	if err == balancerepo.ErrBalanceNotFound {
		return aggregate.Balance{}, ErrBalanceNotFound
	} else if err != nil {
		return aggregate.Balance{}, err
	}

	return balance, nil
}

// UpdateBalance updates balance
func (b *BalanceServiceImpl) UpdateBalance(
	ctx context.Context,
	balanceID uuid.UUID,
	balance float64,
) error {
	balanceModel, err := b.repo.GetBalance(ctx, balanceID)
	if err == balancerepo.ErrBalanceNotFound {
		return ErrBalanceNotFound
	} else if err != nil {
		return err
	}

	balanceModel.SetBalance(balance)

	err = b.repo.UpdateBalance(ctx, balanceModel)
	if err != nil {
		return err
	}

	return nil
}

// CreateBalance creates new balance
func (b *BalanceServiceImpl) CreateBalance(
	ctx context.Context,
	userID uuid.UUID,
) (aggregate.Balance, error) {
	if _, err := b.ug.GetUser(ctx, userID.String()); err != nil {
		return aggregate.Balance{}, ErrUserNotFound
	}

	balance := aggregate.NewBalance(0, userID)

	err := b.repo.CreateBalance(ctx, balance)
	if err != nil {
		return aggregate.Balance{}, err
	}

	return balance, nil
}
