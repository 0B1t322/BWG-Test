package models

import (
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/0B1t322/BWG-Test/internal/models/entity"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

const UserTableName = "Users"

type UserField string

const (
	UserFieldID       UserField = `"Id"`
	UserFieldUsername UserField = `"Username"`
)

func (e UserField) String() string {
	return string(e)
}

func (e UserField) WithTable() string {
	return fieldWithTable(UserTableName, e.String())
}

type UserEdge string

func (u UserEdge) String() string {
	return string(u)
}

const (
	UserEdgeBalance UserEdge = "Balance"
)

type User struct {
	ID       uuid.UUID `gorm:"column:Id;type:uuid;primaryKey"                json:"id"`
	Username string    `gorm:"column:Username;type:varchar(255);uniqueIndex" json:"username"`

	// Balance edge
	Balance *Balance `gorm:"foreignKey:UserId;references:Id" json:"balance"`
}

func (User) TableName() string {
	return UserTableName
}

func UserModelFrom(user aggregate.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
	}
}

func UserModelTo(user *User) aggregate.User {
	return aggregate.User{
		User: &entity.User{
			ID:       user.ID,
			Username: user.Username,
		},
		Balance: lo.TernaryF(
			user.Balance != nil,
			func() *aggregate.Balance {
				return lo.ToPtr(BalanceModelTo(user.Balance))
			},
			func() *aggregate.Balance {
				return nil
			},
		),
	}
}
