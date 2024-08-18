package telegram

import (
	"fmt"
	"log"
	"notification-api/config"
	"notification-api/models"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// startMessageHandler -> handle only </start> user message
func startMessageHandler(dto *models.CommandsDto) {
	txt := strings.Join([]string{
		"Hi, ", dto.UserName, "! ",
		"Welcome to AuthBot!", "should add some text <-"}, "")
	msg := tgbotapi.NewMessage(dto.ChatId, txt)
	dto.Bot.Send(msg)
}

// helpMessageHandler -> handle only </help> user message
func helpMessageHandler(dto *models.CommandsDto) {

	greetings := "Let's start to teach you how to interact with me!"
	addArea := "_"
	editArea := "_"
	lessonArea := "_ will update soon"

	txt := strings.Join([]string{greetings, addArea, editArea, lessonArea}, "\n")
	msg := tgbotapi.NewMessage(dto.ChatId, txt)
	dto.Bot.Send(msg)
}

// authHandler -> handle auth messages from <sign in via telegram>
func authHandler(update *tgbotapi.Update) bool {

	fmt.Println("u ->\n", update)

	// var userDto any
	// userDto.UserID = update.Message.From.ID
	// userDto.UserEmail = update.Message.PassportData.Data[0].Email
	// userDto.UserPasswordOrHash = ""
	// userDto.UserName = update.Message.From.UserName
	// userDto.FirstName = update.Message.From.FirstName
	// userDto.LastName = update.Message.From.LastName

	// fmt.Println("cur dto obj is -> ", userDto)
	// send userDto to the client API to sign up new user <-
	return false
	// return clientapi.SendSignUpData(userDto)
}

// defaultMessageHandler -> send default message to unknown message
func defaultMessageHandler(chatId int64, bt *tgbotapi.BotAPI) {
	txt := "AuthBot is only for the auth and don't handle any user messages. If you have some questions please contact support via profile->support!"
	msg := tgbotapi.NewMessage(chatId, txt)
	bt.Send(msg)
}

// #######################################################################
// ###################### -> connect area <- #############################
// #######################################################################

// initErrorBot -> init tg bot for send ERRORs
func initErrorBot() *tgbotapi.BotAPI {
	var err error
	var t string = config.GetErrorHandlerToken()

	// fmt.Println("t => ", t)
	bot, err := tgbotapi.NewBotAPI(t)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	// fmt.Println(bot.Self.UserName)
	// bot.Debug = true
	return bot
}

// InitNotificationBot -> init tg bot for notification
func initNotificationBot() *tgbotapi.BotAPI {
	token := config.GetNitificationBotToken()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println("bon init error >", err)
		os.Exit(1)
	}
	// bot.Debug = true
	return bot
}

// #######################################################################################################
// ######################### update response example from the auth bot ###################################
// #######################################################################################################

// response: {
// "ok":true,
// "result":[
// 	{
// 		"update_id":550215773,
// 		"message": {
// 			"message_id":15,
// 			"from":{
// 				"id":949347540,
// 				"is_bot":false,
// 				"first_name":"John",
// 				"last_name":"Falstaff",
// 				"username":"falstaff_john",
// 				"language_code":"en"
// 				},
// 			"chat":{
// 				"id":949347540,
// 				"first_name":"John",
// 				"last_name":"Falstaff",
// 				"username":"falstaff_john",
// 				"type":"private"
// 				},
// 			"date":1708887311,
// 			"passport_data":{
// 				"data":[
// 					{
// 						"type":"email",
// 						"email":"niko8elich@gmail.com",
// 						"hash":"P1rpp7fUma7Y0kdkvVUyL7TKJhy04bvofWYFqRC6twA="
// 						}
// 					],
// 				"credentials":{
// 					"data":"ahamPJF5HcEWfLYGa12fSixjGkpiMwZcQufWEqHqhK4WrFcCwatx1QVqnQ2AbfvHaOL1q34FUVHlxb1uvDr4If54Swrspdi7o+1456PWWynP5fjFgXCJNNuErGxdc8YnGbJeLtrjGJ+0x603pvwsYwaIVnNoRGAjfqDw/F8wefO1MoqVTLxE/y5KQ/5FNBq72dPVlloRgCBh2i033K4avgBa9AXm8Y11sMLrQBNlfzA=",
// 					"hash":"ahXDHU91x6qgD+w+iguuJXXmvle3ImWXhZFg2NjTCVA=",
// 					"secret":"obtNBTNo++PiqXeDARqGLhfnHTBkm1Kyet7EF+/b7Pkfi4eri4DarQzHem+pskN5h/YcwQiDXUdE5f8dbQX4waANAa9hvxYtmi63xwLnLt1fuSHCvORD//knCBq5BXd3FsTHcySMhKemr+e/6/uqtDSkux7SwI+LeFwe0R4DZrgH1gdrwIg+0cF4Xn6GNglIqzeBWY9mS60cWEQXJCh75sjS9AYvu88ZzL4T26jpDp2H08ZBeCASbJvD302ZgRhWVcb5TNZXVU/KRS9xDGGFJxgRJwl1zPrxlqiebrsMYG3xUN/+rCQB2EzhgnJyiIK3bbI5J1+kr18ShTp6iGy4JA=="
// 					}
// 				}
// 			}
// 		}
// 	]
// }

// #######################################################################################################
