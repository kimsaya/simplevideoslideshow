package main

import (
	"fmt"
	"log"
	"net/http"
)

var c int = 0

func main() {
	// c := 0

	http.HandleFunc("/hello", webl)
	http.HandleFunc("/counter", countr)
	http.HandleFunc("/fileTest", sfile)
	http.HandleFunc("/urlInfo", echoUrlInfo)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func webl(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Avatar")
	log.Printf("Page Visited\n")
}
func countr(w http.ResponseWriter, req *http.Request) {
	c++
	fmt.Fprintf(w, "Hi there<br/>You're visiter #%v", c)
	log.Printf("Countr visited %v times\n", c)
}
func sfile(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "test2")
	log.Printf("Served File\n")
}
func echoUrlInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hi there<br/>The URL you're visiting is made up of these component parts!<br/>Gocode!:<br/>%T<br/>%#v", req, req)
	log.Printf("Hit URL Info page\n")
}
