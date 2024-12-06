package netdemo

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func ReadChat(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}

func StartTcpChat() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	go ReadChat(conn)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("you:")
		fmt.Fprintln(conn, input.Text())
	}
}
