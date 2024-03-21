package chat

import (
	"godemo/internal/goconcurrency"
	"testing"
)

func TestChatServer(t *testing.T) {
	goconcurrency.StartChat()
}
