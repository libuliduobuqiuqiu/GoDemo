package pprofdemo

import (
	"fmt"
	"godemo/internal/golib/httpdemo"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func AnalysisFibByPprof(path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	if err := pprof.StartCPUProfile(f); err != nil {
		return err
	}
	defer pprof.StopCPUProfile()

	var n = 10
	for i := 0; i <= 5; i++ {
		fmt.Printf("Fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
	return nil
}

func AnalysisFibByTrace(path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	if err := trace.Start(f); err != nil {
		return err
	}
	defer trace.Stop()

	var n = 10
	for i := 0; i <= 5; i++ {
		fmt.Printf("Fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
	return nil
}

func AnalysisHttpServer() {
	httpdemo.HandleHttpRequest()
}
