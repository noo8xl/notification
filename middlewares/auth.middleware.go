package middlewares

import (
	"fmt"
	controller "notification-api/controllers"
	"notification-api/service/database"

	"github.com/gofiber/fiber/v2"
)

// -> right "AccessToken" header should came with request <-*
type AccessToken struct {
	AccessToken string
}

// AuthMiddleware -> compare received access token with user key in database
func AuthMiddleware(c *fiber.Ctx) error {

	var err error
	var content *controller.NotificationRequestDto // <- the main request type for calling /api/v1/

	err = c.BodyParser(&content)
	if err != nil {
		fmt.Println("requestDto parser failed with error:\n", err)
		c.Status(500)
		return err
	}

	var t = new(AccessToken)

	// get company unique access hash str and compare it with received header
	access := database.GetAccessToken(content.DomainName) // --- >> get hash from db, NOT FROM CONFIG
	err = c.ReqHeaderParser(t)

	// fmt.Println("headers ojs => /n", t)
	// fmt.Println("access str is  => ", access)
	// fmt.Println("retrieved header is => ", t.AccessToken)

	if err != nil {
		fmt.Println("auth middleware was failed with error:\n", err.Error())
		return err
	} else {
		if t.AccessToken != access {
			c.Status(403).JSON(fiber.Map{
				"Ok":     false,
				"Reason": "Permission denied",
			})
			fmt.Println("wrong auth token from ", content.DomainName)
			return nil
		} else {
			return c.Next()
		}
	}
}
