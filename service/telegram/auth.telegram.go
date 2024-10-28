package telegram

import (
	"log"
	"notification-api/config"
	"notification-api/models"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// InitAuthBot -> init auth bot for use telegram.passport
func InitAuthBot() *tgbotapi.BotAPI {
	var err error
	var conf = config.GetTelegramConfig()
	//fmt.Println("conf is ->\n", conf)

	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	// fmt.Println(bot.Self.UserName)
	// bot.Debug = true
	return bot
}

// HandleUpdates -> waiting for update with telegram passport data (telegram-passport auth)
// and send data to the set API URL with an <accessKey> header
func HandleUpdates(update tgbotapi.Update, bt *tgbotapi.BotAPI) {
	// fmt.Println("current update item -> \n", update)

	dto := models.CommandsDto{
		UserName: update.Message.From.FirstName,
		Bot:      bt,
		ChatId:   update.Message.Chat.ID,
	}

	if update.Message != nil {
		switch update.Message.Text {
		case "/start":
			startMessageHandler(&dto)
			time.Sleep(time.Millisecond * time.Duration(1500))
			return
		case "/help":
			helpMessageHandler(&dto)
			time.Sleep(time.Millisecond * time.Duration(5000))
			return
		default:
			if update.Message.PassportData != nil {
				isSigned := authHandler(&update)
				if !isSigned {
					SendErrorMessage("Telegram auth API return en error. Auth status -> false")
					return
				} else {
					// send 2fa bot link and w8 for the 2fa code in auth bot
					return
				}
			} else {
				defaultMessageHandler(update.Message.Chat.ID, bt)
				time.Sleep(time.Millisecond * time.Duration(1500))
				return
			}
		}
	}
}
