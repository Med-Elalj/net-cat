package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	// parse flags
	// launch server
	//	if *listen {
	startServer()
	return
	//	}
	// if len(flag.Args()) < 2 {
	// 	fmt.Println("Hostname and port required")
	// 	return
	// }
	// serverHost := flag.Arg(0)
	// serverPort := flag.Arg(1)
	startClient(fmt.Sprintf("%s:%s", HOST, PORT))
}

func startServer() {
	// donc addr is server's address
	// addr := fmt.Sprintf("%s:%d", *host, *port)******************
	// launch TCP server
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
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
}

func startClient(addr string) {
	// net.Dial connect to a TCP server running on localhost at port 1234. (use for connect with our server)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Can't connect to server: %s\n", err)
		return
	}
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
	}
}
