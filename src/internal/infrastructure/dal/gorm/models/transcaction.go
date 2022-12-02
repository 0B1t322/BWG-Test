package models

import (
	"time"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

const TransactionTableName = "Transactions"

type TransactionField string

func (e TransactionField) String() string {
	return string(e)
}

func (e TransactionField) WithTable() string {
	return fieldWithTable(TransactionTableName, e.String())
}

const (
	TransactionFieldID         TransactionField = `"Id"`
	TransactionFieldUserID     TransactionField = `"UserId"`
	TransactionFieldAmount     TransactionField = `"Amount"`
	TransactionFieldCreatedAt  TransactionField = `"CreatedAt"`
	TransactionFieldExecutedAt TransactionField = `"ExecutedAt"`
	TransactionFieldStatus     TransactionField = `"Status"`
)

type Transaction struct {
	ID         uuid.UUID `gorm:"column:Id;type:uuid;primaryKey"`
	UserID     uuid.UUID `gorm:"column:UserId;type:uuid;"`
	Amount     float64   `gorm:"column:Amount;type:double precision;"`
	Status     int       `gorm:"column:Status;type:int;"`
	CreatedAt  time.Time `gorm:"column:CreatedAt;type:timestamp;"`
	ExecutedAt time.Time `gorm:"column:ExecutedAt;type:timestamp;"`

	// Edges
	User    *User    `gorm:"foreignKey:UserId;references:Id"`
	Balance *Balance `gorm:"foreignKey:UserId"`
}

func (Transaction) TableName() string {
	return TransactionTableName
}

func TransactionModelFrom(transaction aggregate.Transaction) Transaction {
	return Transaction{
		ID:         transaction.ID,
		UserID:     transaction.Balance.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status.Int(),
		CreatedAt:  transaction.CreatedAt,
		ExecutedAt: transaction.ExecutedAt,

		Balance: lo.ToPtr(BalanceModelFrom(transaction.Balance)),
	}
}

func TransactionModelTo(transaction *Transaction) aggregate.Transaction {
	return aggregate.Transaction{
		ID:         transaction.ID,
		Balance:    lo.ToPtr(BalanceModelTo(transaction.Balance)),
		Amount:     transaction.Amount,
		Status:     aggregate.TransactionStatus(transaction.Status),
		CreatedAt:  transaction.CreatedAt,
		ExecutedAt: transaction.ExecutedAt,
	}
}
