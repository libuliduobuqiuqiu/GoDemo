package gomutex

import "sync"

func UseMutexDemo() {
	var b int
	a := sync.Mutex{}
	a.Lock()
	b += 1
	a.Unlock()
}
