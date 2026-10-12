package golib

import (
	"godemo/internal/golib/osdemo"
	"os"
	"path/filepath"
	"testing"
)

func TestFileOperations(t *testing.T) {
	dir := t.TempDir()
	source := filepath.Join(dir, "demo.json")
	if err := os.WriteFile(source, []byte("hello, world\n"), 0600); err != nil {
		t.Fatal(err)
	}

	for _, run := range []struct {
		name string
		fn   func() error
	}{
		{"open", func() error { return osdemo.SimpleOpenFile(source) }},
		{"open advanced", func() error { return osdemo.AdvancedOpenFile(source) }},
		{"read", func() error { return osdemo.SimpleReadFile(source) }},
		{"read all", func() error { return osdemo.IOReadFile(source) }},
		{"write", func() error { return osdemo.SimpleWriteFile(filepath.Join(dir, "written.json")) }},
		{"write advanced", func() error { return osdemo.AdvancedWriteFile(dir) }},
		{"write io", func() error { return osdemo.IOWriteFile(filepath.Join(dir, "io.json")) }},
		{"copy", func() error { return osdemo.SimpleCopyFile(source, filepath.Join(dir, "copy.json")) }},
		{"copy advanced", func() error { return osdemo.AdvancedCopyFile(source, filepath.Join(dir, "advanced-copy.json")) }},
		{"copy io", func() error { return osdemo.IOCopyFile(source, filepath.Join(dir, "io-copy.json")) }},
	} {
		t.Run(run.name, func(t *testing.T) {
			if err := run.fn(); err != nil {
				t.Fatal(err)
			}
		})
	}

	data, err := osdemo.AdvancedReadFile(source)
	if err != nil || string(data) != "hello, world\n" {
		t.Fatalf("advanced read = %q, %v", data, err)
	}
	for _, name := range []string{"copy.json", "advanced-copy.json", "io-copy.json"} {
		copied, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil || string(copied) != string(data) {
			t.Fatalf("%s = %q, %v", name, copied, err)
		}
	}

	if err := osdemo.DeleteFile(filepath.Join(dir, "copy.json")); err != nil {
		t.Fatal(err)
	}
	deleteDir := filepath.Join(dir, "to-delete")
	if err := os.Mkdir(deleteDir, 0700); err != nil {
		t.Fatal(err)
	}
	if err := osdemo.DeleteDir(deleteDir); err != nil {
		t.Fatal(err)
	}
}
