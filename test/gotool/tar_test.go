//go:build integration

package gotool

import (
	"godemo/internal/gotool/compressdemo"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractTGZ(t *testing.T) {
	archive := os.Getenv("GO_DEMO_ARCHIVE_TGZ")
	if archive == "" {
		t.Skip("set GO_DEMO_ARCHIVE_TGZ to a local .tgz file")
	}
	err := compressdemo.ExtractTGZ(archive, t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateTGZ(t *testing.T) {
	source := os.Getenv("GO_DEMO_ARCHIVE_SOURCE_DIR")
	if source == "" {
		t.Skip("set GO_DEMO_ARCHIVE_SOURCE_DIR to a local directory")
	}
	err := compressdemo.CreateTGZ(source, filepath.Join(t.TempDir(), "demo.tgz"))
	if err != nil {
		t.Fatal(err)
	}
}
