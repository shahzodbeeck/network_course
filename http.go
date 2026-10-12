package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Proto   string
	Headers map[string]string
	Body    string
}

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Portni band qila olmadi")
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server localhost:8080 da ishga tushdi")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ulanishda hatoli bo'ldi")
			continue
		}
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	requestLine, _ := reader.ReadString('\n')
	requestLine = strings.TrimSpace(requestLine)
	parts := strings.SplitN(requestLine, " ", 3)
	if len(parts) < 3 {
		fmt.Println("Buzuq request yuborilgan")
		return
	}
	request := Request{
		Method:  parts[0],
		Path:    parts[1],
		Proto:   parts[2],
		Headers: make(map[string]string),
		Body:    "",
	}
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		headerParts := strings.SplitN(line, ": ", 2)
		if len(headerParts) == 2 {
			key := strings.TrimSpace(headerParts[0])
			value := strings.TrimSpace(headerParts[1])
			request.Headers[key] = value
		}
	}

	//conn.Write([]byte("Http/1.1 200 OK\r\n Content-type: text/html \r\n\r\n <h1>Hello world</h1>"))
	sendResponse(conn, 200, "OK", "text/html", "<h1>Hello world</h1>")
}

func sendResponse(conn net.Conn, statusCode int, statusMessage string, contentType string, body string) {
	contentLenght := len(body)
	responseHeader := fmt.Sprintf(
	"HTTP/1.1 %d %s\r\n"+
	"Content-type:text/html\r\n"+
	"Content-lenght: %d\r\n"+
	"Connection:close \r\n"+
	"\r\n",
	statusCode,statusMessage,contentLenght,
	)
	responseHeader =responseHeader+body
	conn.Write([]byte(responseHeader))
}
