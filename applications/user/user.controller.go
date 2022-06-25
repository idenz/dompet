package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService IUserService
}

/** Constructors */
func New(_userService IUserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (u *UserController) CreateUser(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, "denny")
}

func (u *UserController) GetUser(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, "dennt")
}
