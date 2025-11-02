package database

import (
	"context"
	"notification-api/config"
	"notification-api/pkg/exceptions"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseService struct {
	db   *mongo.Client
	name string
}

func InitDatabaseService() *DatabaseService {
	return initDatabaseConnection()
}

// connectDb -> is a connector to a mongodb database with required params
func initDatabaseConnection() *DatabaseService {

	ctx := context.Background()
	opts := config.GetMONGOdatabaseConfig()
	// uri := strings.Join([]string{"mongodb+srv://", configs.User, ":", configs.Password, "@cluster001.sipjs.mongodb.net/?retryWrites=true&w=majority"}, "")
	clientOptions := options.Client().ApplyURI(opts[0])

	conn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		exceptions.HandleAnError("db connection got an err: " + err.Error())
	}
	return &DatabaseService{db: conn, name: opts[1]}
}
