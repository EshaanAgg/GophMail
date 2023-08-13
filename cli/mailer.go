package cli

import "fmt"

func (flags *InputFlags) sendMail(content string, recipient string) {
	fmt.Println(content)
	fmt.Println(recipient)
}
