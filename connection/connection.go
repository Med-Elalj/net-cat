package connection

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func BrodCast(msg string, conn net.Conn) {
	msgs = append(msgs, msg[:len(msg)-1])
	currentTime := time.Now().Format(time.DateTime)
	// lock the mutex before accessing the data
	mu.Lock()
	defer mu.Unlock()
	for client, name := range Clients {
		if client != conn {
			client.Write([]byte("\n" + msg))
		}
		client.Write([]byte(fmt.Sprintf("[%s][%s]:", currentTime, name)))
	}
}

func Connection(conn net.Conn, Logo []byte) {
	defer conn.Close()
	defer delete(Clients, conn)
	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	conn.Write([]byte(Logo))
	if len(Clients) > 10 {
		conn.Write([]byte("sorry server full cannot join the chat"))
		return
	}
	name := ""
	read := bufio.NewReader(conn)
	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		name, _ = read.ReadString('\n')
		name = strings.TrimSpace(name)
		if !Validname(name) {
			conn.Write([]byte(name + " invalid charachters used." + "\n"))
		} else {
			break
		}
	}
	mu.Lock()
	Clients[conn] = name
	mu.Unlock()
	defer BrodCast(name+" has left our chat...\n", conn)
	fmt.Println(len(Clients))
	if len(msgs) != 0 {
		conn.Write([]byte(strings.Join(msgs, "\n") + "\n"))
	}
	if len(Clients) != 1 {
		BrodCast(name+" has joined the chat...\n", conn)
	} else {
		conn.Write([]byte(fmt.Sprintf("[%s][%s]:", time.Now().Format(time.DateTime), name)))
	}
	for {
		msg, err := read.ReadString('\n')
		if err != nil {
			break
		}
		if len(msg) == 1 {
			conn.Write([]byte("you can't enter in empty message\n"))
			conn.Write([]byte(fmt.Sprintf("[%s][%s]:", time.Now().Format(time.DateTime), Clients[conn])))
			continue
		}
		if !isPrintable(msg[:len(msg)-1]) {
			conn.Write([]byte("you just entered an invalid text\n"))
			conn.Write([]byte(fmt.Sprintf("[%s][%s]:", time.Now().Format(time.DateTime), Clients[conn])))
			continue
		}
		BrodCast(fmt.Sprintf("[%s][%s]:%s", time.Now().Format(time.DateTime), Clients[conn], msg), conn)
	}
}
