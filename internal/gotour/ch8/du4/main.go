package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var sema = make(chan struct{}, 20)
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func walkDir(dirName string, n *sync.WaitGroup, fileSizes chan int64) {
	defer n.Done()

	if cancelled() {
		return
	}

	for _, entire := range dirents(dirName) {
		if entire.IsDir() {
			n.Add(1)
			path := filepath.Join(dirName, entire.Name())
			go walkDir(path, n, fileSizes)
		} else {
			fileSizes <- entire.Size()
		}
	}

}

func dirents(dirName string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}

	entries, err := ioutil.ReadDir(dirName)

	defer func() {
		<-sema
	}()

	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()

	fileSizes := make(chan int64)
	n := sync.WaitGroup{}
	roots := flag.Args()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	go func() {
		for _, root := range roots {
			n.Add(1)
			walkDir(root, &n, fileSizes)
		}
	}()

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range fileSizes {

			}
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}

			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
