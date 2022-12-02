package models

import (
	"github.com/0B1t322/BWG-Test/internal/models/entity"
	"github.com/google/uuid"
)

const UserTableName = "Users"

type UserField string

const (
	UserFieldID       UserField = `"ID"`
	UserFieldUsername UserField = `"Username"`
)

func (e UserField) String() string {
	return string(e)
}

func (e UserField) WithTable() string {
	return fieldWithTable(UserTableName, e.String())
}

type User struct {
	ID       uuid.UUID `gorm:"column:Id;type:uuid;primaryKey"                json:"id"`
	Username string    `gorm:"column:Username;type:varchar(255);uniqueIndex" json:"username"`
}

func (User) TableName() string {
	return UserTableName
}

func UserModelFrom(user *entity.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
	}
}

func UserModelTo(user *User) entity.User {
	return entity.User{
		ID:       user.ID,
		Username: user.Username,
	}
}
