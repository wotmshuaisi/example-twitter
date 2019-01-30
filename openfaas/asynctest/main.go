package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/callback" || r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 NOT FOUND."))
		return
	}

	os.Stdout.Write([]byte("=================================\n"))
	os.Stdout.Write([]byte("REMOTE:\n"))
	os.Stdout.Write([]byte(r.RemoteAddr))
	os.Stdout.Write([]byte("\nHEADERS:\n"))
	for k, v := range r.Header {
		s := fmt.Sprintf("%s = %s\n", k, v[0])
		os.Stdout.Write([]byte(s))
	}
	os.Stdout.Write([]byte("BODY:\n"))
	defer r.Body.Close()
	data, _ := ioutil.ReadAll(r.Body)
	os.Stdout.Write(data)
	os.Stdout.Write([]byte("\n=================================\n"))

	w.WriteHeader(200)
}

func main() {
	log.Println("starting asynctest server...")
	if err := http.ListenAndServe(":3001", handler{}); err != nil {
		log.Fatal(err)
	}
}
