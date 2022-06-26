package finances

import (
	"encoding/json"
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

	finance, err := c.Service.Create(payload)
	if err != nil {
		return err
	}

	result := new(Response)
	bytes, err := json.Marshal(&finance)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *Controller) GetAll(ctx echo.Context) (err error) {
	return ctx.JSON(http.StatusOK, "ok")
}
