package connection

import (
	"fmt"
	"net"
	"os"
)

var (
	Clients = make(map[net.Conn]string)
	File    []byte
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
	if len(connname) < 2 || !isPrintable(connname) {
		fmt.Println("hhh")
		return false
	}
	return true
}

func isnumeric(s string) bool {
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
		if !isnumeric(os.Args[1]) || len(os.Args[1]) > 5 {
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
