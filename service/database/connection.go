package database

import (
	"context"
	"notification-api/config"
	"notification-api/excepriton"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseService struct {
	db   *mongo.Client
	name string
}

func InitDatabaseService() *DatabaseService {
	return &DatabaseService{db: nil, name: ""}
}

// connectDb -> is a connector to a mongodb database with required params
func initDatabaseConnection() *DatabaseService {

	ctx := context.Background()
	opts := config.GetMONGOdatabaseConfig()
	// uri := strings.Join([]string{"mongodb+srv://", configs.User, ":", configs.Password, "@cluster001.sipjs.mongodb.net/?retryWrites=true&w=majority"}, "")
	clientOptions := options.Client().ApplyURI(opts[0])

	conn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		excepriton.HandleAnError("db connection got an err: " + err.Error())
	}
	return &DatabaseService{db: conn, name: opts[1]}
}
