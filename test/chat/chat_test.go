package chat

import (
	"sunrun/internal/goconcurrency"
	"testing"
)

func TestChatServer(t *testing.T) {
	goconcurrency.StartChat()
}
