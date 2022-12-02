package aggregate

import "github.com/0B1t322/BWG-Test/internal/models/entity"

type User struct {
	*entity.User
}

func NewUser(
	username string,
) User {
	return User{
		User: &entity.User{
			Username: username,
		},
	}
}
