package user

import (
	"context"
	"errors"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("User not found")
	ErrUserExists   = errors.New("User with this username already exists")
)

type UserService interface {
	// CreateUser creates new user
	// If user with this username already exists, it will return ErrUserExists
	CreateUser(ctx context.Context, username string) (aggregate.User, error)

	// GetUser return user from the database
	// If user not found, it will return ErrUserNotFound
	GetUser(ctx context.Context, userID uuid.UUID) (aggregate.User, error)
}
