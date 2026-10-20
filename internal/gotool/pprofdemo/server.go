package pprofdemo

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"net/http"
	httppprof "net/http/pprof"
	"time"
)

// NewHandler returns an isolated mux with both the workload and profiling routes.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /debug/pprof/", httppprof.Index)
	mux.HandleFunc("GET /debug/pprof/cmdline", httppprof.Cmdline)
	mux.HandleFunc("GET /debug/pprof/profile", httppprof.Profile)
	mux.HandleFunc("GET /debug/pprof/symbol", httppprof.Symbol)
	mux.HandleFunc("GET /debug/pprof/trace", httppprof.Trace)
	mux.HandleFunc("GET /work/cpu", CPUHandler)
	return mux
}

type cpuResult struct {
	Iterations uint64        `json:"iterations"`
	Checksum   [32]byte      `json:"checksum"`
	Elapsed    time.Duration `json:"elapsed_ns"`
}

// CPUHandler runs a bounded CPU workload and can also be mounted in Gin.
func CPUHandler(w http.ResponseWriter, r *http.Request) {
	duration := 5 * time.Second
	if raw := r.URL.Query().Get("duration"); raw != "" {
		parsed, err := time.ParseDuration(raw)
		if err != nil || parsed < time.Millisecond || parsed > 2*time.Minute {
			http.Error(w, "duration must be between 1ms and 2m (example: 10s)", http.StatusBadRequest)
			return
		}
		duration = parsed
	}
	result := burnCPU(r.Context(), duration)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}

// burnCPU performs real computation throughout the requested wall-clock window.
// Results are returned to the caller so the computation cannot be discarded.
func burnCPU(ctx context.Context, duration time.Duration) cpuResult {
	start := time.Now()
	result := cpuResult{}
	for time.Since(start) < duration {
		if ctx.Err() != nil {
			break
		}
		for i := 0; i < 4096; i++ {
			result.Checksum = sha256.Sum256(result.Checksum[:])
			result.Iterations++
		}
	}
	result.Elapsed = time.Since(start)
	return result
}
