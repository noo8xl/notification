package models

// SendTwoStepCodeDto -> via telegram
type SendTwoStepCodeDto struct {
	Message string `json:"message"`
	ChatID  string `json:"chatId"`
}

// TelegramAuthConfig -> auth bot config type
type TelegramBotConfig struct {
	Token string
	BotID int64
	Key   string
}

// ClientRegistrationDto -> signup a new API client obj
type ClientRegistrationDto struct {
	UserEmail  string `json:"userEmail"`  // client email
	DomainName string `json:"domainName"` // name of client domain
}

// NotificationRequestDto -> is notification request dto
type NotificationRequestDto struct {
	ServiceType string `json:"serviceType"` // email OR telegram
	DomainName  string `json:"domainName"`  // name of client API
	Content     string `json:"content"`     // content to send
	Recipient   string `json:"recipient"`   // user email or telegram chatId
}
