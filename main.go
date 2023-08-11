package main

import (
	"flag"
	"gophmail/cli"
)

func main() {
	var flags cli.InputFlags

	flag.StringVar(&flags.SenderEmail, "s", "", "The email address to send the mails from")
	flag.StringVar(&flags.Password, "p", "", "The password of the sender email account")
	flag.StringVar(&flags.Template, "t", "./template.html", "Relative path to the template file to generate your mail body")
	flag.StringVar(&flags.DataFile, "d", "./data.csv", "Relative path to the data file containing data about the receipients")
	flag.BoolVar(&flags.Help, "help", false, "Help about the default arguments of the CLI")

	flag.Parse()

	if flags.Help {
		flag.PrintDefaults()
	} else {
		flags.Send()
	}
}
