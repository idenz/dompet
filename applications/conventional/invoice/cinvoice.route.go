package conventional_invoice

import "github.com/labstack/echo/v4"

func (cinvoice *Controller) Route(g *echo.Group) {

	path := g.Group("")
	path.POST("", cinvoice.Create)
	path.GET("", cinvoice.GetAll)

}
