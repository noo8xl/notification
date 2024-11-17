package telegram

import (
	"fmt"
	"notification-api/config"
	"notification-api/excepriton"
	"notification-api/models"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type botSettings struct {
	errorChatID     int64
	notificationBot *tgbotapi.BotAPI
	errorBot        *tgbotapi.BotAPI
}

// var errorChatID int64
// var notificationBot *tgbotapi.BotAPI
// var errorBot *tgbotapi.BotAPI

// // InitAuthBot -> init auth bot for use telegram.passport
// func InitAuthBot() *tgbotapi.BotAPI {
// 	var err error
// 	var conf = config.GetTelegramConfig()
// 	//fmt.Println("conf is ->\n", conf)

// 	bot, err := tgbotapi.NewBotAPI(conf.Token)
// 	if err != nil {
// 		excepriton.HandleAnError("auth bot init got an error: " + err.Error())
// 		os.Exit(1)
// 	}

// 	// fmt.Println(bot.Self.UserName)
// 	// bot.Debug = true
// 	return bot
// }

// init -> init notif bot here and set some variables
func initBot() *botSettings {

	var bt = new(botSettings)
	devChatId := config.GetDevChatId()
	temp, _ := strconv.Atoi(devChatId)

	bt.errorChatID = int64(temp)
	bt.errorBot = initErrorBot()
	bt.notificationBot = initNotificationBot()

	return bt
}

// SendUserMessage -> send a two-step code message to the current chatId
func SendUserMessage(dto *models.SendTwoStepCodeDto) error {
	var err error
	var temp, _ = strconv.Atoi(dto.ChatID)
	var chatID = int64(temp)

	//ctx := strings.Join([]string{
	//  "Your new auth code is ",
	//  "< ", dto.Code, " >",
	//  ". This code is available only a few minutes."}, "")

	msg := tgbotapi.NewMessage(chatID, dto.Message)

	bot := initBot()
	_, err = bot.notificationBot.Send(msg)
	if err != nil {
		fmt.Println("err is - >\n", err)
	}

	return err
}

// SendErrorMessage -> send a message to the current chatID (to a developer)
func SendErrorMessage(ctx string) error {
	bot := initBot()
	msg := tgbotapi.NewMessage(bot.errorChatID, ctx)
	_, err := bot.errorBot.Send(msg)
	if err != nil {
		excepriton.HandleAnError("Send message was failed." + err.Error())
		return err
	}
	return nil
}
