package main

import (
	"fmt"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func Crawl(url string) []string {
	fmt.Println(url, len(tokens))
	tokens <- struct{}{}
	list, err := godemo.Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	var n int
	n++

	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- Crawl(link)
				}(link)
			}
		}
	}
}
