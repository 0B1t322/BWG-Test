package models

import (
	"time"

	"github.com/google/uuid"
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
)

type Transaction struct {
	ID         uuid.UUID `gorm:"column:Id;type:uuid;primaryKey"`
	UserID     string    `gorm:"column:UserId;type:uuid;"`
	Amount     float64   `gorm:"column:Amount;type:double precision;"`
	Status     int       `gorm:"column:Status;type:int;"`
	CreatedAt  time.Time `gorm:"column:CreatedAt;type:timestamp;"`
	ExecutedAt time.Time `gorm:"column:ExecutedAt;type:timestamp;"`

	// Edges
	User    *User      `gorm:"foreignKey:UserId;references:Id"`
	Balance []*Balance `gorm:"foreignKey:UserId"`
}

func (Transaction) TableName() string {
	return TransactionTableName
}
