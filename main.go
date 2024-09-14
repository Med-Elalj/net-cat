package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	// launch server
	startServer()
}

func startServer() {
	// func net.Listen  It is used to create a network listener that waits for incoming connections on a specified network address and port
	// net.Listen("tcp" the network type , addr the address and port to listen on)
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		// if we can't launch server for some reason,
		panic(err)
	}
	// the infinite loop accepts TCP connections from clients and processes these connections in separate goroutines
	for {
		// The listener.Accept() method in Go is used to accept a new incoming connection on a network listener.
		// after each successful connection ( by processClient(conn) ) , server returns to the listener.Accept() and gets ready for the next client.
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection from client: %s", err)
		} else {
			// Handle the connection in a new goroutine
			go processClient(conn)
		}
	}
}

// processing's function
// io.Copy way to copy data from one source to a destination.
// conn from which the data will be read (Copying Data from a Network Connection)
// os.Stdout where the data will be copied to. type : standard output which is typically the terminal
func processClient(conn net.Conn) {
	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	conn.Write([]byte("Enter your name:"))
	reder := bufio.NewReader(conn)
	user, _ := reder.ReadString('\n')
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	conn.Write([]byte(currentTime))
	conn.Write([]byte(user))
}
