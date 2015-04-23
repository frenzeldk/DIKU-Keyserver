package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Orkeren/DIKU-Keyserver/golibs/hash" // This is our hash function
	"github.com/Orkeren/DIKU-Keyserver/golibs/mail" // This is our mail function, it does hello
	"html/template"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"strconv"
	//"strings"
	"time"
)

type FastCGIServer struct{}
type Page struct {
	Title string
	Body  []byte
}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	titel := req.URL.Path[len("/"):]
	p, _ := loadPage(titel)
	fmt.Println(req.URL) // Dette viser bare hvordan man får en URL fra req
	//kuid is the KU ID of the student
	kuid := req.FormValue("kuid")
	//ctime is the time of creation of the link (as unix time)
	ctime := req.FormValue("ctime")
	//hash is the padded sha3-512 hash of kuid & ctime)
	coffee_hash := req.FormValue("hash")

	pubkey := req.FormValue("pubkey")

	rcpt := kuid + "@alumni.ku.dk"

	fmt.Println(rcpt)

	if kuid == "" {
		t, _ := template.ParseFiles("/home/keys/Orkeren/DIKU-Keyserver/html_templates/create_link.html")
		t.Execute(resp, p)
	} else if coffee_hash == "" {
		ctime = strconv.FormatInt(time.Now().Unix(), 10)
		coffee_hash = hex.EncodeToString(hash.GetHash(kuid, ctime)[:])

		//body is the plaintext body of the email.
		body := 
`Velkommen til dikukeys. For at afslutte registreringen, tryk venligst på dette link:
http://dikukeys.dk:8081/app?kuid=` + kuid + "&ctime=" + ctime + "&hash=" + coffee_hash
		
		if rcpt != "@alumni.ku.dk" {
			mail.Send(rcpt, body)
		}
		t, _ := template.ParseFiles("/home/dikukeys/Orkeren/DIKU-Keyserver/html_templates/reg_mail_sent.html")
		t.Execute(resp, p)
	} else if hex.EncodeToString(hash.GetHash(kuid, ctime)[:]) == coffee_hash {
		t, _ := template.ParseFiles("/home/dikukeys/Orkeren/DIKU-Keyserver/html_templates/public_key.html")
		t.Execute(resp, p)
	} else {
		resp.Write([]byte("<p>Not a valid link!</p>"))
	}

	type User struct {
		KUID   string
		PUBKEY string
	}
	if kuid != "" {
		if pubkey != "" {
			cuser := User{kuid, pubkey}
			tmpl, err := template.New("test").Parse("{{.KUID}} has submitted the public key {{.PUBKEY}}")
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(os.Stdout, cuser)
			if err != nil {
				panic(err)
			}
		}
	}
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}
