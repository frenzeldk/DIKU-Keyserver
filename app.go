package main

import (
	"encoding/hex"
	//"fmt"
	"github.com/Orkeren/DIKU-Keyserver/golibs/hash"
	"github.com/Orkeren/DIKU-Keyserver/golibs/mail"
	"html/template"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"strconv"
	//"strings"
	"time"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	//kuid is the KU ID of the student
	kuid := req.FormValue("kuid")
	//ctime is the time of creation of the link (as unix time)
	ctime := req.FormValue("ctime")
	//hash is the padded sha3-512 hash of kuid & ctime)
	coffee_hash := req.FormValue("hash")

	pubkey := req.FormValue("pubkey")

	//rcpt is the e-mail address associated with kuid
	//rcpt := strings.Join([]string{kuid, "alumni.ku.dk"}, "@")
	rcpt := kuid

	if kuid == "" {
		resp.Write([]byte("<form>KU-ID:<br><input type='text' name='kuid'>@alumni.ku.dk<br><input type='submit' value='Send'></form>"))
	} else if coffee_hash == "" {
		ctime = strconv.FormatInt(time.Now().Unix(), 10)
		coffee_hash = hex.EncodeToString(hash.GetHash(kuid, ctime)[:])

		//body is the plaintext body of the email.
		body := `Dette er epostens krop.
                 linjeskifte laves ved at have regulære linjeskift.
	             linket er: http://dikukeys.dk:8081/app?kuid=` + kuid + "&ctime=" + ctime + "&hash=" + coffee_hash
		//only send an email if rcpt has a value. This needs to be changed to regex for a valid e-email adress (user@domain.tld)
		if rcpt != "@alumni.ku.dk" {
			mail.Send(rcpt, body)
		}
		resp.Write([]byte("<p>Registration e-mail sent!</p>"))
	} else if hex.EncodeToString(hash.GetHash(kuid, ctime)[:]) == coffee_hash {
		resp.Write([]byte("<form>public key:<br><input type='text' name='pubkey'><br><input type='submit' value='Send'></form>"))
	} else {
		resp.Write([]byte("<p>Not a valid link!</p>"))
	}

	//fmt.Println("A mail has been sent to:", rcpt)
	//fmt.Println(time.Now().Unix())
	//fmt.Println("Deres Hash var", hash.GetHash(rcpt))
	type User struct {
		KUID string
		PUBKEY string
	}
	if kuid != "" { if pubkey != "" {
	cuser := User{kuid, pubkey}
	tmpl, err := template.New("test").Parse("{{.KUID}} has submitted the public key {{.PUBKEY}}")
	if err != nil {panic(err)}
	err = tmpl.Execute(os.Stdout, cuser)
	if err != nil { panic(err)}	
	}}
	
	//fmt.Println(out)
	//fmt.Println(pubkey)
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}
