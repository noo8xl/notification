package http

import (
	database "notification-api/internal/repository"
	"notification-api/internal/service/email"
	"notification-api/internal/service/telegram"
)

// the doc is here ->
// -> https://docs.gofiber.io/api/ctx/

// ===============================

type AuthService struct {
	db *database.DatabaseService
}

type NotificationService struct {
	db *database.DatabaseService
	tg *telegram.TelegramService
	e  *email.EmailService
}

func InitAuthService() *AuthService {
	return &AuthService{
		db: database.InitDatabaseService(),
	}
}

func InitNotificationService() *NotificationService {
	db := database.InitDatabaseService()
	tg := telegram.NewTelegramService()
	e := email.NewEmailService()

	return &NotificationService{db, tg, e}
}

func InitController() {

}
