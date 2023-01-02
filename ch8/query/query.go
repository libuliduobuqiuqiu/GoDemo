package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println(mirroredQuery())
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("http://www.baidu.com") }()
	go func() { responses <- request("http://www.sina.com") }()
	go func() { responses <- request("http://www.qq.com") }()
	return <-responses
}

func request(hostname string) (response string) {
	resp, err := http.Get(hostname)

	if err != nil {
		log.Fatal("Get Error: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read Error: ", err)
	}
	response = string(body)
	return
}
