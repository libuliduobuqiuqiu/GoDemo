//go:build integration

package concurrency

import (
	"godemo/internal/goconcurrency/gosync"
	goconcurrency "godemo/internal/goconcurrency/gosyncmap"
	"os"
	"testing"
)

func TestUseSyncCond(t *testing.T) {
	if os.Getenv("GO_DEMO_INTEGRATION") != "1" {
		t.Skip("waits for an external signal; set GO_DEMO_INTEGRATION=1 to run")
	}
	gosync.UseSyncCond()
}

func TestSyncMap(t *testing.T) {

	goconcurrency.ConcurrentSyncMap()

}

func TestDeviceSyncMap(t *testing.T) {
	goconcurrency.DeviceExecCommands()
}
