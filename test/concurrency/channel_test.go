package concurrency

import (
	"godemo/internal/goconcurrency/gochannel"
	"testing"
)

func TestUseChannel(t *testing.T) {
	gochannel.UseChannel()
}

func TestUseChannelSelect(t *testing.T) {
	gochannel.UseChannelSelect()
}

func TestUseChannelDone(t *testing.T) {
	gochannel.UseChannelDone()
}
