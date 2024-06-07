package golib

import (
	"godemo/internal/golib/osdemo"
	"testing"
)

func TestSimpleOpenFile(t *testing.T) {
	err := osdemo.SimpleOpenFile()
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdvancedOpenFile(t *testing.T) {
	err := osdemo.AdvancedOpenFile()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleReadFile(t *testing.T) {
	err := osdemo.SimpleReadFile()
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdvancedReadFile(t *testing.T) {
	data, err := osdemo.AdvancedReadFile()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}

func TestIOReadFile(t *testing.T) {
	targetPath := "/data/GoDemo/internal/golib/osdemo/demo2.json"
	err := osdemo.IOReadFile(targetPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleWriteFile(t *testing.T) {
	err := osdemo.SimpleWriteFile()
	if err != nil {
		t.Fatal(err)
	}
	targetPath := "/data/GoDemo/internal/golib/osdemo/demo.json"
	err = osdemo.IOReadFile(targetPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdvancedWriteFile(t *testing.T) {
	err := osdemo.AdvancedWriteFile()
	if err != nil {
		t.Fatal(err)
	}
}

func TestIOWriteFile(t *testing.T) {
	err := osdemo.IOWriteFile()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleCopyFile(t *testing.T) {
	err := osdemo.SimpleCopyFile()
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdvancedCopyFile(t *testing.T) {
	originPath := "/data/GoDemo/internal/golib/osdemo/demo.json"
	targetPath := "/data/GoDemo/internal/golib/osdemo/demo2.json"

	err := osdemo.AdvancedCopyFile(originPath, targetPath)
	if err != nil {
		t.Fatal(err)
	}

}

func TestIOCopyFile(t *testing.T) {
	originPath := "/data/GoDemo/internal/golib/osdemo/demo.json"
	targetPath := "/data/GoDemo/internal/golib/osdemo/demo2.json"

	err := osdemo.IOCopyFile(originPath, targetPath)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDeleteFile(t *testing.T) {
	originPath := "/data/GoDemo/internal/golib/osdemo/demo.json"

	err := osdemo.DeleteFile(originPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteDir(t *testing.T) {
	originPath := "/data/GoDemo/internal/golib/osdemo/test"

	err := osdemo.DeleteDir(originPath)
	if err != nil {
		t.Fatal(err)
	}
}
