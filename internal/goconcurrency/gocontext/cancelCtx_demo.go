package gocontext

import (
	"context"
	"fmt"
	"time"
)

func UseCancelCtx() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		case <-time.After(10 * time.Second):
			fmt.Println("TimeOut")
			return
		}
	}

}

func secondCancelCtx(ctx context.Context) {
	time.Sleep(5 * time.Second)
	v := ctx.Value("name")
	fmt.Println(v)
}
