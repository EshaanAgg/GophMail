package cli

import "fmt"

type InputFlags struct {
	SenderEmail string
	Subject     string
	Password    string

	Template string
	DataFile string

	AppPassword string
	SMTPSever   string
	SMTPPort    int

	Help bool
}

func (flags *InputFlags) Send() {
	flags.parseEnv()
	flags.validate()

	data := flags.parseData()
	mails := flags.generateMailContent(data)

	if !flags.confirmMail(mails[0], data[0]["Recipient"]) {
		sendError("Operation cancelled.")
	}

	fmt.Print("Sending mails...\n\n")

	for i, mail := range mails {
		flags.sendMail(mail, data[i]["Recipient"], true)
	}
}

func (flags *InputFlags) validate() {

	if len(flags.SenderEmail) == 0 {
		sendError("Sender email is a REQUIRED argument which can't be blank. Please specify the same using the '-e' flag.")
	}

	if len(flags.Password) == 0 {
		sendError("Password is a REQUIRED argument which can't be blank. Please specify the same using the '-p' flag.")
	}
}
