//go:build integration

package goothers

import (
	"godemo/internal/goothers"
	"testing"
)

func TestParseP12(t *testing.T) {
	var err error
	password := "111111"
	// err := goothers.ParseSM2P12("sm2enc.p12", password)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	//
	// err = goothers.ParseSM2P12Bydata("sm2sign.p12", password)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	//
	err = goothers.ParseSM2P12("down.p12", password)
	if err != nil {
		t.Fatal(err)
	}
}
