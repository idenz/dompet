package authentication

import "github.com/labstack/echo/v4"

func (auth *Controller) Route(g *echo.Group) {

	path := g.Group("")
	path.POST("/register", auth.Register)
	path.POST("/login", auth.Login)

}
