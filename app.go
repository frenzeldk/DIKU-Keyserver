package main

import (
	"fmt"
	"github.com/Orkeren/DIKU-Keyserver/golibs/hash"
	"github.com/Orkeren/DIKU-Keyserver/golibs/mail"
	"net"
	"net/http"
	"net/http/fcgi"
	"time"
	//  "strings"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<form>KU-ID:<br><input type='text' name='ku_id'><br><input type='submit' value='Send'></form>"))
	//kuid is the KU ID of the student
	kuid := req.FormValue("ku_id")
	//rcpt is the e-mail address associated with kuid
	rcpt := strings.Join([]string{kuid, "alumni.ku.dk"}, "@")
	//body is the plaintext body of the email.
	body := `Dette er epostens krop.
	         linjeskifte laves ved at have regulære linjeskift.`
	//only send an email if rcpt has a value. This needs to be changed to regex for a valid e-email adress (user@domain.tld)
	if rcpt != "" {
		mail.Send(rcpt, body)
	}
	fmt.Println("A mail has been sent to:", rcpt)
	fmt.Println(time.Now())
	fmt.Println("Deres Hash var", hash.GetHash(rcpt))
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}
