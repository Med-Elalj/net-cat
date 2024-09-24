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

var PORT string

func main() {
	File, err := os.ReadFile("../connection/logo.txt")
	if err != nil {
		panic(err)
	}
	if PORT == "invalid port" {
		fmt.Println("invalid port")
		return
	}
	listener, err := net.Listen(TYPE, ":"+PORT)
	if err != nil { //**********
		fmt.Println("Error strating server :", err)
		return
	}
	fmt.Println("Starting server at localhost " + PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection from client: %s", err)
			conn.Close()
			continue
		}
		go connection.Connection(conn, File)
	}
}

func init() {
	if len(os.Args) == 2 {
		if !connection.Isnumeric(os.Args[1]) || len(os.Args[1]) > 5 {
			PORT = "invalid port"
		}
		PORT = os.Args[1]
	} else if len(os.Args) == 1 {
		PORT = "8989"
	} else {
		PORT = ("invalid port")
	}
}
