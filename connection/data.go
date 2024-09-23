package connection

import (
	"net"
	"os"
	"sync"
)

var (
	Clients = make(map[net.Conn]string)
	File    []byte
	msgs    []string
	// create a mutex by declaring a variable of type sync.Mutex
	mu sync.Mutex
)

func isPrintable(s string) bool {
	for _, v := range s {
		if v > 126 || v < 32 {
			return false
		}
	}
	return true
}

func Validname(connname string) bool {
	for _, name := range Clients {
		if connname == name {
			return false
		}
	}
	if len(connname) == 0 || !isPrintable(connname) {
		return false
	}
	return true
}

func Isnumeric(s string) bool {
	for _, v := range s {
		if v < 48 || v > 57 {
			return false
		}
	}
	return true
}

func CheckPort() string {
	var PORT string
	if len(os.Args) == 2 {
		if !Isnumeric(os.Args[1]) || len(os.Args[1]) > 5 {
			return "invalid port"
		}
		PORT = os.Args[1]
	} else if len(os.Args) == 1 {
		PORT = "8989"
	} else {
		return ("invalid port")
	}
	return PORT
}
