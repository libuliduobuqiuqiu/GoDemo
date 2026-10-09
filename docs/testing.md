# 测试与演示

在仓库根目录运行默认验证：

```bash
go build ./...
go test ./...
```

默认测试应在没有私有配置、数据库、设备连接和固定目录的环境中结束。`test/` 中既有自动测试，也有用于观察行为的手动演示。需要外部条件的测试文件使用 `//go:build integration` 标记；默认命令不会执行它们。

可以只编译全部手动演示，检查代码和依赖是否仍然有效：

```bash
go test -tags=integration -run '^$' ./...
```

运行手动演示时，请指定包和测试名称。例如，在准备好 MySQL 配置与服务后：

```bash
go test -tags=integration -run '^TestGetUsers$' ./test/gostorage
```

不同演示的前提如下：

| 位置 | 运行前提 |
| --- | --- |
| `test/gostorage`、`test/gotool/ssh_test.go` | 本地 `configs/conf.json`，以及相应的数据库或 SSH 设备 |
| `test/gogin`、`test/concurrency/chat`、`test/golib/httpdemo_test.go` | 相应的本地服务或设备；部分用例会启动服务并持续运行 |
| `test/golib/filedemo_test.go`、`test/golib/osdemo_test.go`、`test/goothers`、`test/gotool` 中的文件演示 | 代码中指定的示例文件和目录；部分用例会写入或删除文件 |
| `test/concurrency/channel_test.go`、`test/concurrency/context_test.go`、`test/concurrency/sync_test.go` | 部分演示使用长时间等待，或需要外部信号结束 |
| `test/golib/rate_test.go`、`test/pkg/traceroute_test.go` | 分别需要较长时间等待、特定网络环境 |

`test/common_test.go` 中的 `TestCountTime` 还使用固定的本机日志目录；只有设置 `GO_DEMO_INTEGRATION=1` 才会运行。手动演示应逐个选择，不要一次运行整个带标签的测试集合。
