package timedemo

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	SecondLimiter *rate.Limiter
	MinuteLimiter *rate.Limiter
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		SecondLimiter: rate.NewLimiter(rate.Limit(10), 10),
		MinuteLimiter: rate.NewLimiter(rate.Limit(1), 60),
	}
}

func (r *RateLimiter) Wait() {
	r1 := r.SecondLimiter.Reserve()
	r2 := r.MinuteLimiter.Reserve()

	delay1 := r1.Delay()
	delay2 := r2.Delay()

	delay := max(delay1, delay2)
	fmt.Println("Delay: ", delay)
	time.Sleep(delay)
}

func GenMultiRequest() {
	rL := NewRateLimiter()
	// ticker := time.NewTicker(1 * time.Second)
	//
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 	}
	// }
	//
	index := 0
	wg := &sync.WaitGroup{}
	for range 100 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rL.Wait()
			fmt.Println(i)
		}(index)
		index++
	}
	wg.Wait()
}
