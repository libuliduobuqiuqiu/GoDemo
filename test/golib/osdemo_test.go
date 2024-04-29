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
