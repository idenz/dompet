package main

import (
	"dompet/config"
	"dompet/helper"
	"dompet/utils"
	"fmt"

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

	/**
	 * 1. Check the string is palindrome or not
	 */
	palindrome, value := helper.IsPallindrome("abbccd")
	fmt.Println("1. Palindrome of " + "'" + value + "'" + " is " + fmt.Sprintf("%v", palindrome))

	/**
	 * 2. Find prime number by range
	 */
	fmt.Println("2. " + fmt.Sprintf("%v", helper.FindPrimeByRange(11, 40)))

	app := echo.New()
	app.Validator = &utils.CustomValidator{Validator: validator.New()}

	app.Logger.Fatal(app.Start(":" + PORT))
}
