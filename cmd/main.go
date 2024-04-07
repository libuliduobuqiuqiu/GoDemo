package main

import "godemo/internal/goconcurrency"

// "github.com/beego/beego/v2/server/web"
// _ "godemo/internal/goweb/gowebsockets"

func main() {
	// web.BConfig.CopyRequestBody = true
	// web.Run(":8090")
	// goconcurrency.PrintFib()
	goconcurrency.CountBalance()
}
