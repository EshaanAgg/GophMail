package cli

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func (flags *InputFlags) parseData() []map[string]string {
	f, err := os.Open(flags.DataFile)
	if err != nil {
		sendError(fmt.Sprintf("Unable to read data file at %s. Please make sure that the file exists at this location.", flags.DataFile))
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		sendError(fmt.Sprintf("Unable to read data file at %s. Please make sure that the file is a valid CSV.", flags.DataFile))
	}

	if len(records) == 0 {
		sendError("The parsed CSV file is empty.")
	}

	headers := records[0]
	data := []map[string]string{}
	foundReceipient := false

	for _, header := range headers {
		if header == "Recipient" {
			foundReceipient = true
		}
	}

	if !foundReceipient {
		sendError("The data file has no header named 'Recipient'. Please make sure that the same is present.")
	}

	for i := 1; i < len(records); i++ {
		record := records[i]
		dataEntry := map[string]string{}

		for ind, header := range headers {
			dataEntry[header] = record[ind]
		}

		data = append(data, dataEntry)
	}

	return data
}

func (flags *InputFlags) generateMailContent(data []map[string]string) []string {
	t, err := template.ParseFiles(flags.Template)

	if err != nil {
		sendError("The given template file does not exist or is not of valid type. Please check the requirements.")
	}

	mails := []string{}

	for _, d := range data {
		var mail bytes.Buffer
		t.Execute(&mail, d)
		mails = append(mails, mail.String())
	}

	return mails
}

func (flags *InputFlags) parseEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	fmt.Println(".env file found. Populating variables from the same.")

	if len(os.Getenv("MAIL_ID")) != 0 {
		flags.SenderEmail = os.Getenv("MAIL_ID")
	}

	if len(os.Getenv("MAIL_PASSWORD")) != 0 {
		flags.Password = os.Getenv("MAIL_PASSWORD")
	}

	if len(os.Getenv("SMTP_SERVER")) != 0 {
		flags.SMTPSever = os.Getenv("SMTP_SERVER")
	}

	if len(os.Getenv("SMTP_SERVER_PORT")) != 0 {
		p, err := strconv.Atoi(os.Getenv("SMTP_SERVER_PORT"))
		if err != nil {
			fmt.Println("Invalid SMTP port. Using default port instead.")
		}
		flags.SMTPPort = p
	}

	if len(os.Getenv("SMTP_APP_PASSWORD")) != 0 {
		flags.AppPassword = os.Getenv("SMTP_APP_PASSWORD")
	}
}
