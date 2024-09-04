package controller

import (
	"fmt"
	"net/url"
	"notification-api/models"
	"notification-api/service/database"
	"notification-api/service/email"
	"notification-api/service/telegram"
	"time"

	"github.com/gofiber/fiber/v2"
)

// the doc is here ->
// -> https://docs.gofiber.io/api/ctx/

// ===============================

// Registration -> sign a new client
func Registration(c *fiber.Ctx) error {
	// generate a unique hash str by dto
	var err error
	var statusCode int = 201
	dto := new(models.ClientRegistrationDto)

	err = c.BodyParser(dto)
	if err != nil {
		fmt.Println("Registration requestDto parser failed with error:\n", err)
		c.Status(500)
		return err
	}

	resp := database.SignNewClient(*dto)
	if !resp {
		statusCode = 400
	}
	// send email with auth hash here ?*

	c.Status(statusCode)
	return err
}

// ConfirmSignUp -> confirm registration via client email
func ConfirmSignUp(c *fiber.Ctx) error {
	return nil
}

// RenewAuthKey -> generate a new client auth key and send it via email
func RenewAuthKey(c *fiber.Ctx) error {
	// date, email
	return nil
}

// SendMessage -> handle a request from the API and send a notification via chosen service
func SendMessage(c *fiber.Ctx) error {
	var err error
	var status int = 200
	dto := new(NotificationRequestDto)

	err = c.BodyParser(dto)
	fmt.Println("dto is ->\n", dto)
	if err != nil {
		fmt.Println("SendMessage requestDto parser failed with error:\n", err.Error())
		c.Status(500)
		return err
	}

	fmt.Println("dto is ->\n", dto)

	switch dto.ServiceType {
	case "telegram":
		err = dto.SendTelegramMessage()
	case "email":
		err = dto.SendEmailMessage()
	default:
		fmt.Println("Received wrong service type.")
		status = 400
	}

	if err != nil {
		status = 500
	}

	dto.SaveHistory()
	c.Status(status)
	return err
}

// HandleError -> handle all critical project errors and send msg to developer
func HandleError(c *fiber.Ctx) error {

	ctx := c.AllParams()["msg"]

	decodedMsg, err := url.QueryUnescape(ctx)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		c.Status(400)
		return err
	}

	telegram.SendErrorMessage(decodedMsg)

	c.Status(200)
	return nil
}

// ##################################################################################
// ##################################################################################
// ##################################################################################

// SendTelegramMessage -> handle only a telegram notification sending
func (dto NotificationRequestDto) SendTelegramMessage() error {
	d := models.SendTwoStepCodeDto{
		ChatID:  dto.Recipient,
		Message: dto.Content,
	}
	return telegram.SendUserMessage(d)
}

// SendEmailMessage - > handle only an email notification sending
func (dto NotificationRequestDto) SendEmailMessage() error {
	d := models.EmailDto{
		ServiceType: dto.ServiceType,
		DomainName:  dto.DomainName,
		Content:     dto.Content,
		Recipient:   dto.Recipient,
	}
	return email.PrepareEmailMessage(d)
}

func (dto NotificationRequestDto) SaveHistory() {
	historyItem := models.NotificationHistory{
		DateTime:    time.Now().Format(time.UnixDate),
		Recipient:   dto.Recipient,
		DomainName:  dto.DomainName,
		MessageBody: dto.Content,
		SentVia:     dto.ServiceType,
	}
	// fmt.Println("hist ->\n", historyItem)
	database.SaveHistory(historyItem)
}
