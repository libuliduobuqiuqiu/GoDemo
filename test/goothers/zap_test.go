package goothers

import (
	"godemo/internal/goothers/zapdemo"
	"testing"
)

func TestZapLogging(t *testing.T) {
	err := zapdemo.UseZapLogging()
	if err != nil {
		t.Fatal(err)
	}
	// zapdemo.UseZapExample()
	// zapdemo.UseZapProduction()
	// zapdemo.UseZapProductionSuger()
}
