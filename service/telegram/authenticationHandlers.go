package telegram

import (
	"log"
	"notification-api/excepriton"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// authHandler -> handle auth messages from <sign in via telegram>
func authHandler(update *tgbotapi.Update) bool {

	log.Println("update ->\n", update)

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

// defaultMessageHandler -> send a default message
// to the auth bot if got an unexpected message
func defaultMessageHandler(chatId int64, bt *tgbotapi.BotAPI) {
	txt := "AuthBot is only for the auth and don't handle any user messages. If you have some questions please contact support via profile->support!"
	msg := tgbotapi.NewMessage(chatId, txt)
	_, err := bt.Send(msg)
	if err != nil {
		excepriton.HandleAnError("auth bot init got an error: " + err.Error())
		return
	}
}
