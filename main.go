package main

import (
	"bufio"
	"log"
	"net"
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
		// fmt.Println(client, name)
		if client != conn {
			client.Write([]byte(msg))
		}
		currentTime := time.Now().Format(time.DateTime)
		client.Write([]byte("[" + currentTime + "]" + "[" + name + "]:"))
	}
}

func Connection(conn net.Conn) {
	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	conn.Write([]byte("Enter your name:"))
	currentTime := time.Now().Format(time.DateTime)
	read := bufio.NewReader(conn)
	name, _ := read.ReadString('\n')
	name = strings.TrimSpace(name)
	clients[conn] = name
	BrodCast("\n"+clients[conn]+" has joined"+"\n", conn)
	for {
		msg, err := read.ReadString('\n')
		if err != nil {
			BrodCast("\n"+name+" has left\n", conn)
			break
		}
		BrodCast("\n"+"["+currentTime+"]"+"["+clients[conn]+"]:"+msg, conn)
	}
}
