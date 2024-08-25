package telegram

import (
  "fmt"
  "log"
  "notification-api/config"
  "notification-api/models"
  "os"
  "strconv"
  "time"

  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var errorChatID int64
var notificationBot *tgbotapi.BotAPI
var errorBot *tgbotapi.BotAPI

// init -> init bot here and set some variables
func init() {

  var devChatId string = config.GetDevChatId()
  var temp, _ = strconv.Atoi(devChatId)
  errorChatID = int64(temp)
  // fmt.Println("errorChatID ", errorChatID)

  errorBot = initErrorBot()
  notificationBot = initNotificationBot()
}

// SendUserMessage -> send a two-step code message to the current chatId
func SendUserMessage(dto models.SendTwoStepCodeDto) error {
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
func SendErrorMessage(ctx string) {
  msg := tgbotapi.NewMessage(errorChatID, ctx)
  _, err := errorBot.Send(msg)
  if err != nil {
    fmt.Println("Send message was failed.")
    return
  }
}

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
