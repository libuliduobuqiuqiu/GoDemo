package golib

import (
	"godemo/internal/golib/compressdemo"
	"testing"
)

func TestExtractTGZ(t *testing.T) {
	err := compressdemo.ExtractTGZ("/data/Bak/ebook.tgz", "/data/Bak/")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateTGZ(t *testing.T) {
	err := compressdemo.CreateTGZ("/data/GoDemo", "/data/tmp/godemo.tgz")
	if err != nil {
		t.Fatal(err)
	}
}
