# GophMail

This is a simple `CLI` built in `Go` that enables you to send mails to multiple users in multiple templates, without ever leaving your command line.

## About the CLI

The CLI needs two external files to generate and send the emails:

- `Data File`: This is a `csv` file with headers which contains the various data about the recepients.
- `Template File`: This is a `html` file which instructs the CLI on how to create the body of the mail.

## Usage

While using the CLI, you can use the following flags:

- `-e`: Specify the sender's email address
- `-p`: Specify the password for the sender's email address
- `-t`: Specify the path to the template file
- `-d`: Specify the path to the data containing

You can use the `-h` to see the help about the various flags you can pass.

#### Constraints on the Data File

- It must be a `.csv` or `.tsv` file.
- The delimeter must be `,` and the new line delimeter must be `\n`.
- The file must have headers.
- It must contain atleast one header called `Recipients` which contains the email address of the recipients.
