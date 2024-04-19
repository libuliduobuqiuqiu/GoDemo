package gocontext

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func execCmd(ctx context.Context, deviceName string) {
	connectDevice(ctx, deviceName)
}

func connectDevice(ctx context.Context, deviceName string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(deviceName + " had exit.")
			return
		default:
			fmt.Println(deviceName + " connecting.")
			time.Sleep(2 * time.Second)
		}
	}
}

func UseContext() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 10; i++ {
		go execCmd(ctx, fmt.Sprintf("device%d", i))
	}

	time.Sleep(4 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("Done.")

}
