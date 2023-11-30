package cli

import (
	"fmt"
)

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

	// Keeps track of the number of threads that we want to create
	registerChan := make(chan int)
	go func() {
		for i := 0; i < 25; i++ {
			registerChan <- i
		}
	}()

	// Asks like a task queue, keeping track of the tasks that are yet to be scheduled
	tasks := make(chan int)
	go func() {
		for i := 0; i < len(mails); i++ {
			tasks <- i
		}
	}()

	// Keeps track of the number of tasks that have been completed successfully
	successTasks := 0
	loop := true

	for loop {
		select {
		case task := <-tasks:
			go func() {
				worker := <-registerChan
				sent := flags.sendMail(mails[task], data[task]["Recipient"], true)
				if sent {
					successTasks++
					registerChan <- worker
				} else {
					tasks <- task
				}
			}()

		default:
			if successTasks == len(mails) {
				close(tasks)
				loop = false
			}
		}
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
