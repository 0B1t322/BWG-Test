package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type TransactionStatus int

const (
	TransactionStatusCreated TransactionStatus = iota
	TransactionStatusSuccess
	TransactionStatusDenied
)

func (t TransactionStatus) Int() int {
	return int(t)
}

type OperationType int

const (
	OperationTypeAdd OperationType = iota
	OperationTypeSub
)

type Transaction struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	Amount        float64
	Status        TransactionStatus
	OperationType OperationType
	CreatedAt     time.Time
	ExecutedAt    time.Time
}

func NewTransaction(
	amount float64,
	operationType OperationType,
	userID uuid.UUID,
) Transaction {
	return Transaction{
		ID:            uuid.New(),
		UserID:        userID,
		Amount:        amount,
		Status:        TransactionStatusCreated,
		OperationType: operationType,
		CreatedAt:     time.Now().UTC(),
	}
}

func (t Transaction) Execute(balance Balance) {
	switch t.OperationType {
	case OperationTypeAdd:
		balance.AddAmount(t.Amount)
	case OperationTypeSub:
		balance.SubAmount(t.Amount)
	}
	t.ExecutedAt = time.Now().UTC()
}

func (t *Transaction) SetDenied() {
	t.Status = TransactionStatusDenied
}

func (t *Transaction) SetSuccess() {
	t.Status = TransactionStatusSuccess
}

func (t Transaction) IsSuccess() bool {
	return t.Status == TransactionStatusSuccess
}

func (t Transaction) CanExecute(balance Balance) bool {
	if t.OperationType == OperationTypeSub && balance.GetBalance() < t.Amount {
		return false
	}

	return true
}
