package authentication

import (
	users "dompet/applications/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserService users.IService
	Service     IService
}

/** Constructors */
func New(_userService users.IService, _service IService) *Controller {
	return &Controller{
		UserService: _userService,
		Service:     _service,
	}
}

func (c *Controller) Login(ctx echo.Context) (err error) {

	payload := new(AuthLoginRequest)
	if err = ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, &ResponseError{
			Code:   http.StatusBadRequest,
			Status: "Can't find your credential",
		})
	}

	if err = ctx.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, &ResponseError{
			Code:   http.StatusBadRequest,
			Status: "Validation Failed",
		})
	}

	token, err := c.Service.Login(*payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &ResponseError{
			Code:   http.StatusBadRequest,
			Status: "Can't find your credential",
		})
	}

	result := &ResponseLogin{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   *token,
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *Controller) Register(ctx echo.Context) (err error) {

	payload := new(users.RegisterRequest)
	if err = ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, &ResponseLogin{
			Code:   http.StatusBadRequest,
			Status: "Failed create user",
		})
	}

	if err = ctx.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, &ResponseLogin{
			Code:   http.StatusBadRequest,
			Status: "Validation Failed",
		})
	}

	user, err := c.UserService.Create(payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &ResponseLogin{
			Code:   http.StatusBadRequest,
			Status: "Failed create user",
		})
	}

	result := &ResponseRegister{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   *user,
	}

	return ctx.JSON(http.StatusOK, result)
}
