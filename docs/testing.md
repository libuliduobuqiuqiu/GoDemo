# 测试与演示

在仓库根目录运行默认验证：

```bash
go build ./...
go test ./...
```

默认测试应在没有私有配置、数据库、设备连接和固定目录的环境中结束。`test/` 中既有自动测试，也有用于观察行为的手动演示。需要外部条件的测试文件使用 `//go:build integration` 标记；默认命令不会执行它们。

文件读写和复制测试位于 `test/golib/filedemo_test.go`，使用 `t.TempDir()`，已纳入默认测试。`Makefile` 的 `wsl`、`mac` 目标仅对当前构建命令设置交叉编译环境，不会修改全局 `go env`。

需要连接数据库、SSH 或代理设备时，先复制示例配置并指定路径：

```bash
cp configs/conf_demo.json configs/conf.json
# 填写 configs/conf.json 中所需的连接信息
export GO_DEMO_CONFIG="$PWD/configs/conf.json"
```

`configs/conf.json` 已被 Git 忽略。配置读取优先使用代码传入的非空路径，其次使用 `GO_DEMO_CONFIG`；两者都没有时返回错误。默认构建和测试无需执行上面的步骤。

可以只编译全部手动演示，检查代码和依赖是否仍然有效：

```bash
go test -tags=integration -run '^$' ./...
```

MySQL 部分按[无数据库静态检查](advanced/go_database.md)验收；默认测试会运行配置、SQL 参数、GORM 映射和 `DryRun` 检查。无需准备 MySQL 服务。真实数据库往返只保留为可选的手动演示，未在当前阶段运行。

若自行运行手动演示，请指定包和测试名称。例如，在准备好[专用 MySQL 测试数据库](advanced/go_database.md)及配置后：

```bash
go test -tags=integration -run '^TestStorageRoundTrip$' -v ./test/gostorage
```

不同演示的前提如下：

| 位置 | 运行前提 |
| --- | --- |
| `test/gostorage`、`test/gotool/ssh_test.go` | 手动运行时设置 `GO_DEMO_CONFIG` 并准备相应的数据库或 SSH 设备；数据库静态检查见[sqlx 与 GORM 静态检查](advanced/go_database.md) |
| `test/gogin`、`test/concurrency/chat`、`test/golib/httpdemo_test.go` | 相应的本地服务或设备；部分用例会启动服务并持续运行 |
| `test/golib/osdemo_test.go`、`test/goothers`、`test/gotool` 中的手动文件演示 | 按具体用例准备输入；部分用例会写入或删除文件。归档用例通过 `GO_DEMO_ARCHIVE_SOURCE_DIR`、`GO_DEMO_ARCHIVE_TGZ`、`GO_DEMO_ARCHIVE_ZIP` 指定输入，输出位于临时目录 |
| `test/concurrency/channel_test.go`、`test/concurrency/context_test.go`、`test/concurrency/sync_test.go` | 文件中仍有较长时间运行的观察性演示；快速、可断言的并发测试位于 `internal/goconcurrency/` |
| `test/golib/rate_test.go`、`test/pkg/traceroute_test.go` | 分别需要较长时间等待、特定网络环境 |

`test/common_test.go` 中的 `TestCountTime` 只有同时设置 `GO_DEMO_INTEGRATION=1` 和 `GO_DEMO_LOG_DIR` 后才会遍历指定日志目录。手动演示应逐个选择，不要一次运行整个带标签的测试集合。
