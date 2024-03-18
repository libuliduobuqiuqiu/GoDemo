package contextdemo

import (
	"context"
	"fmt"
	"time"
)

func StartContextTask() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	go startSubContextTask(ctx, 2000*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("Main task: ", ctx.Err())
	}
}

func startSubContextTask(ctx context.Context, timeout time.Duration) {

	subCtx, cancel := context.WithTimeout(ctx, timeout)

	defer cancel()

	go subTask(subCtx, 5*time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("From main task:", ctx.Err())
	case <-subCtx.Done():
		fmt.Println("sub task: ", subCtx.Err())
	}
}

func subTask(ctx context.Context, timeout time.Duration) {
	go helloWorld(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("From sub task: ", ctx.Err())
	case <-time.After(timeout):
		fmt.Println("sub sub task timeout.")
	}
}

func helloWorld(ctx context.Context) {
	defer fmt.Println("hello world close.")
	for {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("hello, world")
		case <-ctx.Done():
			return
		}
	}
}
