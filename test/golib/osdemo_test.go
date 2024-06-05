package golib_test

import (
	"godemo/internal/golib/osdemo"
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
	if err := osdemo.HandleFile(); err != nil {
		t.Fatal(t)
	}
}
