package views

import (
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/samber/lo"
)

type (
	UserView struct {
		ID       string       `json:"id"       format:"uuid"`
		Username string       `json:"username"`
		Balance  *BalanceView `json:"balance"`
	}

	UsersView struct {
		Users []UserView `json:"users"`
	}
)

func UserViewFrom(user aggregate.User) UserView {
	return UserView{
		ID:       user.ID.String(),
		Username: user.Username,
		Balance: lo.Ternary(
			user.Balance != nil,
			lo.ToPtr(BalanceViewFrom(*user.Balance)),
			nil,
		),
	}
}

func UsersViewFrom(users []aggregate.User) UsersView {
	usersView := UsersView{
		Users: make([]UserView, len(users)),
	}

	for i, user := range users {
		usersView.Users[i] = UserViewFrom(user)
	}

	return usersView
}
