//go:build integration

package gotool

import (
	"godemo/internal/gotool/pprofdemo"
	"os"
	"path/filepath"
	"testing"
)

func TestAnalysisFibProfile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "cpu.profile")
	if err := pprofdemo.AnalysisFibByPprof(path); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Size() == 0 {
		t.Fatal("CPU profile is empty")
	}
}
