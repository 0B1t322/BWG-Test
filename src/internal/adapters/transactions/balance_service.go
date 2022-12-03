package transactions

import (
	"context"

	balance "github.com/0B1t322/BWG-Test/internal/domain/balance/service"
	transactionssrv "github.com/0B1t322/BWG-Test/internal/domain/transaction/service"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

type BalanceService struct {
	service balance.BalanceService
}

func NewBalanceService(service balance.BalanceService) transactionssrv.BalanceService {
	return &BalanceService{service: service}
}

func (s *BalanceService) GetBalance(
	ctx context.Context,
	userID uuid.UUID,
) (aggregate.Balance, error) {
	b, err := s.service.GetBalance(ctx, userID)
	if err == balance.ErrBalanceNotFound {
		return aggregate.Balance{}, transactionssrv.ErrBalanceNotFound
	} else if err != nil {
		return aggregate.Balance{}, err
	}

	return b, nil
}

func (s *BalanceService) UpdateBalance(
	ctx context.Context,
	b aggregate.Balance,
) error {
	err := s.service.UpdateBalance(ctx, b.ID, b.GetBalance())
	if err == balance.ErrBalanceNotFound {
		return transactionssrv.ErrBalanceNotFound
	} else if err != nil {
		return err
	}

	return nil
}
