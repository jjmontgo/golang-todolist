package frame

import (
	"fmt"
	"net"
	"net/smtp"
	"net/url"
	"net/http"
	"crypto/tls"
	"log"
	"os"
)

// for sending email locally
func Email(to string, subject string, body string, from string) {
	if from == "" {
		from = os.Getenv("EMAIL_FROM")
	}

	if from == "" {
		log.Print("Controller.Email(): No From address provided")
		return
	}

	if to == "" {
		log.Print("Controller.Email(): No To address provided")
		return
	}

	if os.Getenv("MODE") == "prod" {
		sendViaEmailService(to, subject, body, from)
		return
	}

	sendSmtpEmail(to, subject, body, from)
}

func sendViaEmailService(to string, subject string, body string, from string) {
	// "https://k8fwzy59a0.execute-api.us-east-1.amazonaws.com/prod/"
	http.PostForm(os.Getenv("EMAIL_SERVICE_URL"), url.Values{
		"to": {to},
		"subject": {subject},
		"body": {body},
		"from": {from},
	})
}

// local dev emails are sent through SMTP
func sendSmtpEmail(to string, subject string, body string, from string) {
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k,v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := os.Getenv("EMAIL_HOST") + ":465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("",os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_HOST"))

	// TLS config
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}
