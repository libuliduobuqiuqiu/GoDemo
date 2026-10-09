package golib_test

import (
	"os"
	"path/filepath"
	"testing"

	"godemo/internal/golib/osdemo"
)

func TestCopyFileInTempDir(t *testing.T) {
	for _, tc := range []struct {
		name string
		copy func(string, string) error
	}{
		{name: "advanced", copy: osdemo.AdvancedCopyFile},
		{name: "io", copy: osdemo.IOCopyFile},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			source := filepath.Join(dir, "source.txt")
			target := filepath.Join(dir, "target.txt")
			want := []byte("hello, world\n")
			if err := os.WriteFile(source, want, 0600); err != nil {
				t.Fatal(err)
			}
			if err := tc.copy(source, target); err != nil {
				t.Fatal(err)
			}
			got, err := os.ReadFile(target)
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != string(want) {
				t.Fatalf("copied content = %q, want %q", got, want)
			}
		})
	}
}
