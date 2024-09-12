package main

import (
	"flag"
	"fmt"
)

var (
	listen = flag.Bool("l", false, "Listen")
	host   = flag.String("h", "localhost", "Host")
	port   = flag.Int("p", 0, "Port")
)

func main() {
	// parse flags
	flag.Parse()
	// launch server
	if *listen {
		startServer()
		return
	}
}

func startServer() {
	addr := fmt.Sprintf("%s:%d", *host, *port)
	// launch TCP server
	// listener, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	// if we can't launch server for some reason,
	// 	panic(err)
	// }
	fmt.Println(addr)
}
