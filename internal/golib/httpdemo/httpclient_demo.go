package httpdemo

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// GetRequest 通过context控制http请求
func GetRequest(ctx context.Context, url string) (rawResbody []byte, err error) {
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return ReadRespDataByBuffer(resp.Body)
}

func ReadRespDataByOnceTime(r io.Reader) (raws []byte, err error) {
	// 无缓冲直接读取所有
	raws, err = io.ReadAll(r)
	return
}

// ReadRespDataByBuffer 根据缓存分区读取响应内容
func ReadRespDataByBuffer(r io.Reader) (raws []byte, err error) {
	tmp := make([]byte, 1024)
	reader := bufio.NewReader(r)

	for {
		_, err := reader.Read(tmp)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		raws = append(raws, tmp...)
	}
	return
}

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
