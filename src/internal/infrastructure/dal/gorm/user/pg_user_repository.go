package user

import (
	"context"

	user "github.com/0B1t322/BWG-Test/internal/domain/user/repository"
	"github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm/models"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type PGUserRepository struct {
	db *gorm.DB
}

func NewPGUserRepository(db *gorm.DB) *PGUserRepository {
	return &PGUserRepository{db: db}
}

// CreateUser creates new user
// If user with this username already exists, it will return ErrUserExists
func (u *PGUserRepository) CreateUser(ctx context.Context, user aggregate.User) error {
	model := models.UserModelFrom(user)

	result := u.db.WithContext(ctx).Model(&models.User{}).Create(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetUser return user from the database
// If user not found, it will return ErrUserNotFound
func (u *PGUserRepository) GetUser(ctx context.Context, id uuid.UUID) (aggregate.User, error) {
	var model models.User
	{
		result := u.db.WithContext(ctx).
			Model(&models.User{}).
			Where(models.UserFieldID.WithTable()+" = ?", id).
			Joins(models.UserEdgeBalance.String()).
			First(&model)

		if result.Error == gorm.ErrRecordNotFound {
			return aggregate.User{}, user.ErrUserNotFound
		} else if result.Error != nil {
			return aggregate.User{}, result.Error
		}
	}

	return models.UserModelTo(&model), nil
}

// GetUsers return all users from the database
func (u *PGUserRepository) GetUsers(ctx context.Context) ([]aggregate.User, error) {
	var users []models.User
	{
		result := u.db.WithContext(ctx).
			Model(&models.User{}).
			Joins(models.UserEdgeBalance.String()).
			Find(&users)

		if result.Error != nil {
			return nil, result.Error
		}
	}

	return lo.Map(
		users,
		func(user models.User, _ int) aggregate.User {
			return models.UserModelTo(&user)
		},
	), nil
}

func isDuplacateKeyError(err error) bool {
	return err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\""
}
