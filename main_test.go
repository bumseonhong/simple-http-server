package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	server := NewServer()
	server.HandleFunc("/", func(rw ResponseWriter, r Request) {
		rw.Headers(map[string]string{
			"Content-Type": "text/plain",
		})
		rw.SendByString("Hello World")
		rw.Close()
	})
	server.Listen(8080)

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if string(body) != "Hello World" {
		t.Fail()
	}

	server.Close()
}
