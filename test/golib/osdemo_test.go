//go:build integration

package golib_test

import (
	"godemo/internal/golib/osdemo"
	"os"
	"path/filepath"
	"testing"
)

func TestUseFileUtil(t *testing.T) {
	if err := osdemo.UseFileUtil(); err != nil {
		t.Fatal(err)
	}
}

func TestUseBufioMaxInsert(t *testing.T) {
	osdemo.UseBufioMaxInsert()
}

func TestHandleFile(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "demo.txt")
	if err := os.WriteFile(filePath, []byte("hello"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := osdemo.HandleFile(filePath); err != nil {
		t.Fatal(err)
	}
}

func TestPrintFilePath(t *testing.T) {

	if err := osdemo.PrintFilePath(t.TempDir()); err != nil {
		t.Fatal(err)
	}

}
