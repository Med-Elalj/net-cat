package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"netcat/connection"
)

const (
	TYPE = "tcp"
)

func main() {
	File, err := os.ReadFile("../connection/logo.txt")
	if err != nil {
		panic(err)
	}
	PORT := connection.CheckPort()
	if PORT == "invalid port" {
		fmt.Println("invalid port")
		return
	}
	listener, err := net.Listen(TYPE, ":"+PORT)
	if err != nil {
		fmt.Println("you cannot connected")
		return
	}
	fmt.Println("Starting server at localhost " + PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection from client: %s", err)
			continue
		}
		go connection.Connection(conn, File)
	}
}
