package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
)

func customDirector(req *http.Request) {
	req.URL.Path = "/person/info"
	req.URL.Host = "127.0.0.1:8999"
	req.URL.Scheme = "http"
	req.Header.Add("X-Custom-Header", "Modified Header")
}

func customModifyResponse(resp *http.Response) error {
	body := "Hello,World"
	resp.Body = io.NopCloser(strings.NewReader(body))
	resp.Header["Content-Length"] = []string{fmt.Sprintf("%d", len(body))}
	return nil
}

func main() {
	server := &httputil.ReverseProxy{Director: customDirector, ModifyResponse: customModifyResponse}
	http.Handle("/python/services", server)
	_ = http.ListenAndServe(":9000", nil)
}
