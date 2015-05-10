package mail

import (
	"bytes"
	"log"
	"net/smtp"
)

func Send(rcpt, body string) {
	stmp_server := "dikukeys.dk:25"
	from := `From: DIKU Keys <noreply@dikukeys.dk>`
	to := "To: " + rcpt + ``
	subject := `Subject: Velkommen til DIKU Keys`
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"

	epost := from + to + subject + mime + body
	// Connect to the remote SMTP server.
	c, err := smtp.Dial(stmp_server)
	if err != nil {
		log.Fatal(err)
	}
	// Set the sender and recipient.
	c.Hello("dikukeys.dk")
	c.Mail("noreply@dikukeys.dk")
	c.Rcpt(rcpt)
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString(epost)
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
