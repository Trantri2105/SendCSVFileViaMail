# Send CSV file via mail
## Overview
- Go project show how to send CSV file via mail
## Setup
- To use gmail for sending CSV file, you have to enable Two-Factor Authentication and [create an application password](https://support.google.com/mail/answer/185833?hl=vi)
- Set environment variable in .env file
  - `FROM`: Mail for sending
  - `PASSWORD`: Application password
  - `SMTP_HOST`: Smtp host. Example: `smtp.gmail.com` for ...@gmail.com
  - `SMTP_PORT`: Smtp port
## How the code run
1. Load environment variables in .env file
2. Create email and send by calling `sendCSVFileViaMail(subject string, body string, filePath string, to []string)` method
   - This method do the following
     - Write subject to email header
     - Write body message to email body
     - Attach csv file and send to smtp server
   - Parameter
     - `subject`: subject of an email
     - `body`: body of an email
     - `filePath`: csv file path
     - `to`: slice of received email