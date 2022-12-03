package users

import (
	balancesrv "github.com/0B1t322/BWG-Test/internal/domain/balance/service"
	usersrv "github.com/0B1t322/BWG-Test/internal/domain/user/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	service        usersrv.UserService
	balanceService balancesrv.BalanceService
}

func NewUserController(
	service usersrv.UserService,
	balanceService balancesrv.BalanceService,
) *UserController {
	return &UserController{
		service:        service,
		balanceService: balanceService,
	}
}

func (u UserController) Build(r gin.IRouter) {
	r.GET("/users", u.GetUsers)
	r.GET("/users/:id", u.GetUser)
	r.POST("/users", u.CreateUser)
}

// GetUsers
//
// @Summary Get all users
// @Description Get all users
// @Router /users [get]
// @Tags users
// @Produce json
// @Success 200 {object} UsersView
func (u UserController) GetUsers(c *gin.Context) {
	users, err := u.service.GetUsers(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(
		200,
		UsersViewFrom(users),
	)
}

// GetUser
// @Summary Get user by id
// @Description Get user by id
// @Router /users/{id} [get]
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserView
func (u UserController) GetUser(c *gin.Context) {
	var req GetUserReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := u.service.GetUser(c, id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(
		200,
		UserViewFrom(user),
	)
}

// CreateUser
// @Summary Create user
// @Description Create user
// @Tags users
// @Router /users [post]
// @Accept json
// @Produce json
// @Param userBody body CreateUserReq true "User body"
// @Success 200 {object} UserView
func (u UserController) CreateUser(c *gin.Context) {
	var req CreateUserReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	user, err := u.service.CreateUser(c, req.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	logrus.Info(user.ID)

	b, err := u.balanceService.CreateBalance(
		c,
		user.ID,
	)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Balance = &b

	c.JSON(
		200,
		UserViewFrom(user),
	)
}
