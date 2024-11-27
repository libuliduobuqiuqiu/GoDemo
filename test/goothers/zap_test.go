package goothers

import (
	"godemo/internal/goothers/zapdemo"
	"testing"
)

func TestZapLogging(t *testing.T) {
	zapdemo.UseZapLogging()
	// zapdemo.UseZapExample()
	// zapdemo.UseZapProduction()
	// zapdemo.UseZapProductionSuger()
}
