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

type OperationType int

const (
	OperationTypeAdd OperationType = iota
	OperationTypeSub
)

type Transaction struct {
	ID            string
	Balance       *entity.Balance
	Amount        float64
	Status        TransactionStatus
	OperationType OperationType
	CreatedAt     time.Time
	ExecutedAt    time.Time
}

func NewTransaction(
	amount float64,
	operationType OperationType,
	balance *entity.Balance,
) Transaction {
	return Transaction{
		Balance:       balance,
		Amount:        amount,
		Status:        TransactionStatusCreated,
		OperationType: operationType,
		CreatedAt:     time.Now().UTC(),
	}
}

func (t Transaction) Execute() {
	switch t.OperationType {
	case OperationTypeAdd:
		t.Balance.Balance += t.Amount
	case OperationTypeSub:
		t.Balance.Balance -= t.Amount
	}
	t.ExecutedAt = time.Now().UTC()
}

func (t *Transaction) SetDenied() {
	t.Status = TransactionStatusDenied
}

func (t *Transaction) SetSuccess() {
	t.Status = TransactionStatusSuccess
}

func (t Transaction) CanExecute() bool {
	if t.OperationType == OperationTypeSub && t.Balance.Balance < t.Amount {
		return false
	}

	return true
}
