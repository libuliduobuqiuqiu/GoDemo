package chat

import (
	"testing"

	"github.com/libuliduobuqiuqiu/chat-client"
)

func TestChatClient(t *testing.T) {
	chat.StartChatClient("tcp", "127.0.0.1:8080")
}
