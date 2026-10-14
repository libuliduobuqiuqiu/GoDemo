package osdemo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFileDataUsesOnlyReadBytes(t *testing.T) {
	for _, content := range []string{"", "ok"} {
		path := filepath.Join(t.TempDir(), "input.txt")
		if err := os.WriteFile(path, []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
		file, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		got, readErr := readFileData(file)
		closeErr := file.Close()
		if readErr != nil || closeErr != nil || got != content {
			t.Fatalf("read %q: got %q, read error %v, close error %v", content, got, readErr, closeErr)
		}
	}
}
