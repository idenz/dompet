package users

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Service IService
}

/** Constructors */
func New(_userService IService) *Controller {
	return &Controller{
		Service: _userService,
	}
}

func (u *Controller) CreateUser(c echo.Context) (err error) {

	payload := new(RegisterRequest)
	if err = c.Bind(payload); err != nil {
		return err
	}

	if err = c.Validate(payload); err != nil {
		return err
	}

	user, err := u.Service.Create(payload)
	if err != nil {
		return nil
	}

	result := new(Response)
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

func (u *Controller) GetUser(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, "dennt")
}
