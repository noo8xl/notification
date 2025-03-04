package database

import (
	"context"
	"errors"
	"fmt"
	"notification-api/excepriton"
	"notification-api/helpers"
	"notification-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SignNewClient -> prepare and save registration data of a new client
func (s *DatabaseService) SignNewClient(dto *models.ClientRegistrationDto) error {

	s = initDatabaseConnection()
	candidate := new(models.CompanyList)
	clientDetails := new(models.CompanyDetails)
	filter := bson.D{{Key: "companyDomain", Value: dto.DomainName}}
	doc := models.CompanyList{CompanyDomain: dto.DomainName}

	userHashKey := helpers.CreateClientKey(38)
	ctx := context.TODO()

	db := s.db.Database(s.name)
	baseCollection := db.Collection("CompanyList")
	detailsCollection := db.Collection("CompanyDetails")
	defer s.db.Disconnect(ctx)

	cand := baseCollection.FindOne(ctx, filter)
	if cand != nil {
		cand.Decode(&candidate)

		if candidate.CompanyDomain != "" {
			return errors.New("user already exists")
		}
	}

	result, err := baseCollection.InsertOne(ctx, &doc)
	if err != nil {
		excepriton.HandleAnError("db insertion err: " + err.Error())
		return err
	}

	clientDetails.CompanyId = fmt.Sprintf("%v", result.InsertedID)
	clientDetails.DomainName = dto.DomainName
	clientDetails.UserEmail = dto.UserEmail
	clientDetails.JoinDate = time.Now().Format(time.UnixDate)
	clientDetails.UniqueKey = userHashKey

	_, err = detailsCollection.InsertOne(ctx, &clientDetails)
	if err != nil {
		excepriton.HandleAnError("db insertion err: " + err.Error())
		return err
	}
	return nil
}

// GetAccessToken -> get client access token for middleware
func (s *DatabaseService) GetAccessToken(d string) (string, error) {

	s = initDatabaseConnection()
	var result *models.CompanyDetails

	db := s.db.Database(s.name)
	collection := db.Collection("CompanyDetails")
	filter := bson.D{{Key: "domainName", Value: d}}
	ctx := context.TODO()
	defer s.db.Disconnect(ctx)

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			excepriton.HandleAnError("Not found err: " + err.Error())
			return "", err
		}
		excepriton.HandleAnError("GetAccessToken func err: " + err.Error())
		return "", err
	}

	return result.UniqueKey, nil
}

// SaveHistory -> save notification details (s -> service)
func (s *DatabaseService) SaveToTheHistory(item *models.NotificationHistory) error {

	s = initDatabaseConnection()
	db := s.db.Database(s.name)
	collection := db.Collection("CompanyDetails")
	ctx := context.TODO()
	defer s.db.Disconnect(ctx)

	_, err := collection.InsertOne(ctx, &item)
	if err != nil {
		excepriton.HandleAnError("save notification history was failed: " + err.Error())
		return err
	}
	return nil
}

// GetNotificationHistotyList -> get a list of notifications
func (s *DatabaseService) GetNotificationHistotyList(skip int, limit int, recepient string) ([]*models.NotificationHistory, error) {

	return nil, nil
}
