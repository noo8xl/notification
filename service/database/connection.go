package database

import (
	"context"
	"notification-api/config"
	"notification-api/excepriton"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type database struct {
	db   *mongo.Client
	name string
}

// connectDb -> is a connector to a mongodb database with required params
func initDatabaseConnection() (*database, error) {

	ctx := context.Background()
	configs := config.GetMONGOdatabaseConfig()
	uri := strings.Join([]string{"mongodb+srv://", configs.User, ":", configs.Password, "@cluster001.sipjs.mongodb.net/?retryWrites=true&w=majority"}, "")
	clientOptions := options.Client().ApplyURI(uri)

	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		excepriton.HandleAnError("db connection got an err: ", err)
		return nil, err
	}
	return &database{db: db, name: configs.Name}, nil
}
