package applications

import (
	"context"
	users "dompet/applications/user"
	"dompet/config"
	"dompet/database"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type Module struct {
	collections []mongo.Collection
}

func Init(g *echo.Group) {
	var (
		APP     = config.Config.Server.AppName
		VERSION = config.Config.Server.Version
	)

	g.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	user_collection := db.Collection("users")
	user_service := users.NewService(user_collection, context.TODO())
	users.New(user_service).Route(g.Group("/user"))

}
