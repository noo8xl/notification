package main

import (
	"log"
	controller "notification-api/controllers"
	"notification-api/middlewares"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// var botConfig tgbotapi.UpdateConfig = config.BotConfig()
// var authBot tgbotapi.BotAPI

// func init() {
// 	authBot = *telegram.InitAuthBot()
// }

// api/v1 -> provides only email letters
// and telegram messages
// <-

// api/v2 - > will provides another types of notifications

func main() {

	// ######################### -> init fiber <- ################################
	app := fiber.New()
	app.Use(recover.New())
	// AuthMiddleware -> check client access key
	app.Use("/api/v1/", middlewares.AuthMiddleware)

	// ###################### -> routes list <- ##################################

	// signup for a new company
	app.Post("/notification/api/auth/sign-up/", controller.Registration) // sign a new client
	app.Get("/notification/api/auth/sign-up/confirmation/:code/", controller.ConfirmSignUp)
	app.Get("", controller.RenewAuthKey) // set a new client auth key

	// -> notification handlers
	app.Post("/api/v1/notification/send-user-message/", controller.SendMessage)
	// handle <project> errors (fatal or server errors)
	// chatId (as const) => developer tg chat id
	app.Get("/api/v1/notification/handle-error/:msg/", controller.HandleError)

	// #################### > handle telegram 0Auth <- ###########################
	// go func() {
	// 	fmt.Println("-> auth bot was started.")
	// 	var updates tgbotapi.UpdatesChannel = authBot.GetUpdatesChan(botConfig)
	// 	for update := range updates {
	// 		telegram.HandleUpdates(update, &authBot)
	// 	}
	// }()

	// ##################### > start fiber server <- #############################
	// fmt.Println("-> fiber was started.")
	log.Fatal(app.Listen(":7493"))
}
