package productive_invoice

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
		return ctx.JSON(http.StatusBadRequest, &ResponseError{
			Code:   http.StatusBadRequest,
			Status: "Failed create data",
		})
	}

	if err = ctx.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, &ResponseError{
			Code:   http.StatusBadRequest,
			Status: "Validation Failed",
		})
	}

	pinvoice, err := c.Service.Create(payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &ResponseError{
			Code:   http.StatusBadRequest,
			Status: "Failed create data",
		})
	}

	result := &Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   *pinvoice,
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *Controller) GetAll(ctx echo.Context) (err error) {

	pinvoices, err := c.Service.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &ResponseError{
			Code:   http.StatusBadRequest,
			Status: "Internal Server Error",
		})
	}

	result := &ResponseMany{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   pinvoices,
	}

	return ctx.JSON(http.StatusOK, result)
}
