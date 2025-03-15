package controller

import (
	"log"
	"net/url"
	"notification-api/excepriton"
	"notification-api/models"
	"notification-api/service/database"
	"notification-api/service/email"
	"notification-api/service/telegram"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// the doc is here ->
// -> https://docs.gofiber.io/api/ctx/

// ===============================

type AuthService struct {
	db *database.DatabaseService
}

type NotificationService struct {
	db *database.DatabaseService
	tg *telegram.TelegramService
	e  *email.EmailService
}

func InitAuthService() *AuthService {
	return &AuthService{
		db: database.InitDatabaseService(),
	}
}

func InitNotificationService() *NotificationService {
	return &NotificationService{
		db: database.InitDatabaseService(),
		tg: telegram.NewTelegramService(),
		e:  email.NewEmailService(),
	}
}

func InitController() {

}

// Registration -> sign a new client
func (a *AuthService) Registration(c *fiber.Ctx) error {

	dto := new(models.ClientRegistrationDto)

	err := c.BodyParser(&dto)
	if err != nil {
		excepriton.HandleAnError("Registration requestDto parser failed with error:" + err.Error())
		c.Status(500)
		return err
	}

	err = a.db.SignNewClient(dto)
	if err != nil {
		if err.Error() == "user already exists" {
			c.Status(400).JSON(fiber.Map{
				"Ok":      false,
				"Message": err.Error(),
			})
			return nil
		} else {
			c.Status(417)
		}
	}

	// send email with auth hash here ?*
	// or in service
	c.Status(201)
	return nil
}

// ConfirmSignUp -> confirm registration via client email
func (a *AuthService) ConfirmSignUp(c *fiber.Ctx) error {
	return nil
}

// RenewAuthKey -> generate a new client auth key and send it via email
func (a *AuthService) RenewAuthKey(c *fiber.Ctx) error {
	// date, email
	return nil
}

// SendMessage -> handle a request from the API and send a notification via chosen service
func (n *NotificationService) SendMessage(c *fiber.Ctx) error {
	var err error
	dto := new(models.NotificationRequestDto)

	err = c.BodyParser(dto)
	if err != nil {
		excepriton.HandleAnError("SendMessage bodyparser got an error: " + err.Error())
		c.Status(417)
		return err
	}

	switch dto.ServiceType {
	case "telegram":
		d := &models.SendTwoStepCodeDto{
			ChatID:  dto.Recipient,
			Message: dto.Content,
		}
		err = n.tg.SendUserMessage(d)
		if err != nil {
			c.Status(417)
			return err
		}
	case "email":
		d := &models.EmailDto{
			ServiceType: dto.ServiceType,
			DomainName:  dto.DomainName,
			Content:     dto.Content,
			Recipient:   dto.Recipient,
		}
		err = n.e.PrepareEmailMessage(d)
		if err != nil {
			c.Status(417)
			return err
		}
	// case "slack":
	// 	err = dto.SendEmailMessage()
	// 	if err != nil {
	// 	c.Status(417)
	// 	return err
	// }
	default:
		excepriton.HandleAnError("Received wrong service type.")
		c.Status(400)
		return err
	}
	historyItem := &models.NotificationHistory{
		DateTime:    time.Now().Format(time.UnixDate),
		Recipient:   dto.Recipient,
		DomainName:  dto.DomainName,
		MessageBody: dto.Content,
		SentVia:     dto.ServiceType,
	}
	err = n.db.SaveToTheHistory(historyItem)
	if err != nil {
		c.Status(417)
		return err
	}

	c.Status(200)
	return nil
}

// HandleError -> handle all critical project errors and send msg to developer
func (n *NotificationService) HandleError(c *fiber.Ctx) error {
	log.Println("log1")
	ctx := c.AllParams()["msg"]

	decodedMsg, err := url.QueryUnescape(ctx)
	if err != nil {
		excepriton.HandleAnError("Error decoding URL:" + err.Error())
		c.Status(417)
		return err
	}

	err = n.tg.SendErrorMessage(decodedMsg)
	if err != nil {
		c.Status(417)
		return err
	}

	c.Status(200)
	return nil
}

// HandleError -> handle all critical project errors and send msg to developer
func (n *NotificationService) GetHistoryList(c *fiber.Ctx) error {
	params := c.AllParams()

	skip, _ := strconv.Atoi(params["skip"])
	lim, _ := strconv.Atoi(params["limit"])
	recipient := params["recipient"]

	list, err := n.db.GetNotificationHistotyList(skip, lim, recipient)
	if err != nil {
		c.Status(417)
		return err
	}

	c.Status(200).JSON(fiber.Map{
		"Ok":   true,
		"List": list,
	})
	return nil

}
