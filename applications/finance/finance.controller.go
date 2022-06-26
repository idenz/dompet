package finances

import "github.com/labstack/echo/v4"

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
	return nil
}

func (c *Controller) GetAll(ctx echo.Context) (err error) {
	return nil
}
