package email

import (
  "fmt"
  "notification-api/service/telegram"

  "gopkg.in/gomail.v2"
)

// sendMessageViaEmail -> Send the email
func sendMessageViaEmail(msg *gomail.Message) error {

  // sandbox implementation <-
  user := "32807a4cddbd87"
  pwd := "d349cff32824c7"
  // pwd := config.GetGmailSecret() // <- my gmail secrete pwd here
  smtpHost := "sandbox.smtp.mailtrap.io"
  smtpPort := 25

  n := gomail.NewDialer(smtpHost, smtpPort, user, pwd)
  err := n.DialAndSend(msg)
  if err != nil {
    fmt.Println("email handler error.")
    telegram.SendErrorMessage("email handler got an error.")
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
