package conventional_osf

import "github.com/labstack/echo/v4"

func (cosf *Controller) Route(g *echo.Group) {

	path := g.Group("")
	path.POST("", cosf.Create)
	path.GET("", cosf.GetAll)

}
