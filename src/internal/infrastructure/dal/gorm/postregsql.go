package dal

import (
	"github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToPostgreSQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	// In default set logger mode to Silent
	db.Logger = db.Logger.LogMode(logger.Silent)

	return db, nil
}

func CreateModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Balance{},
		&models.Transaction{},
	)
}
