package golib

import (
	"godemo/internal/golib/compressdemo"
	"testing"
)

func TestCompressZipFile(t *testing.T) {
	dirPath := "/root/PythonScript/"
	err := compressdemo.CompressFiles(dirPath, "PythonScript.zip")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnCompressZipFile(t *testing.T) {
	zipPath := "/data/Bak/ebook.zip"
	bakPath := "/data/Bak/"
	err := compressdemo.UnCompressZip(zipPath, bakPath)
	if err != nil {
		t.Fatal(err)
	}

}
