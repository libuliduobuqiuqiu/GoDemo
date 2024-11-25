package goothers

import (
	"godemo/internal/goothers/logrusdemo"
	"testing"
)

func TestUseLogrus(t *testing.T) {
	logrusdemo.UseLogrus()
}
