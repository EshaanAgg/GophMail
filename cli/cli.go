package cli

import (
	"fmt"
	"os"
)

type InputFlags struct {
	SenderEmail string
	Password    string
	Template    string
	DataFile    string
	Help        bool
}

func (flags *InputFlags) Send() {
	flags.validate()
}

func (flags *InputFlags) validate() {
	valid := true

	if len(flags.SenderEmail) == 0 {
		fmt.Println("Sender email is a REQUIRED argument which can't be blank. Please specify the same using the '-s' flag.")
		valid = false
	}

	if len(flags.Password) == 0 {
		fmt.Println("Password is a REQUIRED argument which can't be blank. Please specify the same using the '-p' flag.")
		valid = false
	}

	if !valid {
		os.Exit(0)
	}
}
