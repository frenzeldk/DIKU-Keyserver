package mail

import (
    "gopkg.in/gomail.v1"
)

func Send(rcpt, body string) {
    msg := gomail.NewMessage()
    msg.SetHeader("From", "noreply@dikukeys.dk")
    msg.SetHeader("To", rcpt)
    msg.SetHeader("Subject", "Hello!")
    msg.SetBody("text/html", "Hello <b>Bob</b>!")

    // Send the email to Bob
    mailer := gomail.NewMailer("localhost", "dikukeys", "", 25)
    if err := mailer.Send(msg); err != nil {
        panic(err)
    }
}

/* 
import (
        "bytes"
        "log"
        "net/smtp"
)

func Send(rcpt, body string) {
        // Connect to the remote SMTP server.
        c, err := smtp.Dial("dikukeys.dk:25")
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
        buf := bytes.NewBufferString(body)
        if _, err = buf.WriteTo(wc); err != nil {
                log.Fatal(err)
        }
} */