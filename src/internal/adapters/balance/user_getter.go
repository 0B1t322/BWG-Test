package balance

import (
	"context"

	balancesrv "github.com/0B1t322/BWG-Test/internal/domain/balance/service"
	usersrv "github.com/0B1t322/BWG-Test/internal/domain/user/service"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
)

type UserGetter struct {
	userService usersrv.UserService
}

func NewUserGetter(
	userService usersrv.UserService,
) balancesrv.UserGetter {
	return &UserGetter{
		userService: userService,
	}
}

func (u UserGetter) GetUser(ctx context.Context, userID uuid.UUID) (aggregate.User, error) {
	user, err := u.userService.GetUser(ctx, userID)
	if err == usersrv.ErrUserNotFound {
		return aggregate.User{}, balancesrv.ErrUserNotFound
	} else if err != nil {
		return aggregate.User{}, err
	}

	return user, nil
}
