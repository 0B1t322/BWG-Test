package users

import (
	"github.com/0B1t322/BWG-Test/internal/controllers/http/shared/views"
	"github.com/0B1t322/BWG-Test/internal/controllers/http/users/dto"
)

// Alias dto
type (
	GetUserReq    = dto.GetUserReq
	CreateUserReq = dto.CreateUserReq
)

// Alias Views
type (
	UserView  = views.UserView
	UsersView = views.UsersView
)

// Alias Mappers
var (
	UserViewFrom  = views.UserViewFrom
	UsersViewFrom = views.UsersViewFrom
)
