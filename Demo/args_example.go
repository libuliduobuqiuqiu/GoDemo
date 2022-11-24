package Demo

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func Args_main() {
	start := time.Now()

	ch := make(chan string)
	for _, url := range os.Args[1:] {

		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	seconds := time.Since(start).Seconds()
	fmt.Printf("%.2f elapsed\n", seconds)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	httpCheck := strings.HasPrefix(url, "http://")
	httpsCheck := strings.HasPrefix(url, "https://")

	if !httpCheck && !httpsCheck {
		url = "http://" + url
	}
	fmt.Println(url)

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("Get error: %s ", err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("Read Body error: %s", err)
		return
	}
	seconds := time.Since(start).Seconds()

	fmt.Printf("%s: %s", url, resp.Status)
	ch <- fmt.Sprintf("%.2f %7d %s", seconds, nbytes, url)
}
