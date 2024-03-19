package goconcurrency

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

type ClientMessage struct {
	Client   client
	Messages string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan ClientMessage)
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

	cm := ClientMessage{Client: ch}

	remoteAddr := conn.RemoteAddr().String()
	ch <- "You IP Address: " + remoteAddr

	cm.Messages = remoteAddr + " has arrived."
	messages <- cm
	entering <- ch

	buffer := bufio.NewScanner(conn)
	for buffer.Scan() {
		cm.Messages = remoteAddr + ": " + buffer.Text()
		messages <- cm
	}

	leaving <- ch
	cm.Messages = remoteAddr + " has left."
	messages <- cm
	conn.Close()
}

func readWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}

}
