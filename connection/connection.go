package connection

import (
	"bufio"
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
	mu.Lock()
	Clients[conn] = name
	mu.Unlock()
	if len(Clients) > 10 {
		conn.Write([]byte("you cannot join the chat"))
		return
	}
	
	if len(msgs) != 0 {
		conn.Write([]byte(strings.Join(msgs, "") + "\n"))
		conn.Write([]byte("[" + time.Now().Format(time.DateTime) + "][" + Clients[conn] + "]:"))
		// mssags = ""
	}

	BrodCast("\n"+Clients[conn]+" has joined our chat..."+"\n", conn)
	// mssags += "\n" + Clients[conn] + " has joined our chat..."
	for {
		msg, err := read.ReadString('\n')
		if err != nil {
			BrodCast("\n"+name+" has left our chat...\n", conn)
			break
		}
		if len(msg) == 1 {
			conn.Write([]byte("you can't enter in empty message\n"))
			conn.Write([]byte("[" + time.Now().Format(time.DateTime) + "][" + Clients[conn] + "]:"))
			continue
		}
		if !isPrintable(msg[:len(msg)-1]) {
			conn.Write([]byte("you just entered in invalid text\n"))
			conn.Write([]byte("[" + time.Now().Format(time.DateTime) + "][" + Clients[conn] + "]:"))
			continue
		}
		BrodCast("\n"+"["+time.Now().Format(time.DateTime)+"]"+"["+Clients[conn]+"]:"+msg, conn)
		// msg = strings.TrimSpace(msg)
		// msgs = append(msgs, "\n"+"["+time.Now().Format(time.DateTime)+"]"+"["+Clients[conn]+"]:"+msg)
		// for i := 0; i < len(msgs); i++ {
		// 	mssags += msgs[i]
		// }
	}
}
