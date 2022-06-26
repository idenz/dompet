package authentication

import (
	users "dompet/applications/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Service users.IService
}

/** Constructors */
func New(_service users.IService) *Controller {
	return &Controller{
		Service: _service,
	}
}

func (c *Controller) Register(ctx echo.Context) (err error) {

	payload := new(users.RegisterRequest)
	if err = ctx.Bind(payload); err != nil {
		return err
	}

	if err = ctx.Validate(payload); err != nil {
		return err
	}

	user, err := c.Service.Create(payload)
	if err != nil {
		return err
	}

	result := &ResponseRegister{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   *user,
	}

	return ctx.JSON(http.StatusOK, result)
}
