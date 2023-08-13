# GophMail

This is a simple `cli` built in `Go` that enables you to send mails to multiple users in multiple templates, without ever leaving your command line.

## About the CLI

The CLI needs two external files to generate and send the emails:

- `Data File`: This is a `csv` file with headers which contains the various data about the recepients.
- `Template File`: This is a `html` file which instructs the CLI on how to create the body of the mail.

## Usage

While using the CLI, you can use the following flags:

| Flag | Description                             | Default Value   | Required |
| ---- | --------------------------------------- | --------------- | -------- |
| `-e` | Sender's email address                  | NA              | Y        |
| `-p` | Password for the sender's email service | NA              | Y        |
| `-s` | Status of the mails to be sent          | ""              | N        |
| `-t` | Path to the template file               | `template.html` | N        |
| `-d` | Path to the data file                   | `data.csv`      | N        |

You can use the `-h` to see the help about the various flags you can pass.

#### Constraints on the Data File

- It must be a `.csv` or `.tsv` file.
- The delimeter must be `,` and the new line delimeter must be `\n`.
- The file must have headers.
- All the headers which are are to be used in the template must be single words.
- It must contain atleast one header called `Recipient` which contains the email address of the recipients.

#### Constraints on the Template File

- The template file would would need to specify the property `X` as `{{ .X }}` in the template to be populated.
- All the names are case-sensitive.

#### Note

You can checkout a sample [template file](./template.html) and [data file](./data.csv) here, or can edit them as the need arises.
