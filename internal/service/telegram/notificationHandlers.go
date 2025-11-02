package telegram

import (
	"notification-api/config"
	"notification-api/pkg/exceptions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// // startMessageHandler -> handle only </start> user message
// func startMessageHandler(dto *models.CommandsDto) {
// 	txt := strings.Join([]string{
// 		"Hi, ", dto.UserName, "! ",
// 		"Welcome to AuthBot!", "should add some text <-"}, "")
// 	msg := tgbotapi.NewMessage(dto.ChatId, txt)
// 	_, err := dto.Bot.Send(msg)
// 	if err != nil {
// 		exceptions.HandleAnError("<startMessageHandler> got an error: " + err.Error())
// 		return
// 	}
// }

// // helpMessageHandler -> handle only </help> user message
// func helpMessageHandler(dto *models.CommandsDto) {

// 	greetings := "Let's start to teach you how to interact with me!\n"
// 	info1 := "_"
// 	info2 := "_"
// 	info3 := "_ will update soon"

// 	txt := strings.Join([]string{greetings, info1, info2, info3}, "\n")
// 	msg := tgbotapi.NewMessage(dto.ChatId, txt)
// 	_, err := dto.Bot.Send(msg)
// 	if err != nil {
// 		exceptions.HandleAnError("<helpMessageHandler> got an error: " + err.Error())
// 		return
// 	}
// }

// ##########################################################################
// ###################### -> connection area <- #############################
// ##########################################################################

// initErrorBot -> init tg bot for send ERRORRs
func initErrorBot() *tgbotapi.BotAPI {
	token := config.GetErrorHandlerToken()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		exceptions.HandleAnError("error handler bot init got an error: " + err.Error())
	}
	return bot
}

// initNotificationBot -> init tg bot for notification
func initNotificationBot() *tgbotapi.BotAPI {
	token := config.GetNitificationBotToken()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		exceptions.HandleAnError("notification bot init got an error: " + err.Error())
	}
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
