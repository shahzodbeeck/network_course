package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Error listening port")
	}
	defer listener.Close()
	fmt.Println("Server listening on  :8000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting pakage")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())
	buffer := make([]byte,4096)
	n,err := conn.Read(buffer)
	if err != nil {
	fmt.Println("Eror reading connection")
	return
	}
	fmt.Printf("Received : %s\n", string(buffer[:n]))
	response := []byte("Hello from go Server !\n")
	_, err = conn.Write(response)
	if err != nil{
	fmt.Println("Eror connection")
	return
	}
	


}