package golib

import (
	"context"
	"godemo/internal/golib/httpdemo"
	"testing"
)

func TestHttpClientGet(t *testing.T) {
	url := "http://127.0.0.1:8989/person"
	data, err := httpdemo.GetRequest(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}

func TestHttpServer(t *testing.T) {
	httpdemo.HandleHttpRequest()
}
