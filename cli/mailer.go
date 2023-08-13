package cli

import (
	"fmt"

	"github.com/go-gomail/gomail"
)

func (flags *InputFlags) confirmMail(content string, recipient string) bool {
	fmt.Print("Here is a sample mail that was generated:\n\n")
	fmt.Printf("To: %s\n", recipient)
	fmt.Printf("Mail Body:\n%s", content)
	fmt.Print("\n\nShould I send the mails? (Y/N) ")

	var response string
	n, err := fmt.Scanln(&response)
	if n != 1 || err != nil {
		sendError("Invalid response. Aborting execution.")
	}

	return response == "Y"
}

func (flags *InputFlags) sendMail(content string, recipient string) {
	m := gomail.NewMessage()

	m.SetHeader("From", flags.SenderEmail)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", flags.Subject)
	m.SetBody("text/html", content)

	// SMTP configuration
	d := gomail.NewDialer(flags.SMTPSever, flags.SMTPPort, flags.SenderEmail, flags.AppPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		sendError("Mail to " + recipient + " failed.")
	}

	fmt.Printf("Email to %s sent successfully.\n", recipient)
}
