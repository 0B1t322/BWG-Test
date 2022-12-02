package dal_test

import (
	"testing"

	"github.com/0B1t322/BWG-Test/internal/config"
	dal "github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm"
	"github.com/stretchr/testify/require"
)

func TestFunc_ConnectToDb(t *testing.T) {
	db, err := dal.ConnectToPostgreSQL(config.GlobalConfig.DB.PostgreSQLDSN)
	require.NoError(t, err)

	err = dal.CreateModels(db)
	require.NoError(t, err)
}
