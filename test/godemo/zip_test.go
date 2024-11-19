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
