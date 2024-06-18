package gocontext

import "context"

func UseCancelCtx() {
	b := context.Background()
	cancelCtx, cancel := context.WithCancel(b)

	deadCtx, cancel := context.WithDeadline(cancelCtx)

	cancel()
}
