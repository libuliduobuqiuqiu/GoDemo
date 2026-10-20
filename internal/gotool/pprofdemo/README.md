# pprof CPU 分析示例

项目入口 `main.go` 使用 Cobra 管理命令，`internal/` 存放各类 Go 示例，`test/` 存放分类测试。
已有 `pprof_demo.go` 演示离线 CPU profile 和 trace；本示例补充 HTTP 在线采样和持续 CPU 负载。
旧的 `AnalysisHttpServer` 使用自己的 mux，没有注册 pprof 路由，仅空导入 `net/http/pprof` 并不能让该 mux 支持 pprof。

## 启动

在项目根目录执行（独立入口，不依赖数据库、Gin 或项目配置）：

```sh
go run ./cmd/pprof-demo
```

默认监听 `127.0.0.1:6060`，可用 `-addr 127.0.0.1:7070` 修改。
浏览器访问 <http://127.0.0.1:6060/debug/pprof/> 查看 profile 列表。

也已接入项目现有 Gin 服务：

```sh
go run . gin
```

此时以下命令中的端口全部改为 `8999`。例如在两个终端分别执行：

```sh
curl --fail-with-body -w '\nHTTP %{http_code}, 耗时 %{time_total}s\n' 'http://127.0.0.1:8999/work/cpu?duration=60s'
go tool pprof 'http://127.0.0.1:8999/debug/pprof/profile?seconds=10'
```

修改代码后需停止并重启服务。若请求立即返回，请检查 HTTP 状态和响应内容；
`404` 表示命中了未注册该接口的服务或旧进程，`400` 表示 duration 参数错误。
原 `/json-group/print-json` 仅执行少量计算，会很快返回，持续采样请使用 `/work/cpu`。
负载接口与 profile 必须请求同一个服务进程、同一个端口。

## 采样时产生 CPU 活动

终端 A 发起持续 60 秒的计算请求：

```sh
curl 'http://127.0.0.1:6060/work/cpu?duration=60s'
```

请求执行期间，在终端 B 采样 10 秒并打开分析终端：

```sh
go tool pprof 'http://127.0.0.1:6060/debug/pprof/profile?seconds=10'
```

在 `(pprof)` 提示符下执行：

```text
top
top -cum
list burnCPU
quit
```

`top` 按函数自身 CPU 时间排序；`top -cum` 包含被调用函数的时间。
预期能看到 SHA-256 相关函数，累计热点中能看到 `pprofdemo.burnCPU`。
具体名称和占比会随 CPU 架构、编译器和调度变化。

也可以保存采样后再分析：

```sh
curl -o /tmp/cpu.pprof 'http://127.0.0.1:6060/debug/pprof/profile?seconds=10'
go tool pprof -top /tmp/cpu.pprof
go tool pprof -http=127.0.0.1:8080 /tmp/cpu.pprof
```

图形调用图需要系统安装 Graphviz 的 `dot`。

CPU profile 只记录采样窗口内实际使用 CPU 的工作。必须让计算请求与采样重叠；
如果先等待计算请求完成，再采样空闲服务器，`top` 就可能为空。`sleep`、等待网络或等待锁不能替代 CPU 计算。
接口默认计算 5 秒，`duration` 支持 `1ms` 到 `2m`，请求取消后会停止计算。
响应包含计算次数、校验值和实际耗时；循环持续执行 SHA-256，不使用 sleep 制造耗时。
一个请求主要占用一个 CPU 核，需要更多负载可同时发起多个请求。

## 测试

```sh
go test ./internal/gotool/pprofdemo ./cmd/pprof-demo -v
go test ./internal/goweb/gogin/app -run TestCPUWorkRouteOverHTTP -v
go test -race ./internal/gotool/pprofdemo
```

测试覆盖参数边界、实际计算结果、取消请求，以及在计算接口运行时采样 2 秒，
调用真正的 `go tool pprof -top` 验证 `burnCPU` 出现在 profile 中，避免仅检查文件非空。
`go test -short ./internal/gotool/pprofdemo` 可跳过 CPU 采样测试。
如果精简工具链缺少 `go tool pprof`，可通过 `PPROF_BINARY=/path/to/pprof` 指定独立 pprof 可执行文件运行测试；
交互分析命令也相应使用该可执行文件替代 `go tool pprof`。

这里使用独立 mux 显式注册 pprof 路由，CPU profile 采集的是整个进程，不能同时启动多个 CPU 采样。
默认仅绑定本机；这是用于本地分析的计算接口和调试服务。
