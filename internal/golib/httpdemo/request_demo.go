package httpdemo

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// RequestHtml 请求一个网页，解析获取网页中的节点元素
func RequestHtml(url string) {
	transport := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: &transport,
	}

	resp, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)

	htmlText, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	visitHtml(nil, htmlText)
}

func visitHtml(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visitHtml(stack, c)
	}
}
