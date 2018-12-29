package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	fmt.Printf("Hello World!")
}

type ResponseWriter struct {
}

func (ResponseWriter) Headers(kv map[string]string) {

}

func (ResponseWriter) SendByString(data string) {

}

func (ResponseWriter) Close() {

}

type Request struct {
}

type HttpServer interface {
	HandleFunc(string, func(ResponseWriter, Request))
	Close()
	Listen(int)
}

type HttpServerImpl struct {
	handlers map[string]func(ResponseWriter, Request)
}

func (h HttpServerImpl) HandleFunc(pattern string, cb func(rw ResponseWriter, r Request)) {
	h.handlers[pattern] = cb
}

func (HttpServerImpl) Close() {

}

func (h HttpServerImpl) Listen(port int) {
	net, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

	go handleConn(net)
}

// Header
func handleConn(net net.Listener) {
	for {
		conn, err := net.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(fmt.Sprintf("connect: %v", conn))

		// FIXME: 개행문자까지 header를 읽고 Header struct를 생성
		content := make([]byte, 1000)
		readN, err := conn.Read(content)
		if err != nil {
			log.Output(2, fmt.Sprintf("%v", err))
		}
		fmt.Println(readN)
		fmt.Println(string(content))
	}
}

func NewServer() HttpServer {
	server := HttpServerImpl{}
	server.handlers = make(map[string]func(ResponseWriter, Request))
	return server
}
