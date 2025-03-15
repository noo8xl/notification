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

type appController struct {
	n controller.NotificationService
	a controller.AuthService
}

func InitController() *appController {
	return &appController{
		n: *controller.InitNotificationService(),
		a: *controller.InitAuthService(),
	}
}

func main() {
	app := fiber.New()
	app.Use(recover.New())
	// app.Use("/api/v1/", middlewares.AuthMiddleware)
	svc := InitController()

	// sign a new company
	app.Post("/notification/api/auth/sign-up/", svc.a.Registration) // sign a new client
	app.Get("/notification/api/auth/sign-up/confirmation/:code/", svc.a.ConfirmSignUp)
	app.Get("", svc.a.RenewAuthKey) // set a new client auth key

	// -> notification handlers
	app.Post("/api/v1/notification/send-user-message/", svc.n.SendMessage)
	// handle <project> errors (fatal or server errors)
	// chatId (as const) => developer tg chat id
	app.Get("/api/v1/notification/handle-error/:msg/", svc.n.HandleError)
	// -> get a list of sent notif by recipient string value
	app.Get("/api/v1/notification/get-history/:skip/:limit/:recipient/", svc.n.GetHistoryList)

	// svc.httpRoutesHandlerRegister(app)

	// // #################### > handle telegram 0Auth <- ###########################
	// // go func() {
	// // 	fmt.Println("-> auth bot was started.")
	// // 	var updates tgbotapi.UpdatesChannel = authBot.GetUpdatesChan(botConfig)
	// // 	for update := range updates {
	// // 		telegram.HandleUpdates(update, &authBot)
	// // 	}
	// // }()

	// ##################### > start the server <- #############################
	err := app.Listen(":7493")
	if err != nil {
		log.Fatal("Can't start the server: ", err)
	}
}
