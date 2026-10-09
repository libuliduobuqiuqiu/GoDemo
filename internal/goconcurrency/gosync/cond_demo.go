package gosync

import (
	"fmt"
	"sort"
	"sync"
)

// BroadcastToListeners waits until every listener is ready, then wakes them all.
func BroadcastToListeners(count int) []int {
	if count <= 0 {
		return nil
	}
	cond := sync.NewCond(&sync.Mutex{})
	ready := make(chan struct{}, count)
	done := make(chan int, count)
	released := false

	for i := 0; i < count; i++ {
		go func(id int) {
			cond.L.Lock()
			ready <- struct{}{}
			for !released {
				cond.Wait()
			}
			cond.L.Unlock()
			done <- id
		}(i)
	}

	for i := 0; i < count; i++ {
		<-ready
	}
	cond.L.Lock()
	released = true
	cond.Broadcast()
	cond.L.Unlock()

	listeners := make([]int, 0, count)
	for i := 0; i < count; i++ {
		listeners = append(listeners, <-done)
	}
	sort.Ints(listeners)
	return listeners
}

func UseSyncCond() {
	fmt.Println(BroadcastToListeners(10))
}
