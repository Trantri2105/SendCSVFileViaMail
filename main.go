package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"os"

	"github.com/joho/godotenv"
)

func loadEnvVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func sendCSVFileViaMail(subject string, body string, filePath string, to []string) error {
	password := os.Getenv("PASSWORD")
	from := os.Getenv("FROM")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var fileInfo fs.FileInfo
	fileInfo, err = file.Stat()
	if err != nil {
		return err
	}
	fileContent := make([]byte, fileInfo.Size())
	file.Read(fileContent)
	defer file.Close()

	var email bytes.Buffer
	writer := multipart.NewWriter(&email)

	//Email header
	email.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	email.WriteString("MIME-Version: 1.0\r\n")
	email.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n", writer.Boundary()))
	email.WriteString("\r\n")

	//Email body
	bodyHeader := textproto.MIMEHeader{}
	bodyHeader.Set("Content-Type", "text/plain")
	bodyPath, err := writer.CreatePart(bodyHeader)
	if err != nil {
		return err
	}
	bodyPath.Write([]byte(body))

	//Email attachment
	attachmentHeader := textproto.MIMEHeader{}
	attachmentHeader.Set("Content-Type", "text/csv")
	attachmentHeader.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name()))
	attachment, err := writer.CreatePart(attachmentHeader)
	if err != nil {
		return err
	}
	attachment.Write(fileContent)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, from, to, email.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func main() {
	loadEnvVariable()
	filePath := "example.csv"
	to := []string{"example@gmail.com"}
	subject := "CSV file"
	body := "Hello world"
	err := sendCSVFileViaMail(subject, body, filePath, to)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sent mail successfully")
}
