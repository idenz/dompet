package database

import (
	"context"
	"dompet/config"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() (*mongo.Database, error) {
	var (
		USERNAME = config.Config.Database.DBUser
		PASSWORD = config.Config.Database.DBPass
		HOST     = config.Config.Database.DBHost
		PORT     = config.Config.Database.DBPort
		URI      = fmt.Sprint("mongodb://" + HOST + ":" + PORT)
		DBNAME   = config.Config.Database.DBName
	)

	credential := options.Credential{
		Username: USERNAME,
		Password: PASSWORD,
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(URI)
	if USERNAME != "" && PASSWORD != "" {
		clientOptions.SetAuth(credential)
	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client.Database(DBNAME), nil
}
