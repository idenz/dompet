package applications

import (
	"context"
	cinvoice "dompet/applications/convetional/invoice"
	cosf "dompet/applications/convetional/osf"
	finances "dompet/applications/finance"
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

	g.GET("", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	/** User Router */
	user_collection := db.Collection("users")
	user_service := users.NewService(user_collection, context.TODO())
	users.New(user_service).Route(g.Group("/user"))
	/** End User Router */

	/** Finance Router */
	finance_collection := db.Collection("finances")
	finance_service := finances.NewService(finance_collection, context.TODO())
	finances.New(finance_service).Route(g.Group("/finance"))
	/** End Finance Router */

	/** Conventional Invoice Router */
	cinvoice_collection := db.Collection("conventional_invoice")
	cinvoice_service := cinvoice.NewService(cinvoice_collection, context.TODO())
	cinvoice.New(cinvoice_service).Route(g.Group("/conventional/invoice"))
	/** End Conventional Invoice Router */

	/** Conventional Invoice Router */
	cosf_collection := db.Collection("conventional_osf")
	cosf_service := cosf.NewService(cosf_collection, context.TODO())
	cosf.New(cosf_service).Route(g.Group("/conventional/osf"))
	/** End Conventional Invoice Router */

}
