//go:build integration

package golib

import (
	"testing"

	"godemo/internal/golib/timedemo"
)

func TestRateLimite(t *testing.T) {
	timedemo.GenMultiRequest()
}
