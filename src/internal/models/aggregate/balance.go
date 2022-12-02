package aggregate

import (
	"github.com/0B1t322/BWG-Test/internal/models/entity"
	"github.com/google/uuid"
)

type Balance struct {
	*entity.Balance
	// user edge
	User *User
}

func NewBalance(
	balance float64,
	userId uuid.UUID,
) Balance {
	return Balance{
		Balance: &entity.Balance{
			ID:      uuid.New(),
			UserID:  userId,
			Balance: balance,
		},
	}
}

func (b *Balance) AddAmount(amount float64) {
	b.Balance.Balance += amount
}

func (b *Balance) SubAmount(amount float64) {
	b.Balance.Balance -= amount
}

func (b Balance) GetBalance() float64 {
	return b.Balance.Balance
}

func (b *Balance) SetBalance(balance float64) {
	b.Balance.Balance = balance
}
