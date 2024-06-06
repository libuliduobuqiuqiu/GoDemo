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
	err := osdemo.IOReadFile()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleWriteFile(t *testing.T) {
	err := osdemo.SimpleWriteFile()
	if err != nil {
		t.Fatal(err)
	}

	err = osdemo.IOReadFile()
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
