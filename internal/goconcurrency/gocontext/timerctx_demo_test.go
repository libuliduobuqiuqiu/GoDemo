package gocontext

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWaitForCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	if err := WaitForCancellation(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("deadline error = %v", err)
	}

	canceled, stop := context.WithCancel(context.Background())
	stop()
	if err := WaitForCancellation(canceled); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancel error = %v", err)
	}
}

func TestGenNumStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	values := GenNum(ctx)
	for want := 0; want < 3; want++ {
		if got := <-values; got != want {
			t.Fatalf("value %d = %d", want, got)
		}
	}
	cancel()
	deadline := time.After(time.Second)
	for {
		select {
		case _, ok := <-values:
			if !ok {
				return
			}
		case <-deadline:
			t.Fatal("generator did not stop after cancellation")
		}
	}
}
