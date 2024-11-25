package goothers

import (
	"godemo/internal/goothers/zapdemo"
	"testing"
)

func TestZapLogging(t *testing.T) {
	zapdemo.UseZapLogging()
}

func TestZapFasterLogging(t *testing.T) {
	zapdemo.UseFasterZapLogging()
}

func TestZapExampleLogging(t *testing.T) {
	zapdemo.UseZapExampleLogging()
}
