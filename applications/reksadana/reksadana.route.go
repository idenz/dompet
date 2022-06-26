package reksadana

import "github.com/labstack/echo/v4"

func (reksadana *Controller) Route(g *echo.Group) {

	path := g.Group("")
	path.POST("", reksadana.Create)
	path.GET("", reksadana.GetAll)

}
