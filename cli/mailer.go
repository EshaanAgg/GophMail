package cli

import (
	"fmt"

	"github.com/go-gomail/gomail"
)

func (flags *InputFlags) confirmMail(content string, recipient string) bool {
	sent := flags.sendMail(content, flags.SenderEmail, false)

	if !sent {
		sendError("The mail could not be sent due to TCP timeout. Please try again.")
	}

	fmt.Printf("A sample mail for %s was generated and mailed to the sender's email address (%s).\n", recipient, flags.SenderEmail)
	fmt.Print("You can check the same for content formating.\n\nShould I send all the mails? (Y/N) ")

	var response string
	n, err := fmt.Scanln(&response)
	if n != 1 || err != nil {
		sendError("Invalid response. Aborting execution.")
	}

	fmt.Print("\n\n")

	return response == "Y" || response == "y"
}

func (flags *InputFlags) sendMail(content string, recipient string, displayMessage bool) bool {
	m := gomail.NewMessage()

	m.SetHeader("From", flags.SenderEmail)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", flags.Subject)
	m.SetBody("text/html", content)

	// SMTP configuration
	d := gomail.NewDialer(flags.SMTPSever, flags.SMTPPort, flags.SenderEmail, flags.AppPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Mail to " + recipient + " failed.")
		fmt.Println(err)
		return false
	} else if displayMessage {
		fmt.Printf("Email to %s sent successfully.\n", recipient)
	}

	return true
}
