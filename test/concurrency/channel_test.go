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

func TestRecvClosedChannel(t *testing.T) {
	gochannel.RecvClosedChannel()
}

func TestUseChannelPrintUser(t *testing.T) {
	gochannel.UseChannelPrintUser()
}

func TestUseLimitGoroutine(t *testing.T) {
	var tasks []string
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			tasks = append(tasks, "hello,world")
		} else {
			tasks = append(tasks, "hello,main")
		}
	}

	gochannel.UseLimitGoroutine(5, tasks)
}

func TestUseChannelTicker(t *testing.T) {
	gochannel.UseChannelTicker()
}

func TestUseChannelTimer(t *testing.T) {
	gochannel.UseChannelTimer()
}

func TestUseProducerConsumer(t *testing.T) {
	gochannel.UseProducerConsumer()
}

func TestUseChannelClosedGracefully(t *testing.T) {
	gochannel.UseChannelClosedGracefully()
}

func TestUseChannelRange(t *testing.T) {
	gochannel.UseChannelRange()
}
