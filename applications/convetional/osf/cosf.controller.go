package conventional_osf

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Service IService
}

/** Constructors */
func New(_service IService) *Controller {
	return &Controller{
		Service: _service,
	}
}

func (c *Controller) Create(ctx echo.Context) (err error) {

	payload := new(CreateRequest)
	if err = ctx.Bind(payload); err != nil {
		return err
	}

	if err = ctx.Validate(payload); err != nil {
		return err
	}

	cosf, err := c.Service.Create(payload)
	if err != nil {
		return err
	}

	result := &Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   *cosf,
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *Controller) GetAll(ctx echo.Context) (err error) {

	cosfs, err := c.Service.GetAll()
	if err != nil {
		return err
	}

	result := &ResponseMany{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   cosfs,
	}

	return ctx.JSON(http.StatusOK, result)
}
