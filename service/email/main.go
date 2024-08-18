package email

import (
	"notification-api/models"
	lettershtmlcontent "notification-api/service/email/letters.html.content"
	"strings"

	"gopkg.in/gomail.v2"
)

// PrepareEmailMessage -> set email opts and then call sendMessageViaEmail func
func PrepareEmailMessage(dto models.EmailDto) error {

	from := strings.Join([]string{"no-repl@", dto.DomainName}, "")
	to := dto.Recipient
	subject := "Authentication"                                      // -> up
	htmlCtx := lettershtmlcontent.GeTwoFactorCodeLetter(dto.Content) // get an HTML doc here <-

	// Choose an auth method and set it up

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", htmlCtx)
	// msg.Attach("/home/cats/cat.jpg")

	return sendMessageViaEmail(msg)
}
