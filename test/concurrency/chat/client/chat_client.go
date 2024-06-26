package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/libuliduobuqiuqiu/chat-client"
)

func readConn(c net.Conn) {
	fromChatData := bufio.NewScanner(c)

	for fromChatData.Scan() {
		fmt.Println(fromChatData.Text())
	}
}

func startChatClient() {
	fmt.Println("start connect chat server...")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect successfully.")

	go readConn(conn)

	fmt.Println("you: ")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("you: ")
		fmt.Fprintln(conn, input.Text())
	}
}

func main() {
	chat.StartChatClient("tcp", "127.0.0.1:8080")
}
