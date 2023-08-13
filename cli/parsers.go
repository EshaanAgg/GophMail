package cli

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
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
