package gochannel

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client chan<- string

type ClientMessage struct {
	Client   client
	Name     string
	Messages string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan ClientMessage)
	users    = make(map[string]struct{})
)

func StartChat() {
	fmt.Println("Start Listen 127.0.0.1:8080")
	server, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	go brocaster()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func brocaster() {
	var clients = make(map[client]bool)

	for {
		select {
		case cm := <-messages:
			for c := range clients {
				if c != cm.Client {
					c <- cm.Messages
				}
			}
			fmt.Println(cm.Messages)
		case c := <-entering:
			clients[c] = true
		case c := <-leaving:
			delete(clients, c)
			close(c)
		}
	}

}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go readWriter(conn, ch)

	buffer := bufio.NewScanner(conn)
	reader := bufio.NewReader(conn)
	cm := ClientMessage{Client: ch}

	// 输入对应的用户名称
loop:
	for {
		ch <- "Send your name: "
		remoteName, err := reader.ReadString('\n')
		remoteName = strings.Split(remoteName, "\n")[0]
		if err != nil {
			fmt.Println(err)
			continue loop
		}

		if _, ok := users[remoteName]; ok {
			ch <- "Error: " + remoteName + " is Exist."
		} else {
			cm.Name = remoteName
			users[remoteName] = struct{}{}
			break loop
		}
	}

	remoteAddr := conn.RemoteAddr().String()
	ch <- "You IP Address: " + remoteAddr

	cm.Messages = cm.Name + " has arrived."
	messages <- cm
	entering <- ch

	for buffer.Scan() {
		cm.Messages = cm.Name + ": \n" + buffer.Text()
		messages <- cm
	}

	leaving <- ch
	cm.Messages = cm.Name + " has left."
	messages <- cm
	conn.Close()
}

func readWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}

}
