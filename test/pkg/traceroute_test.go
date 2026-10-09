//go:build integration

package pkg

import (
	"fmt"
	"godemo/pkg/traceroute"
	"testing"
	"time"
)

func TestTraceroute(t *testing.T) {
	start := time.Now()
	results, err := traceroute.Traceroute("192.168.23.64", traceroute.DefaultConfig())
	if err != nil {
		t.Fatalf("traceroute failed: %v", err)
	}

	fmt.Printf("\n追踪结果:\n")
	for _, result := range results {
		if result.Success {
			fmt.Printf("%d  %s  %v\n", result.Hop, result.IP, result.RTT)
		} else {
			fmt.Printf("%d  * * *\n", result.Hop)
		}
	}

	fmt.Printf("\n总耗时: %v\n", time.Since(start))
}
