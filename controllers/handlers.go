package controller

import (
	"notification-api/models"
	"notification-api/service/database"
	"notification-api/service/email"
	"notification-api/service/telegram"
	"time"
)

// NotificationRequestDto -> is notification request dto
type NotificationRequestDto struct {
	ServiceType string `json:"serviceType"` // email OR telegram
	DomainName  string `json:"domainName"`  // name of the client API
	Content     string `json:"content"`     // content to send
	Recipient   string `json:"recipient"`   // user email or telegram chatId
}

// ##################################################################################
// ########################## => dto handlers .- ####################################
// ##################################################################################

// SendTelegramMessage -> handle only a telegram notification sending
func (dto *NotificationRequestDto) SendTelegramMessage() error {
	d := models.SendTwoStepCodeDto{
		ChatID:  dto.Recipient,
		Message: dto.Content,
	}
	return telegram.SendUserMessage(&d)
}

// SendEmailMessage - > handle only an email notification sending
func (dto *NotificationRequestDto) SendEmailMessage() error {
	d := models.EmailDto{
		ServiceType: dto.ServiceType,
		DomainName:  dto.DomainName,
		Content:     dto.Content,
		Recipient:   dto.Recipient,
	}
	return email.PrepareEmailMessage(&d)
}

func (dto *NotificationRequestDto) SaveToTheHistory() error {
	historyItem := models.NotificationHistory{
		DateTime:    time.Now().Format(time.UnixDate),
		Recipient:   dto.Recipient,
		DomainName:  dto.DomainName,
		MessageBody: dto.Content,
		SentVia:     dto.ServiceType,
	}
	// fmt.Println("hist ->\n", historyItem)
	return database.SaveToTheHistory(&historyItem)
}
