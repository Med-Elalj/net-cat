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

var msg string

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
		conn.Write([]byte("Welcome to TCP-Chat!\n"))
		conn.Write([]byte("Enter your name:"))
		read := bufio.NewReader(conn)
		name, _ := read.ReadString('\n')
		name = strings.TrimSpace(name)
		clients[conn] = name
		go BrodCast(conn)
	}
}

func BrodCast(conn net.Conn) {
	for {
		for client, name := range clients {
			reder := bufio.NewReader(client)
			if client != conn {
				currentTime := time.Now().Format(time.DateTime)
				client.Write([]byte("\n" + "[" + currentTime + "]" + "[" + clients[conn] + "]:" + msg))
			} else {
				currentTime := time.Now().Format("2006-01-02 15:04:05")
				client.Write([]byte("\n" + "[" + currentTime + "]" + "[" + string(name) + "]:"))
				msg, _ = reder.ReadString('\n')
			}
		}
	}
}
