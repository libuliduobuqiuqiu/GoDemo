# 项目整理记录

本文记录仓库重新作为个人 Go 知识库使用时的整理过程。知识主题的入口见后续的知识库目录；测试的具体运行方式见[测试与演示](testing.md)。

## 第一阶段：建立可复现的基线（2026-10-09）

### 整理前发现的问题

- `go build ./...` 因 RPC chat 的两个空 `main` 包失败。
- `go test -run '^$' ./...` 在当前 Go 1.25 工作区中遇到格式化调用检查错误；存储测试在导入包时就读取私有数据库配置，无法完成编译检查。
- 测试混有纯本地用例、长时间等待的并发演示、数据库与设备连接、固定路径文件操作。直接运行全部测试可能等待外部信号或访问本机资源。
- 部分测试断言本身有误；`test/manual` 的 `TestMain` 没有传递 `m.Run()` 的退出码。

### 已完成的改动

1. 修正格式化调用、字符串数字转换和 RPC 日志格式；将未实现的 RPC chat 入口标为普通包，使所有包可以构建。
2. 将 sqlx 与 GORM 的数据库连接从包导入阶段移至实际使用时；Coze 配置改为启动聊天时读取。GORM 连接失败时返回错误，避免继续对空对象操作。
3. 为 22 个需要环境或交互的测试文件加 `integration` 构建标签；固定本机日志目录和外部信号用例另加显式环境开关。保留了这些演示的代码与单独运行能力。
4. 修正 GoConvey、错误处理和回文示例中的错误断言；修正 `TestMain` 的退出码；等待生产者和消费者 goroutine 完成。
5. 增加使用 `t.TempDir()` 的文件复制测试，并将目录扫描示例改为相对路径。
6. 更新 README 与[测试说明](testing.md)，明确默认验证命令与手动演示前提。

本次基线还纳入整理前工作区已有的 Go 1.25、`golang.org/x/time` 限流示例和 `test/common_test.go` 中的新增示例；限流压力演示归入手动集成测试。

### 验证结果

在 Go 1.25.5、无私有配置的当前环境中，以下命令均通过：

```bash
go build ./...
go test -timeout 20s ./...
go test -tags=integration -run '^$' ./...
go run . --help
git diff --check
```

第三条命令只编译带标签的演示，不执行外部服务、设备和固定路径操作。它们的运行结果仍需在对应环境中逐个验证。后续新增默认测试时，应继续保持不依赖本机目录、私有配置或外部服务。

## 第二阶段：建立知识库入口（2026-10-09）

- 增加[知识库首页](README.md)、四个目录索引和[收纳与写作约定](knowledge-guide.md)，使文章可以按基础、进阶、框架、面试查找，并区分已整理内容与原始笔记。
- 将 Go 语言问答移入 `interview/`，将泛型和文件操作移入 `base/`；原有文章内容保留，索引按新路径链接。
- 将[错误处理](base/go_error.md)从问题提纲整理为专题，修正对应示例的 `Unwrap` 方法并用测试验证错误链。
- 校正[并发笔记](base/go_concurrency.md)中 nil channel 的行为，以及[测试笔记](advanced/go_testing.md)中的函数命名和 `TestMain` 退出码。

验证：`go test ./test/godemo -run '^TestHandleError$'` 和 `go test -timeout 20s ./...` 通过；已检查 README 与 `docs/` 中 28 个 Markdown 文件的本地链接，未发现失效路径。

当时确定下一批先处理并发与测试笔记，再整理数据库、RPC 和 Web 的运行前提。原始笔记的“待复核”状态表示尚未确认其中所有结论，不能把目录索引当作内容正确性的证明。

## 第三阶段：并发与测试专题（2026-10-09）

- 增加[并发控制专题](base/go_concurrency_practice.md)，用 Channel、Context、`sync.Cond` 三个可结束的示例解释等待、取消和广播；旧并发笔记继续标注为待全面复核。
- 将原先会无限等待的 `UseChannelSelect`、依赖十秒等待的 `UseTimerCtx`、依赖 Ctrl+C 的 `UseSyncCond` 改为可控示例。给每个核心行为增加包内测试，验证输入、取消和全部监听者完成。
- 重写[测试笔记](advanced/go_testing.md)，用仓库实际代码解释默认测试、`TestMain`、基准测试、模糊测试与竞态检测。基准测试改用 `b.N`；字符串反转模糊测试现在检查双重反转性质，并明确有效 UTF-8 前提。

验证结果：默认 `go test -timeout 20s ./...`、带标签的编译检查、并发包 `-count=10`、启用 CGO 的 `-race -count=3` 均通过；`BenchmarkEqual` 的短时运行与 `FuzzReverse` 的 2 秒模糊运行也通过。下一批整理数据库、RPC 与 Web：逐个明确最小运行环境、默认测试和手动演示的边界。

## 第四阶段：汇总待办与整理顺序（2026-10-10）

- 新增[知识库待办](../TODO.md)，从知识库首页和项目 README 提供入口。把未完成工作按数据库与配置、RPC 与 Web、旧笔记复核排列，并给每项写明涉及位置和完成标准。
- 下一步先解决数据库配置的绝对路径及错误处理，再把数据库示例与 GORM 笔记连接为可复现的专题；随后处理 RPC 和 Web。旧笔记在逐条验证前继续保留“原始笔记”或“待复核”状态。
- 本阶段只整理任务与导航，尚未改动数据库、RPC、Web 的运行行为；前面三个阶段的构建和测试结果仍以各阶段记录为准。

## 第五阶段：消除配置和本机路径依赖（2026-10-12）

- 配置读取改为显式路径优先、`GO_DEMO_CONFIG` 其次；缺少文件或 JSON 错误会返回带路径的错误。数据库、SSH 和代理示例的调用点同步处理错误，配置字段标签已修正。配置读取和 DSN 生成使用临时文件测试。
- 文件操作演示改为接收路径，`test/golib/filedemo_test.go` 使用 `t.TempDir()` 进入默认测试。手动归档演示从环境变量读取输入、写入临时目录；日志目录演示由 `GO_DEMO_LOG_DIR` 指定。
- `Makefile` 的交叉编译目标只为当前 `go build` 命令设置环境变量，不再写入全局 `go env`。配置和手动演示的准备步骤见[测试与演示](testing.md)。
- 验证：Go 1.25.5 下 `go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...`、`git diff --check` 均通过；`make -n wsl` 和 `make -n mac` 确认不再执行 `go env -w`。数据库和 SSH 的真实连接仍需在对应环境中验证。

## 第六阶段：完成 RPC 基础示例（2026-10-13）

- Hello World 客户端现在用 5 秒上下文约束连接与调用，区分连接和 RPC 错误，并在结束时关闭连接；服务端只监听本机地址。
- [RPC 笔记](advanced/go_rpc.md)写明启动顺序、预期输出和本机无 TLS 的适用边界。[chat 目录](../internal/gorpc/chat/README.md)标明协议代码已生成、客户端和服务端尚未实现；进阶索引与待办状态同步更新。
- 验证：`go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...` 均通过。未启动服务端时客户端约 5 秒返回超时；本机启动服务端后，客户端输出 `hello, world zhangsan`，服务端记录收到 `zhangsan` 请求。数据库和 Web 示例仍按后续待办处理。

## 第七阶段：数据库最小路径与待办盘点（2026-10-13）

- 增加独立的 [`TestStorageRoundTrip`](../test/gostorage/roundtrip_test.go)：sqlx 插入记录，GORM 读取、更新并在事务错误时回滚，最后由 sqlx 核对结果。测试只使用运行时创建的表，并在正常退出时清理；它保留在 `integration` 标签下。
- 增加[数据库准备与清理说明](advanced/go_database.md)，重写[GORM 笔记](advanced/go_gorm.md)的连接、模型、查询和事务部分，修正示例模型中的 GORM 字段标签。进阶索引与[待办](../TODO.md)同步区分“已写好最小路径”和“待真实数据库验证”。
- 在 Go 1.25.0 下，`go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...` 与 `git diff --check` 通过。当前环境没有运行 `TestStorageRoundTrip` 所需的专用 MySQL 数据库；带标签命令仅编译测试。实际连接、建表、读写、回滚和清理仍需按文档验收。
- 后续按待办顺序处理 Gin 基础路由及代理行为，再逐篇复核旧笔记。已有 sqlx、GORM 演示中依赖预置业务表的部分仍属单独验证范围。
