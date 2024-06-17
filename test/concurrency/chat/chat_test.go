package chat

import (
	"godemo/internal/goconcurrency/gochannel"
	"testing"
)

func TestChatServer(t *testing.T) {
	gochannel.StartChat()
}
