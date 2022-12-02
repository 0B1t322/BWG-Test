package models

import (
	"github.com/0B1t322/BWG-Test/internal/models/entity"
	"github.com/google/uuid"
)

const BalanceTableName = "Balances"

type BalanceField string

const (
	BalanceFieldID     BalanceField = `"Id"`
	BalanceFieldUserID BalanceField = `"UserId"`
	BalanceFieldAmount BalanceField = `"Balance"`
)

func (e BalanceField) String() string {
	return string(e)
}

func (e BalanceField) WithTable() string {
	return fieldWithTable(UserTableName, e.String())
}

type Balance struct {
	ID      uuid.UUID `gorm:"column:Id;type:uuid;primaryKey"`
	UserID  uuid.UUID `gorm:"column:UserId;type:uuid;"`
	Balance float64   `gorm:"column:Balance;type:double precision;"`

	// Edges
	User *User `gorm:"foreignKey:UserId;references:Id"`
}

func (Balance) TableName() string {
	return BalanceTableName
}

func BalanceModelFrom(balance *entity.Balance) Balance {
	return Balance{
		ID:      balance.ID,
		UserID:  balance.UserID,
		Balance: balance.Balance,
	}
}

func BalanceModelTo(balance *Balance) entity.Balance {
	return entity.Balance{
		ID:      balance.ID,
		UserID:  balance.UserID,
		Balance: balance.Balance,
	}
}
