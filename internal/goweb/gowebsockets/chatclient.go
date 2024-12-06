package gowebsockets

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func SendMessage(conn *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		conn.WriteMessage(websocket.TextMessage, scanner.Bytes())
	}
}

func StartConnectWS() {
	var name string
	fmt.Println("Please enter your chat's name: ")
	fmt.Scanln(&name)

	chatUrl := "ws://127.0.0.1:8090/chat/join?name=" + name
	conn, _, err := websocket.DefaultDialer.Dial(chatUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	go SendMessage(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(message))
	}
}
