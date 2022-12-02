package user_test

import (
	"context"
	"testing"

	"github.com/0B1t322/BWG-Test/internal/config"
	dal "github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm"
	"github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm/user"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/stretchr/testify/require"
)

// TODO: add delete users after tests
func TestFunc_PGUserRepository(t *testing.T) {
	db, err := dal.ConnectToPostgreSQL(config.GlobalConfig.DB.PostgreSQLDSN)
	require.NoError(t, err)

	repo := user.NewPGUserRepository(db)

	t.Run(
		"Create",
		func(t *testing.T) {
			t.Run(
				"Success",
				func(t *testing.T) {
					u := aggregate.NewUser("test")
					err := repo.CreateUser(
						context.Background(),
						u,
					)
					require.NoError(t, err)
				},
			)
		},
	)

	t.Run(
		"Get",
		func(t *testing.T) {
			t.Run(
				"Success",
				func(t *testing.T) {
					u := aggregate.NewUser("test_2")
					err := repo.CreateUser(
						context.Background(),
						u,
					)
					require.NoError(t, err)

					get, err := repo.GetUser(
						context.Background(),
						u.ID,
					)
					require.NoError(t, err)

					require.Equal(t, u, get)
				},
			)
		},
	)
}
