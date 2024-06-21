package gocontext

import (
	"context"
	"fmt"
)

func UseCancelCtx() {
	b := context.Background()
	cancelCtx, cancel := context.WithCancel(b)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
		default:
			fmt.Println("等待取消。。。")
		}
	}(cancelCtx)
	cancel()
}
