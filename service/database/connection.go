package database

import (
	"context"
	"log"
	"notification-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// databaseName -> available in <database> package
var dbConfigs [2]string = config.GetMONGOdatabaseConfig() // first item -> link, second -> name
var databaseName = dbConfigs[1]

// connectDb -> is connector to mongodb database with required params
func connectDb() *mongo.Client {

	ctx := context.Background()
	mongoPath := dbConfigs[0]
	// dbName := dbConfigs[1] // to use in db ping func <-
	clientOptions := options.Client().ApplyURI(mongoPath)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// // Send a ping to confirm a successful connection
	// if err := client.Database(dbName).RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
	// 	telegram.SendErrorMessage("ping to mongodb was failed")
	// 	return nil
	// }
	// fmt.Println("-> You successfully connected to MongoDB <-")
	return client
}
