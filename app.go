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
	rcpt := req.FormValue("ku_id")
	//epost := strings.Join([]string{rcpt, "alumni.ku.dk"}, "@")
	if rcpt != "" {
		mail.Send(rcpt, `Kære liste,
			denne epost er sendt igennem den golang-app,
			som vi har brugt dagen i dag på at kode.
			
			Venligst
			Thorkil, Mads & Sven`)
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
