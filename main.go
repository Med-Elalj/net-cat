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
	PORT = "8080"
	TYPE = "tcp"
)

var (
	clients = make(map[net.Conn]string)
	File    []byte
)

func main() {
	// default port should be 8989
	// ida dakhel xi wahed port rej3o howa lport
	// num, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	//port ghalet
	// }
	// num > 1024 < 66563
	startServer()
}

func startServer() {
	File, err := os.ReadFile("logo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server at localhost 8080")
	listener, err := net.Listen(TYPE, ":"+PORT)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection from client: %s", err)
			continue
		}
		go Connection(conn, File)
	}
}

func BrodCast(msg string, conn net.Conn) {
	currentTime := time.Now().Format(time.DateTime)
	for client, name := range clients {
		if client != conn {
			client.Write([]byte("[" + currentTime + "]" + msg))
		}
		client.Write([]byte("[" + currentTime + "]" + "[" + name + "]:"))
	}
}

func Connection(conn net.Conn, File []byte) {
	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	conn.Write([]byte(File))
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	read := bufio.NewReader(conn)
	name, _ := read.ReadString('\n')
	name = strings.TrimSpace(name)
	if !Validname(name) {
		conn.Write([]byte(name + "is already used." + "\n"))
		return
	}
	// check if the group is full
	clients[conn] = name
	BrodCast("\n"+clients[conn]+" has joined our chat..."+"\n", conn)
	message := make([]byte, 1024)
	for {
		leng, err := conn.Read(message)
		if err != nil {
			BrodCast("\n"+name+" has left our chat...\n", conn)
			break
		}
		if leng < 2 {
			conn.Write([]byte("you can't enter in empty message\n"))
			conn.Write([]byte("[" + time.Now().Format(time.DateTime) + "][" + clients[conn] + "]:"))
			continue //***
		}
		if !isPrintable(string(message[:leng-1])) {
			conn.Write([]byte("you just entered in invalid text\n"))
			conn.Write([]byte("[" + time.Now().Format(time.DateTime) + "][" + clients[conn] + "]:"))
			continue
		}
		BrodCast("["+clients[conn]+"]:"+string(message), conn)
	}
}

func isPrintable(s string) bool {
	for _, v := range s {
		if v > 126 || v < 32 {
			return false
		}
	}
	return true
}

func Validname(connname string) bool {
	for _, name := range clients {
		if connname == name {
			return false
		}
	}
	return true
}
