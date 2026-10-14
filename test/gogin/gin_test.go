//go:build integration

package test

import (
	"godemo/internal/goweb/gogin/app"
	"godemo/internal/goweb/gogin/proxy/client"
	"testing"
)

func TestGinStart(t *testing.T) {
	if err := app.NewRouter("../../web/templates/index.html").Run("127.0.0.1:8999"); err != nil {
		t.Fatal(err)
	}
}

func TestCommitHttpReq(t *testing.T) {
	tmpUrl := "http://127.0.0.1:8090/netac/base/proxy?proxy_pass=https://10.21.21.64:443/mgmt/tm/ltm/monitor/http/~Common~dfgdf"
	if err := client.CommitDeviceHttpReq(tmpUrl); err != nil {
		t.Fatal(err)
	}
}

// 测试交互式会话
func TestCommitTerminalReq(t *testing.T) {
	addr := "127.0.0.1:8090"
	client.CommitDeviceTerminalReq(addr)
}

// 测试代理Ssh请求
func TestCommitSshReq(t *testing.T) {
	if err := client.CommitDeviceSshReq(); err != nil {
		t.Log(err)
	}
}
