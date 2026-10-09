//go:build integration

package concurrency

import (
	"godemo/internal/goconcurrency/gosync"
	goconcurrency "godemo/internal/goconcurrency/gosyncmap"
	"testing"
)

func TestUseSyncCond(t *testing.T) {
	gosync.UseSyncCond()
}

func TestSyncMap(t *testing.T) {

	goconcurrency.ConcurrentSyncMap()

}

func TestDeviceSyncMap(t *testing.T) {
	goconcurrency.DeviceExecCommands()
}
