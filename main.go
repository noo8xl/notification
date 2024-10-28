package main

import (
	"log"
	controller "notification-api/controllers"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// var botConfig tgbotapi.UpdateConfig = config.BotConfig()
// var authBot tgbotapi.BotAPI

// func init() {
// 	authBot = *telegram.InitAuthBot()
// }

// api/v1 -> provides only email letters
// and telegram messages.

// api/v2 - > will provide another type of notifications

func main() {

	// ######################### -> init fiber <- ################################
	app := fiber.New()
	app.Use(recover.New())
	// AuthMiddleware -> check a client access key
	// app.Use("/api/v1/", middlewares.AuthMiddleware)

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
	// -> get a list of sent notif by recipient string value
	app.Get("/api/v1/notification/get-history/:skip/:limit/:recipient/", controller.GetHistoryList)

	// #################### > handle telegram 0Auth <- ###########################
	// go func() {
	// 	fmt.Println("-> auth bot was started.")
	// 	var updates tgbotapi.UpdatesChannel = authBot.GetUpdatesChan(botConfig)
	// 	for update := range updates {
	// 		telegram.HandleUpdates(update, &authBot)
	// 	}
	// }()

	// ##################### > start fiber server <- #############################
	err := app.Listen(":7493")
	if err != nil {
		log.Fatal("Can't start the server: ", err)
	}
}
