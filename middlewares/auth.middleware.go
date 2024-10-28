package middlewares

import (
	"fmt"
	controller "notification-api/controllers"
	"notification-api/excepriton"
	"notification-api/service/database"

	"github.com/gofiber/fiber/v2"
)

// AccessToken -> correct "AccessToken" header should come with a request <-*
type AccessToken struct {
	AccessToken string
}

// AuthMiddleware -> compare received access token with a user key in a database
func AuthMiddleware(c *fiber.Ctx) error {

	var err error
	var t = new(AccessToken)
	var content *controller.NotificationRequestDto // <- the main request type for calling /api/v1/

	err = c.BodyParser(&content)
	if err != nil {
		excepriton.HandleAnError("requestDto parser failed with error:", err)
		c.Status(500)
		return err
	}

	// get a company unique access hash str and compare it with the received header
	access, err := database.GetAccessToken(content.DomainName)
	if err != nil {
		return err
	}
	err = c.ReqHeaderParser(t)
	if err != nil {
		excepriton.HandleAnError("request header parsing err: ", err)
		c.Status(417)
		return err
	}

	// fmt.Println("headers ojs => /n", t)
	// fmt.Println("access str is  => ", access)
	// fmt.Println("retrieved header is => ", t.AccessToken)

	if t.AccessToken != access {
		c.Status(403).JSON(fiber.Map{
			"Ok":     false,
			"Reason": "Permission denied",
		})

		fmt.Println("got a wrong auth token from ", content.DomainName)
		return nil
	}

	return c.Next()
}
