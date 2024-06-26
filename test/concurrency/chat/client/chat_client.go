package main

import (
	"github.com/libuliduobuqiuqiu/chat-client"
)

func main() {
	chat.StartChatClient("tcp", "127.0.0.1:8080")
}
