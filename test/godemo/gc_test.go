package godemo

import (
	"godemo/internal/godemo"
	"testing"
)

func TestUseAllocate(t *testing.T) {
	if got := godemo.UseAllocate(); got != 496 {
		t.Fatalf("allocation checksum = %d, want 496", got)
	}
}
