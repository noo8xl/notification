package config

// import (
// 	"fmt"
// 	"notification-api/models"
// 	"os"
// 	"strconv"
// 	"strings"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"github.com/joho/godotenv"
// )

// // GetTelegramConfig -> get auth bot config
// func GetTelegramConfig() models.TelegramAuthConfig {

// 	path := getConfigVar("AUTH_RSA_PATH")
// 	key, err := os.ReadFile(path)
// 	if err != nil {
// 		fmt.Println("GetTelegramConfig err -> ", err)
// 	}

// 	id, _ := strconv.Atoi(getConfigVar("AUTH_BOT_ID"))

// 	return models.TelegramAuthConfig{
// 		Token: getConfigVar("AUTH_BOT_TOKEN"),
// 		BotID: int64(id),
// 		Key:   string(key),
// 	}
// }

// func GetErrorHandlerToken() string {
// 	return getConfigVar("ERROR_BOT_TOKEN")
// }

// // BotConfig -> get base config for the apdates
// func BotConfig() tgbotapi.UpdateConfig {
// 	updateConfig := tgbotapi.NewUpdate(0)
// 	updateConfig.Timeout = 120
// 	return updateConfig
// }

// // func GetClientAPIAccessToken() string {
// // 	return getConfigVar("CLIENT_API_ACCESS_TOKEN")
// // }

// // GetAccessToken -> get an access token, which used as header to get access to the API
// func GetNotificationAPIAccessToken() string {
// 	return getConfigVar("API_ACCESS_TOKEN")
// }

// func GetNitificationBotToken() string {
// 	return getConfigVar("NOTIFICATION_BOT_TOKEN")
// }

// // GetDevChatId -> get personal chat id for handle errors *
// func GetDevChatId() string {
// 	return getConfigVar("ERROR_CHAT_ID")
// }

// // GetMONGOdatabaseConfig -> get mongodb config from env file
// func GetMONGOdatabaseConfig() [2]string {
// 	dbUser := getConfigVar("MONGO_DB_USER")
// 	dbPassword := getConfigVar("MONGO_DB_PASSWORD")
// 	dbName := getConfigVar("MONGO_DB_NAME")

// 	mongoConnectString := strings.Join([]string{"mongodb+srv://", dbUser, ":", dbPassword, "@cluster001.sipjs.mongodb.net/?retryWrites=true&w=majority"}, "")

// 	fmt.Println("link -> ", mongoConnectString)
// 	return [2]string{mongoConnectString, dbName}
// }

// // ###################################################################

// func getConfigVar(key string) string {
// 	if err := godotenv.Load(".env"); err != nil {
// 		fmt.Println(err)
// 	}
// 	return os.Getenv(key)
// }

// // func ClientApiAuth() string {
// // 	var apiUrl string = clientAPIUrl()
// // 	var signInPath string = "/auth/passport-telegram/"
// // 	var requestStr string = strings.Join([]string{apiUrl, signInPath}, "")

// // 	fmt.Println("request str -> ", requestStr)
// // 	return requestStr
// // }

// // ####################################################################################

// // func clientAPIUrl() string {
// // 	return "http://127.0.0.1:35891"
// // 	// return "https://127.0.0.1:35891"
// // }
