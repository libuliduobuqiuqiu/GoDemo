package pprofdemo

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func AnalysisFib() error {
	f, err := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var n = 10
	for i := 0; i < 5; i++ {
		fmt.Printf("Fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
	return nil
}
