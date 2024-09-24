package main

import (
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

func TestServerResponse(t *testing.T) {
	fmt.Println("azer")
	// t.Fail()
	PORT = "1234"
	go main()
	// Give the server a moment to start
	time.Sleep(50 * time.Millisecond)
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	// Send data to the server
	// message := "Hello, Server!"
	// _, err = conn.Write([]byte(message))
	// if err != nil {
	// 	t.Fatalf("Error sending data: %v", err)
	// }
	expected := "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |\\dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n[ENTER YOUR NAME]: "

	// Read the response
	buffer := make([]byte, 1024)
	response := ""
	for len(response) < len(expected) {
		_, err = conn.Read(buffer)
		if err != nil {
			t.Fatalf("Error reading from server: %v", err)
		}
		time.Sleep(100 * time.Millisecond)
		response += strings.ReplaceAll(string(buffer[:len(expected)]), "\x00", "")
		clear(buffer)
	}
	if response != expected {
		t.Errorf("Expected response %q, got %q", expected, response)
	}
	fmt.Println(response)
}
