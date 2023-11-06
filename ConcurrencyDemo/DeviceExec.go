package ConcurrencyDemo

import (
	"fmt"
	"sync"
	"time"
)

var deviceMap sync.Map

func ExecCommand(ip string, command string, wg *sync.WaitGroup) {
	defer wg.Done()
	deviceLock, _ := deviceMap.LoadOrStore(ip, &sync.Mutex{})

	if tmpLock, ok := deviceLock.(sync.Locker); ok {
		tmpLock.Lock()
		defer tmpLock.Unlock()
	}

	defer deviceMap.Delete(ip)
	fmt.Printf("Start Device: %s exec command: %s....\n", ip, command)
	time.Sleep(2 * time.Second)
	fmt.Printf("End Device: %s exec command successfully.\n", ip)
}

func DeviceExecCommands() {

	wg := &sync.WaitGroup{}
	deviceList := []string{"192.168.121.1", "192.168.122.2", "192.168.122.3"}
	commandList := []string{"ifconfig", "top", "free -m", "whoami"}

	for _, d := range deviceList {
		for _, c := range commandList {
			wg.Add(1)
			go ExecCommand(d, c, wg)
		}
	}
	wg.Wait()
}
