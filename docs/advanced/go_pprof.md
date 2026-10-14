# pprof 与执行跟踪

状态：已按 Go 1.25 复核运行方式、HTTP 注册条件和输出路径。代码入口：[pprof 示例](../../internal/gotool/pprofdemo/pprof_demo.go)。

## 测试命令采集

`go test` 自带 CPU、内存和 trace 输出参数；针对确实会执行的测试或基准测试使用。示例：

```bash
go test ./internal/gotool/pprofdemo -run '^TestPprofDefaultMuxRegistration$' -o /tmp/go-demo-pprofdemo.test -outputdir /tmp -cpuprofile go-demo-cpu.prof -memprofile go-demo-mem.prof -trace go-demo-test.trace
go tool pprof -top /tmp/go-demo-cpu.prof
go tool trace /tmp/go-demo-test.trace
```

`-o` 把测试二进制文件放在临时目录，`-outputdir` 把 profile 输出放在临时目录，避免污染仓库。该短测试可能几乎没有 CPU 样本，不能据此判断性能。分析真实瓶颈时选用具有代表性的基准或工作负载，先留意运行时长、采样开销和机器差异。内存 profile 描述采样结果，`heap` 与 `allocs` 的统计口径不同；阻塞和互斥 profile 需要另行启用采样。参考 [runtime/pprof](https://pkg.go.dev/runtime/pprof)。

## 程序主动采集

`AnalysisFibByPprof(path)` 和 `AnalysisFibByTrace(path)` 接收显式输出路径，检查 `os.Create`、`pprof.StartCPUProfile` 或 `trace.Start`、文件关闭错误，并在关闭文件前停止采集。CPU 与 trace 是进程级采集，同一进程中不能同时启动两个同类采集器；并行运行的测试需避免争用。手动 CPU 演示仍在 `integration` 标签下：

```bash
go test -tags=integration ./test/gotool -run '^TestAnalysisFibProfile$' -v
```

该测试写入临时目录；退出后测试框架会清理文件。若要保留 profile，可在自己的调用代码中传入持久路径，再用 `go tool pprof -top <路径>` 或 `go tool pprof -http=127.0.0.1:8090 <路径>` 查看。图形视图可能需要 Graphviz；`-top` 不需要。

## HTTP 端点为什么会出现

导入 `_ "net/http/pprof"` 会执行包初始化，向 `http.DefaultServeMux` 注册 `/debug/pprof/` 等处理器。只有 HTTP 服务使用默认 mux（例如 `http.ListenAndServe("127.0.0.1:6060", nil)`）时，导入后才可直接访问。若服务使用自定义 mux（Gin、Beego 或 `http.NewServeMux()`），需要显式注册相应处理器，或启动一个独立且受保护的诊断服务。仓库测试用 `httptest.ResponseRecorder` 验证默认 mux 返回 200、自定义空 mux 返回 404；没有启动网络监听。原笔记的遗留 TODO 至此已回答。参考 [net/http/pprof 官方文档](https://pkg.go.dev/net/http/pprof)。

在可信本机环境启动独立诊断服务后可请求：

```bash
go tool pprof http://127.0.0.1:6060/debug/pprof/heap
go tool pprof 'http://127.0.0.1:6060/debug/pprof/profile?seconds=5'
curl -o trace.out 'http://127.0.0.1:6060/debug/pprof/trace?seconds=5'
go tool trace trace.out
```

上述命令是使用说明，本轮未启动 HTTP 诊断服务，也未验证真实负载的 profile。端点可暴露命令行参数、堆栈及运行时信息，不要把它直接开放给不可信网络。Go 1.22 起这些端点要求 GET；CPU profile 默认采集 30 秒，trace 默认采集 1 秒，可用 `seconds` 调整。参考 [net/http/pprof](https://pkg.go.dev/net/http/pprof)及 [runtime/trace](https://pkg.go.dev/runtime/trace)。

## 本地验证

```bash
go test ./internal/gotool/pprofdemo
go test -tags=integration -run '^$' ./test/gotool
```

第一条验证默认 mux 注册和显式 trace 输出；第二条只编译手动 CPU 演示。trace 文件是否能由工具解析、CPU 采样是否具有统计意义，需对目标工作负载另行验证。
