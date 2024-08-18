package database

import (
	"context"
	"fmt"
	"notification-api/helpers"
	"notification-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// var databaseName [2]string = config.GetMONGOdatabaseConfig() // first elem -> link, second -> name

// SignNewClient -> prepare data and then call insert func to registrate new client
func SignNewClient(dto *models.ClietRegistratioDto) bool {

	filter := bson.D{{Key: "companyDomain", Value: dto.DomainName}}
	if candidate := isDbContains("CompanyList", filter); candidate {
		fmt.Println("Client already exists.")
		return false
	}
	var clietnId string
	clientDetails := new(models.CompanyDetails)
	userHashKey := helpers.CreateClientKey(38)
	doc := models.CompanyList{CompanyDomain: dto.DomainName}

	clietnId = insertData("CompanyList", doc) // ->  save to list and get an id
	clientDetails.CompanyId = clietnId
	clientDetails.DomainName = dto.DomainName
	clientDetails.UserEmail = dto.UserEmail
	clientDetails.JoinDate = time.Now().Format(time.UnixDate)
	clientDetails.UniqueKey = userHashKey

	s := insertData("CompanyDetails", clientDetails)
	fmt.Println("details id is => ", s)
	return true
}

// GetAccessToken -> get client access token for middleware
func GetAccessToken(d *string) *string {

	var result *models.CompanyDetails
	client := connectDb()
	db := client.Database(databaseName)
	collection := db.Collection("CompanyDetails")
	filter := bson.D{{Key: "domainName", Value: d}}
	ctx := context.TODO()

	defer client.Disconnect(ctx)

	cursor := collection.FindOne(ctx, filter)
	cursor.Decode(&result)

	return &result.UniqueKey
}

// SaveHistory -> save sendet notification details
func SaveHistory(item *models.NotificationHistory) {
	s := insertData("NotificationHistory", item)
	fmt.Println("NotificationHistory id is => ", s)
}
