package connection

import (
	"bufio"
	"net"
	"strings"
	"time"
)

func BrodCast(msg string, conn net.Conn) {
	currentTime := time.Now().Format(time.DateTime)
	for client, name := range Clients {
		if client != conn {
			client.Write([]byte(msg))
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
		conn.Write([]byte(name + " invalid used." + "\n"))
		return
	}
	Clients[conn] = name
	if len(Clients) > 10 {
		conn.Write([]byte("you cannot join the chat"))
		return
	}
	BrodCast("\n"+Clients[conn]+" has joined our chat..."+"\n", conn)
	for {
		msg, err := read.ReadString('\n')
		if err != nil {
			BrodCast("\n"+name+" has left our chat...\n", conn)
			break
		}
		if len(msg) < 2 {
			conn.Write([]byte("you can't enter in empty message\n"))
			conn.Write([]byte("[" + time.Now().Format(time.DateTime) + "][" + Clients[conn] + "]:"))
			continue
		}
		if !isPrintable(msg[:len(msg)-1]) { //****
			conn.Write([]byte("you just entered in invalid text\n"))
			conn.Write([]byte("[" + time.Now().Format(time.DateTime) + "][" + Clients[conn] + "]:"))
			continue
		}
		BrodCast("\n"+"["+time.Now().Format(time.DateTime)+"]"+"["+Clients[conn]+"]:"+msg, conn)
	}
}
