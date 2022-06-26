package users

import (
	"encoding/json"
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

	payload := new(UserRegisterRequest)
	if err = c.Bind(payload); err != nil {
		return err
	}

	if err = c.Validate(payload); err != nil {
		return err
	}

	user, err := u.UserService.CreateUser(payload)
	if err != nil {
		return nil
	}

	result := new(UserResponse)
	bytes, err := json.Marshal(&user)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, result)
}

func (u *UserController) GetUser(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, "dennt")
}
