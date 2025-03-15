package telegram

import (
	"fmt"
	"notification-api/config"
	"notification-api/excepriton"
	"notification-api/models"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService struct {
	notificationBot *tgbotapi.BotAPI
	errorBot        *tgbotapi.BotAPI
}

func NewTelegramService() *TelegramService {

	eb := initErrorBot()
	nb := initNotificationBot()

	return &TelegramService{
		notificationBot: nb,
		errorBot:        eb,
	}
}

// SendUserMessage -> send a two-step code message to the current chatId
func (s *TelegramService) SendUserMessage(dto *models.SendTwoStepCodeDto) error {
	var temp, _ = strconv.Atoi(dto.ChatID)
	var chatID = int64(temp)

	msg := tgbotapi.NewMessage(chatID, dto.Message)

	_, err := s.notificationBot.Send(msg)
	if err != nil {
		fmt.Println("err is - >\n", err)
	}

	return err
}

// SendErrorMessage -> send a message to the current chatID (to a developer)
func (s *TelegramService) SendErrorMessage(ctx string) error {

	devChatId := config.GetDevChatId()
	temp, _ := strconv.Atoi(devChatId)
	errorChatID := int64(temp)

	msg := tgbotapi.NewMessage(errorChatID, ctx)
	_, err := s.errorBot.Send(msg)
	if err != nil {
		excepriton.HandleAnError("Send message was failed." + err.Error())
		return err
	}
	return nil
}
