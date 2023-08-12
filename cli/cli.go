package cli

import (
	"fmt"
)

type InputFlags struct {
	SenderEmail string
	Password    string
	Template    string
	DataFile    string
	Status      string
	Help        bool
}

func (flags *InputFlags) Send() {
	flags.validate()
}

func (flags *InputFlags) validate() {

	if len(flags.SenderEmail) == 0 {
		sendError("Sender email is a REQUIRED argument which can't be blank. Please specify the same using the '-e' flag.")
	}

	if len(flags.Password) == 0 {
		sendError("Password is a REQUIRED argument which can't be blank. Please specify the same using the '-p' flag.")
	}

	headers, records := flags.parseData()
	flags.parseTemplate()

	fmt.Println(headers, records)
}
