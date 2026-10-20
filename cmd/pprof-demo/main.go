// pprof-demo runs a CPU workload and pprof endpoints on the same local server.
package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"godemo/internal/gotool/pprofdemo"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:6060", "HTTP listen address")
	flag.Parse()
	server := &http.Server{
		Addr: *addr, Handler: pprofdemo.NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("pprof: http://%s/debug/pprof/ ; CPU workload: /work/cpu?duration=30s", *addr)
	log.Fatal(server.ListenAndServe())
}
