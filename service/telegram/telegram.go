package telegram

import (
	"fmt"
	"notification-api/config"
	"notification-api/excepriton"
	"notification-api/models"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var errorChatID int64
var notificationBot *tgbotapi.BotAPI
var errorBot *tgbotapi.BotAPI

// init -> init notif bot here and set some variables
func init() {

	devChatId := config.GetDevChatId()
	temp, _ := strconv.Atoi(devChatId)
	errorChatID = int64(temp)

	errorBot = initErrorBot()
	notificationBot = initNotificationBot()
}

// SendUserMessage -> send a two-step code message to the current chatId
func SendUserMessage(dto *models.SendTwoStepCodeDto) error {
	fmt.Println("dto is -> ", dto.ChatID, dto.Message)

	var err error
	var temp, _ = strconv.Atoi(dto.ChatID)
	var chatID = int64(temp)

	//ctx := strings.Join([]string{
	//  "Your new auth code is ",
	//  "< ", dto.Code, " >",
	//  ". This code is available only a few minutes."}, "")

	msg := tgbotapi.NewMessage(chatID, dto.Message)

	_, err = notificationBot.Send(msg)
	if err != nil {
		fmt.Println("err is - >\n", err)
	}

	return err
}

// SendErrorMessage -> send a message to the current chatID (to a developer)
func SendErrorMessage(ctx string) error {
	msg := tgbotapi.NewMessage(errorChatID, ctx)
	_, err := errorBot.Send(msg)
	if err != nil {
		excepriton.HandleAnError("Send message was failed.", err)
		return err
	}
	return nil
}
