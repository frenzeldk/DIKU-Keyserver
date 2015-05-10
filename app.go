package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Orkeren/DIKU-Keyserver/golibs/hash" // This is our hash function
	"github.com/Orkeren/DIKU-Keyserver/golibs/mail" // This is our mail function, it does hello
	//	"github.com/Orkeren/DIKU-Keyserver/golibs/dbs" // This is our sqlite function
	"html/template"
	//"io/ioutil"
	"net"
	"net/http"
	"net/http/fcgi"
	//"os"
	"strconv"
	//"strings"
	"regexp"
	"time"
)

type FastCGIServer struct{}
type Page struct {
	Title string
}
type User struct {
	KUID   string
	PUBKEY string
}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	titel := req.URL.Path[len("/"):]
	p, _ := loadPage(titel)
	fmt.Println(titel)
	//fmt.Println(req.URL) // Dette viser bare hvordan man får en URL fra req
	//kuid is the KU ID of the student
	kuid := req.FormValue("kuid")
	//ctime is the time of creation of the link (as unix time)
	ctime := req.FormValue("ctime")
	//hash is the padded sha3-512 hash of kuid & ctime)
	coffee_hash := req.FormValue("hash")

	pubkey := req.FormValue("pubkey")

	if !validKUID(kuid) && kuid != "" {
		resp.Write([]byte("<p>Not a valid link!</p>"))
	}
	rcpt := kuid + "@alumni.ku.dk"

	if kuid == "" && pubkey == "" {
		t, _ := template.ParseFiles("/home/dikukeys/Orkeren/DIKU-Keyserver/html_templates/create_link.html")
		t.Execute(resp, p)
	} else if coffee_hash == "" {
		ctime = strconv.FormatInt(time.Now().Unix(), 10)
		coffee_hash = hex.EncodeToString(hash.GetHash(kuid + ctime)[:])

		//mailbody is the plaintext body of the email.
		mailbody := `English below

Velkommen til dikukeys. For at afslutte registreringen, tryk venligst p&#229; dette link:
http://dikukeys.dk/?kuid=` + kuid + "&ctime=" + ctime + "&hash=" + coffee_hash + `

Welcome to DIKU Keys. To register in the DIKU Keys system please follow this link:
http://dikukeys.dk/?kuid=` + kuid + "&ctime=" + ctime + "&hash=" + coffee_hash

		if rcpt != "@alumni.ku.dk" {
			mail.Send(rcpt, mailbody)
			t, _ := template.ParseFiles("/home/dikukeys/Orkeren/DIKU-Keyserver/html_templates/reg_mail_sent.html")
			t.Execute(resp, p)
		}

	} else if hex.EncodeToString(hash.GetHash(kuid + ctime)[:]) == coffee_hash {
		t, _ := template.ParseFiles("/home/dikukeys/Orkeren/DIKU-Keyserver/html_templates/public_key.html")
		t.Execute(resp, p)
	} else {
		resp.Write([]byte("<p>Not a valid link!</p>"))
	}

	if kuid == "" && pubkey != "" {
		//cuser := User{kuid, pubkey}
		t, _ := template.ParseFiles("/home/dikukeys/Orkeren/DIKU-Keyserver/html_templates/pub_key_succesful.html")
		t.Execute(resp, p)
	}
}

func validKUID(kuid string) (result bool) {
	regpatternKUID := "(?i)^[b-df-hj-np-tv-z]{3}\\d{3}$"
	regmatch, _ := regexp.MatchString(regpatternKUID, kuid)
	return regmatch
}

func loadPage(title string) (*Page, error) {
	return &Page{Title: title}, nil
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}
