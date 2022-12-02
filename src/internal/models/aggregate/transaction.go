package aggregate

import (
	"time"

	"github.com/0B1t322/BWG-Test/internal/models/entity"
)

type TransactionStatus int

const (
	TransactionStatusCreated TransactionStatus = iota
	TransactionStatusSuccess
	TransactionStatusDenied
)

type Transaction struct {
	ID         string
	Balance    *entity.Balance
	Amount     float64
	Status     TransactionStatus
	CreatedAt  time.Time
	ExecutedAt time.Time
}

func NewTransaction(
	amount float64,
	balance *entity.Balance,
) Transaction {
	return Transaction{
		Balance:   balance,
		Amount:    amount,
		Status:    TransactionStatusCreated,
		CreatedAt: time.Now().UTC(),
	}
}

func (t Transaction) Execute() {
	t.Balance.Balance += t.Amount
	t.ExecutedAt = time.Now().UTC()
}

func (t *Transaction) SetDenied() {
	t.Status = TransactionStatusDenied
}

func (t *Transaction) SetSuccess() {
	t.Status = TransactionStatusSuccess
}
