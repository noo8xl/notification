package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// #################################################################################
// ####################### -> internal function only <- ############################
// #################################################################################

// insertData -> insert data to db
func insertData(col string, doc any) string {

	var id string
	client := connectDb()
	db := client.Database(databaseName)
	ctx := context.TODO()
	collection := db.Collection(col)

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Printf("Disconnect db catch an error %s\n", err.Error())
		}
	}(client, ctx)

	result, err := collection.InsertOne(ctx, &doc)
	if err != nil {
		fmt.Println(err)
		return id
	} else {
		id = fmt.Sprintf("%v", result.InsertedID)
		//fmt.Println("ins ID -> ", id)
		fmt.Println("Insertion was done successfully.")
		return id
	}
}

// // updateData -> update received data in db
// func updateData(col string, doc any, filter primitive.D) bool {
// 	return true
// }

// isDbContains -> check if data already exists in db
func isDbContains(col string, filter primitive.D) bool {

	var result any // should update <-
	ctx := context.TODO()
	client := connectDb()
	db := client.Database(databaseName)
	collection := db.Collection(col)

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Printf("Disconnect db catch an error %s\n", err.Error())
		}
	}(client, ctx)

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("isDbContains result is =>\n", result)
	return true
}
