package sbn

import "github.com/labstack/echo/v4"

func (sbn *Controller) Route(g *echo.Group) {

	path := g.Group("")
	path.POST("", sbn.Create)
	path.GET("", sbn.GetAll)

}
