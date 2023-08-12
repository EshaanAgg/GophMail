package cli

import (
	"encoding/csv"
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

func (flags *InputFlags) parseData() (map[string]int, [][]string) {

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
	headerIndex := make(map[string]int)
	foundReceipients := false

	for ind, header := range headers {
		headerIndex[header] = ind
		if header == "Recipients" {
			foundReceipients = true
		}
	}

	if !foundReceipients {
		sendError("The data file has no header named 'Recipients'. Please make sure that the same is present.")
	}

	return headerIndex, records[1:]
}

func (flags *InputFlags) parseTemplate() {}
