package gosync

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func UseAtomic() {
	var count atomic.Int64
	var wg = &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go addValue(&count, wg)
	}

	wg.Wait()
	fmt.Println(count.Load())
}

func ConcurrencyAdd() {
	var count int64
	var wg = sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			count += 1
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func addValue(c *atomic.Int64, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	c.Add(1)
}
