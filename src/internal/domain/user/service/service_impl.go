package user

import (
	"context"

	userrepo "github.com/0B1t322/BWG-Test/internal/domain/user/repository"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	repo userrepo.UserRepository
}

// CreateUser creates new user
// If user with this username already exists, it will return ErrUserExists
func (u *UserServiceImpl) CreateUser(ctx context.Context, username string) (aggregate.User, error) {
	user := aggregate.NewUser(username)
	err := u.repo.CreateUser(ctx, user)
	if err == userrepo.ErrUserExists {
		return aggregate.User{}, ErrUserExists
	} else if err != nil {
		return aggregate.User{}, err
	}

	return user, nil
}

// GetUser return user from the database
// If user not found, it will return ErrUserNotFound
func (u *UserServiceImpl) GetUser(ctx context.Context, userID uuid.UUID) (aggregate.User, error) {
	user, err := u.repo.GetUser(ctx, userID)
	if err == userrepo.ErrUserNotFound {
		return aggregate.User{}, ErrUserNotFound
	} else if err != nil {
		return aggregate.User{}, err
	}

	return user, nil
}
