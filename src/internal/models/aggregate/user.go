package aggregate

import (
	"github.com/0B1t322/BWG-Test/internal/models/entity"
	"github.com/google/uuid"
)

type User struct {
	*entity.User

	// Balance edge
	Balance *Balance
}

func NewUser(
	username string,
) User {
	return User{
		User: &entity.User{
			ID:       uuid.New(),
			Username: username,
		},
	}
}
