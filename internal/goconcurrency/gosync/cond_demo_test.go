package gosync

import "testing"

func TestBroadcastToListeners(t *testing.T) {
	for run := 0; run < 5; run++ {
		got := BroadcastToListeners(20)
		if len(got) != 20 {
			t.Fatalf("listeners = %d, want 20", len(got))
		}
		for i, id := range got {
			if id != i {
				t.Fatalf("listener %d = %d", i, id)
			}
		}
	}
	if got := BroadcastToListeners(0); len(got) != 0 {
		t.Fatalf("zero listeners = %v", got)
	}
}
