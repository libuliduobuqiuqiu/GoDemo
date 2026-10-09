//go:build integration

package gotool

import (
	"godemo/internal/gotool/pprofdemo"
	"testing"
)

func TestAnysisFib(t *testing.T) {
	pprofdemo.AnalysisFibByPprof()
}
