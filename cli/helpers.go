package cli

import (
	"fmt"
	"os"
)

func sendError(message string) {
	fmt.Println(message)
	os.Exit(0)
}
