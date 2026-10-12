//go:build integration

package gotool

import (
	"godemo/internal/gotool/compressdemo"
	"os"
	"path/filepath"
	"testing"
)

func TestCompressZipFile(t *testing.T) {
	dirPath := os.Getenv("GO_DEMO_ARCHIVE_SOURCE_DIR")
	if dirPath == "" {
		t.Skip("set GO_DEMO_ARCHIVE_SOURCE_DIR to a local directory")
	}
	err := compressdemo.CompressFiles(dirPath, filepath.Join(t.TempDir(), "demo.zip"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnCompressZipFile(t *testing.T) {
	zipPath := os.Getenv("GO_DEMO_ARCHIVE_ZIP")
	if zipPath == "" {
		t.Skip("set GO_DEMO_ARCHIVE_ZIP to a local .zip file")
	}
	err := compressdemo.UnCompressZip(zipPath, t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

}
