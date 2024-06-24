package concurrency

import (
	"godemo/internal/goconcurrency/gosync"
	"testing"
)

func TestUseAtomic(t *testing.T) {
	gosync.UseAtomic()
}

func TestConcurrencyAdd(t *testing.T) {
	gosync.ConcurrencyAdd()
}
