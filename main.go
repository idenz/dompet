package main

import (
	"dompet/config"
	"dompet/utils"

	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {

	err := config.Load("")
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

}

func main() {
	var (
		PORT = config.Config.Server.Port
	)

	app := echo.New()
	app.Validator = &utils.CustomValidator{Validator: validator.New()}

	app.Logger.Fatal(app.Start(":" + PORT))
}
