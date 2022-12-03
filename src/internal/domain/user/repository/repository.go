package user

import (
	"context"
	"errors"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

var (
	ErrUserExists   = errors.New("User with this username already exists")
	ErrUserNotFound = errors.New("User not found")
)

type UserRepository interface {
	// CreateUser creates new user
	// If user with this username already exists, it will return ErrUserExists
	CreateUser(ctx context.Context, user aggregate.User) error

	// GetUser return user from the database
	// If user not found, it will return ErrUserNotFound
	GetUser(ctx context.Context, id uuid.UUID) (aggregate.User, error)

	// GetUsers return all users from the database
	GetUsers(ctx context.Context) ([]aggregate.User, error)
}
