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
			client.Write([]byte(msg))
		}
		client.Write([]byte("[" + currentTime + "]" + "[" + name + "]:"))
	}
}

func Connection(conn net.Conn, File []byte) {
	defer conn.Close()
	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	conn.Write([]byte(File))
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
	fmt.Println(len(Clients))
	if len(Clients) >= 2 {
		conn.Write([]byte("you cannot join the chat"))
		return
	}

	if len(msgs) != 0 {
		conn.Write([]byte(strings.Join(msgs, "") + "\n"))
		conn.Write([]byte(fmt.Sprintf("[%s] [%s]:", time.Now().Format(time.DateTime), Clients[conn])))
		// mssags = ""
	}

	BrodCast("\n"+Clients[conn]+" has joined our chat..."+"\n", conn)
	// mssags += "\n" + Clients[conn] + " has joined our chat..."
	for {
		msg, err := read.ReadString('\n')
		if err != nil {
			delete(Clients, conn)
			BrodCast("\n"+name+" has left our chat...\n", conn)
			break
		}
		if len(msg) == 1 {
			conn.Write([]byte("you can't enter in empty message\n"))
			conn.Write([]byte(fmt.Sprintf("[%s] [%s]:", time.Now().Format(time.DateTime), Clients[conn])))
			continue
		}
		if !isPrintable(msg[:len(msg)-1]) {
			conn.Write([]byte("you just entered in invalid text\n"))
			conn.Write([]byte(fmt.Sprintf("[%s] [%s]:", time.Now().Format(time.DateTime), Clients[conn])))
			continue
		}
		BrodCast(fmt.Sprintf("\n[%s] [%s]:%s", time.Now().Format(time.DateTime), Clients[conn], msg), conn)
	}
}
