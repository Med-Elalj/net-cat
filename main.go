package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

var clients = make(map[net.Conn]string)

func main() {
	startServer()
}

func startServer() {
	fmt.Println("Starting server at localhost 8080")
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection from client: %s", err)
		}
		go Connection(conn)
	}
}

func BrodCast(msg string, conn net.Conn) {
	for client, name := range clients {
		if client != conn {
			client.Write([]byte(msg))
		}
		currentTime := time.Now().Format(time.DateTime)
		client.Write([]byte("[" + currentTime + "]" + "[" + name + "]:"))
	}
}

func Connection(conn net.Conn) {
	currentTime := time.Now().Format(time.DateTime)
	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	file, _ := os.ReadFile("logo.txt")
	conn.Write([]byte(file))
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	read := bufio.NewReader(conn)
	name, _ := read.ReadString('\n')
	name = strings.TrimSpace(name)
	clients[conn] = name
	BrodCast("\n"+clients[conn]+" has joined our chat..."+"\n", conn)
	for {
		msg, err := read.ReadString('\n')
		if len(msg) == 0 {
			fmt.Println("pooo")
			continue
		}
		if err != nil {
			BrodCast("\n"+name+" has left our chat...\n", conn)
			break
		}
		
		for i := 0; i < len(msg); i++ {
			if string(msg[i]) == "^[[D" || string(msg[i]) == "^[[B" || string(msg[i]) == "^[[A" || string(msg[i]) == "^[[C" {
				continue
			}
		}
		BrodCast("\n"+"["+currentTime+"]"+"["+clients[conn]+"]:"+msg, conn)
	}
}
