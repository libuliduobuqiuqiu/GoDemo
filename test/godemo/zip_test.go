package godemo

import (
	"godemo/internal/godemo"
	"testing"
)

func TestCompressZipFile(t *testing.T) {
	dirPath := "/root/PythonScript/"
	err := godemo.CompressFiles(dirPath, "PythonScript.zip")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnCompressZipFile(t *testing.T) {
	zipPath := "/data/Bak/ebook.zip"
	bakPath := "/data/Bak/"
	err := godemo.UnCompressZip(zipPath, bakPath)
	if err != nil {
		t.Fatal(err)
	}

}
