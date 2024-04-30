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

func AnalysisFibByPprof() error {
	// f, err := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0666)
	f, err := os.Create("cpu.profile")
	if err != nil {
		return err
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var n = 10
	for i := 0; i <= 5; i++ {
		fmt.Printf("Fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
	return nil
}

func AnalysisFibByTrace() error {

	f, err := os.Create("cpu.trace")
	if err != nil {
		return err
	}

	defer f.Close()
	trace.Start(f)
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
