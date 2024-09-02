package gomutex

import (
	"fmt"
	"sync"
	"time"
)

var balance int
var bLock sync.RWMutex

func Deposite(amount int, wg *sync.WaitGroup) {
	bLock.Lock()

	defer wg.Done()
	defer bLock.Unlock()

	balance += amount
	fmt.Println(balance)
	time.Sleep(2 * time.Second)
}

func Balance(wg *sync.WaitGroup) {
	bLock.RLock()
	defer wg.Done()
	defer bLock.RUnlock()
	fmt.Println("balance=", balance)
	time.Sleep(2 * time.Second)
}

func CountBalance() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Balance(wg)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Deposite(100, wg)
	}
	wg.Wait()
}
