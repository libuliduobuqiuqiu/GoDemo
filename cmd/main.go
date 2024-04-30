package main

import "godemo/internal/gotool/pprofdemo"

// "github.com/beego/beego/v2/server/web"
// _ "godemo/internal/goweb/gowebsockets"
//
//

type Node struct {
	FileNodeIP string
	Status     string
}

func main() {
	// web.BConfig.CopyRequestBody = true
	// web.Run(":8090")
	// goconcurrency.PrintFib()

	// profdemo.AnalysisFibByPprof()
	// pprofdemo.AnalysisFibByTrace()
	pprofdemo.AnalysisHttpServer()
}
