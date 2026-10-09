package gochannel

import (
	"context"
	"errors"
	"testing"
)

func TestSelectFirst(t *testing.T) {
	numbers := make(chan int, 1)
	numbers <- 42
	value, err := SelectFirst(context.Background(), numbers, nil)
	if err != nil || value != "42" {
		t.Fatalf("number case = %q, %v; want 42", value, err)
	}

	texts := make(chan string, 1)
	texts <- "go"
	value, err = SelectFirst(context.Background(), nil, texts)
	if err != nil || value != "go" {
		t.Fatalf("text case = %q, %v; want go", value, err)
	}

	close(texts)
	if _, err := SelectFirst(context.Background(), nil, texts); err == nil {
		t.Fatal("expected an error for a closed channel")
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := SelectFirst(ctx, nil, nil); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled context error = %v", err)
	}
}
