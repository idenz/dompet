package productive_invoice

import "github.com/labstack/echo/v4"

func (pinvoice *Controller) Route(g *echo.Group) {

	path := g.Group("")
	path.POST("", pinvoice.Create)
	path.GET("", pinvoice.GetAll)

}
