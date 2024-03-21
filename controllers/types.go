package controller

type SendUserNotif interface {
	SendEmailMessage() error
	SendTelegramMessage() error
	SaveHistory()
}

// NotificationRequestDto -> is notification request dto
type NotificationRequestDto struct {
	ServiceType string `json:"serviceType"` // email OR telegram
	DomainName  string `json:"domainName"`  // name of client API
	Content     string `json:"content"`     // content to send
	Recipient   string `json:"recipient"`   // user email or telegram chatId
}
