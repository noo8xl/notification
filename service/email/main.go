package email

import (
	"notification-api/excepriton"
	"notification-api/service/telegram"

	"gopkg.in/gomail.v2"
)

type Opts struct {
	host     string
	port     int
	username string
	password string
}
type EmailService struct {
	opts *Opts
	tg   *telegram.TelegramService
}

func NewEmailService() *EmailService {

	// sandbox implementation <-
	opts := &Opts{
		host:     "sandbox.smtp.mailtrap.io",
		port:     25,
		username: "32807a4cddbd87",
		password: "d349cff32824c7",
		// password: config.GetGmailSecret() // <- my gmail secrete pwd here
	}
	tg := telegram.NewTelegramService()
	return &EmailService{tg: tg, opts: opts}
}

// sendMessageViaEmail -> Send the email
func (s *EmailService) sendMessageViaEmail(msg *gomail.Message) error {

	dialer := gomail.NewDialer(s.opts.host, s.opts.port, s.opts.username, s.opts.password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		excepriton.HandleAnError("email handler got an error: " + err.Error())
		s.tg.SendErrorMessage("email handler got an error.")
	}
	return err
}

// ##################################################################################
// ##################################################################################
// ##################################################################################

// production-ready smtp service doc is:
// ---> https://docs.unione.io/en/smtp-api <---
//

// url := "https://us1.unione.io/ru/transactional/api/v1/email/send.json"

// 	"message": {
// 		"recipients": [
// 			{
// 				"email": `${userEmail}`
// 			}
// 		],
// 		"skip_unsubscribe": 0,
// 		"global_language": "en",
// 		"body": {
// 			"html": `${htmlData}`,
// 			"plaintext": `Hello, ${userName}`,
// 			"amp": "<!doctype html><html amp4email><head> <meta charset=\"utf-8\"><script async src=\"https://cdn.ampproject.org/v0.js\"></script> <style amp4email-boilerplate>body{visibility:hidden}</style></head><body> Hello, AMP4EMAIL world.</body></html>"
// 		},
// 		"subject": "email verification",
// 		// "from_email": fromEmail,
// 		"from_email": DOMAIN_TEST,
// 		"from_name": "no-reply",
// 		"track_links": 0,
// 		"track_read": 0,
// 	}
// };
