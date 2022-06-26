package finances

import "github.com/labstack/echo/v4"

func (finance *Controller) Route(g *echo.Group) {

	path := g.Group("")
	path.POST("/", finance.Create)
	path.GET("/", finance.GetAll)

}
