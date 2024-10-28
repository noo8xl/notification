package controller

import (
	"net/url"
	"notification-api/excepriton"
	"notification-api/models"
	"notification-api/service/database"
	"notification-api/service/telegram"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// the doc is here ->
// -> https://docs.gofiber.io/api/ctx/

// ===============================

// Registration -> sign a new client
func Registration(c *fiber.Ctx) error {

	dto := new(models.ClientRegistrationDto)

	err := c.BodyParser(&dto)
	if err != nil {
		excepriton.HandleAnError("Registration requestDto parser failed with error:", err)
		c.Status(500)
		return err
	}

	err = database.SignNewClient(dto)
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
	dto := new(NotificationRequestDto)

	err = c.BodyParser(dto)
	if err != nil {
		excepriton.HandleAnError("SendMessage bodyparser got an error: ", err)
		c.Status(417)
		return err
	}

	switch dto.ServiceType {
	case "telegram":
		err = dto.SendTelegramMessage()
		if err != nil {
			c.Status(417)
			return err
		}
	case "email":
		err = dto.SendEmailMessage()
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
		excepriton.HandleAnError("Received wrong service type.", err)
		c.Status(400)
		return err
	}

	err = dto.SaveToTheHistory()
	if err != nil {
		c.Status(417)
		return err
	}

	c.Status(200)
	return nil
}

// HandleError -> handle all critical project errors and send msg to developer
func HandleError(c *fiber.Ctx) error {
	ctx := c.AllParams()["msg"]

	decodedMsg, err := url.QueryUnescape(ctx)
	if err != nil {
		excepriton.HandleAnError("Error decoding URL:", err)
		c.Status(417)
		return err
	}

	err = telegram.SendErrorMessage(decodedMsg)
	if err != nil {
		c.Status(417)
		return err
	}

	c.Status(200)
	return nil
}

// HandleError -> handle all critical project errors and send msg to developer
func GetHistoryList(c *fiber.Ctx) error {
	params := c.AllParams()

	skip, _ := strconv.Atoi(params["skip"])
	lim, _ := strconv.Atoi(params["limit"])
	recipient := params["recipient"]

	list, err := database.GetNotificationHistotyList(skip, lim, recipient)
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
