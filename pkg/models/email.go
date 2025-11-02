package models

// EmailDto -> describe dto for email message
type EmailDto struct {
	ServiceType string `json:"serviceType"` // email OR telegram
	DomainName  string `json:"domainName"`  // name of API
	Content     string `json:"content"`     // content to send
	Recipient   string `json:"recipient"`   // user email or telegram chatId
}
