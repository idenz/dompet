package users

import "github.com/labstack/echo/v4"

func (user *UserController) Route(g *echo.Group) {

	path := g.Group("")
	path.POST("/register", user.CreateUser)
	path.GET("/profile", user.GetUser)

}
