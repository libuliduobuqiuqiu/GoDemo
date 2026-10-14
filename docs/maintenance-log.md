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

## 第八阶段：MySQL 静态验收（2026-10-14）

- 根据当前验收范围，不连接真实 MySQL。已检查配置/DSN、sqlx 查询与批量插入、GORM 字段和关联、查询及事务示例；`TestStorageRoundTrip` 保留为可选手动演示，实际数据库行为未验证。
- 修正旧 sqlx 批量插入的占位符数量和参数展开方式，移除返回切片的无效 `driver.Valuer` 实现；空名称列表直接返回，`sqlx.In` 错误不再被忽略。修正公司关联示例中混用 `Code` 与 `ID`、空公司列表越界及查询错误未处理的问题。
- 增加无需数据库的批量插入参数测试、GORM schema 关联测试、`DryRun` 查询和更新 SQL 测试。验证命令：`go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...`、`go vet ./internal/gostorage/... ./pkg`、`go vet -tags=integration ./test/gostorage`、`git diff --check`。
- [数据库说明](advanced/go_database.md)、[GORM 笔记](advanced/go_gorm.md)、进阶索引与待办已改为“静态验收完成，真实数据库运行未经验证”。后续工作转到 Gin 与代理示例；不再把连接 MySQL 作为当前阶段的待办。

## 第九阶段：Gin 基础路由与代理（2026-10-14）

- 增加 `go run . gin` 的本机入口，修正模板相对路径，提供 `/index` 和 `/json-group/print-json` 的可复现步骤。代理入口增加 `/healthz`；原本没有处理函数的代理路由现在会转发到启动时 `--proxy-target` 明确允许的 HTTP(S) 源站。
- 代理不再因请求错误调用 `log.Fatal`，格式错误、未允许的源站和上游故障分别给出 HTTP 错误；恢复标准 TLS 证书校验。默认监听地址改为 `127.0.0.1`。SSH/终端路由默认不注册；显式启用时要求 `GO_DEMO_SSH_KNOWN_HOSTS`，SSH 连接验证主机密钥，WebSocket 保留同源检查。SSH 响应错误分支不再继续写成功响应。
- 增加无需外部设备的 Gin 路由、代理允许列表、真实本机 HTTP 转发、自签名 TLS 拒绝、SSH 启动前提及跨源 WebSocket 拒绝测试。同步更新[框架索引](framework/README.md)、[Gin 笔记](framework/go_gin.md)、[测试说明](testing.md)和[待办](../TODO.md)。
- 验证：Go 1.25.5 下，`go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...`、`go vet ./internal/goweb/gogin/... ./cmd` 通过。代理测试需要本机临时端口；受限沙箱禁止监听时，测试在允许 loopback 的环境中运行。静态检查发现旧优雅退出示例使用无缓冲信号通道，已改为缓冲通道并停止信号通知。
- 边界：未连接真实 SSH 设备；旧 Gin 表单、上传、外部 HTTP 处理器和内部源码笔记仍待逐项复核。当前命令无客户端身份认证，文档限定本机与可信隔离环境。

## 第十阶段：基础旧笔记复核（2026-10-14）

- 按[基础索引](base/README.md)重写语言总览、泛型、文件、并发运行时、编译与逃逸及 GC 六篇。每篇标注 Go 1.25 版本、状态、代码入口、验证命令与未覆盖范围。旧版循环变量闭包、固定 CPU 核心数的 `GOMAXPROCS` 默认值和“返回局部变量地址必然堆分配”等绝对化结论已校正；GC 的写屏障、回收流程与 `GOGC`/`GOMEMLIMIT` 边界已补齐。版本相关规则核对了 Go 语言规范、Go 1.22/1.25 发行说明、GC 指南和本机 Go 1.25.5 运行时源码。
- 文件读取示例现在只转换本次实际读取的字节，正确处理空文件并传播读取错误；同包测试覆盖短文件和空文件，复制测试增加目标已有内容时的追加断言。泛型示例把仅支持整数的函数改为 `SumIntegers`，修正约束名称拼写并为结果增加断言。原来的分配示例可能被编译器消除，现改为有界分配、读取并校验结果。增加独立的逃逸示例，避免把编译器优化误写成语言保证。
- 验证：Go 1.25.5 下，`go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...`、`go vet ./internal/godemo ./internal/golib/osdemo ./internal/gomanual/genericsdemo ./test/manual ./test/golib`、并发三个包 `-count=10` 均通过。`go build -gcflags=-m ./internal/godemo` 报告逃逸示例的返回地址变量移到堆、GC 示例的缓冲块逃逸到堆；这只是本机编译结果。Markdown 本地链接检查未发现失效路径。
- 边界：特殊文件、符号链接、文件系统故障、跨平台编译诊断及真实工作负载下的 GC 性能没有验收。调度器内部队列和屏障实现只按 Go 1.25 源码解释，不作为语言 API 承诺。下一步按待办复核反射与 pprof。

## 第十一阶段：反射与 pprof 旧笔记复核（2026-10-14）

- 重写[反射笔记](advanced/go_reflect.md)，核对 `TypeOf`、`ValueOf`、`CanSet`、`CanInterface`、方法调用与 `DeepEqual` 的适用条件；原笔记中跨类型比较的 TODO 已给出结论。修正 `handleStruct` 成功后仍返回错误的问题，并检查无效输入与字段实际修改。
- 重写[pprof 笔记](advanced/go_pprof.md)，解释 `net/http/pprof` 导入时只向默认 mux 注册处理器，回答原笔记遗留 TODO；默认 mux 与空自定义 mux 的响应差异已有测试。CPU 和 trace 演示改为显式输出路径，检查采集启动及文件关闭错误；手动 CPU 测试写临时目录并确认文件非空。进阶索引、[测试说明](testing.md)和[待办](../TODO.md)同步更新。
- 验证：Go 1.25.5 下 `go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...`、`go vet ./internal/gomanual/reflectdemo ./internal/gotool/pprofdemo ./test/manual ./test/gotool`、`go test -tags=integration ./test/gotool -run '^TestAnalysisFibProfile$' -v -timeout 20s` 和 `git diff --check` 通过；文档中的 `go test` profile 参数组合也已运行通过。
- 边界：HTTP pprof 端点只验证了处理器注册，没有启动诊断服务；CPU 手动示例确认文件写入，但未对真实业务负载做性能解释。短测试生成的 CPU 样本可能不足；下一步复核 Beego 与面试问答。

## 第十二阶段：Beego 与面试问答复核（2026-10-14）

- 核对 Beego v2.1.6 依赖和聊天室路由、客户端、广播代码，重写[Beego 笔记](framework/go_beego.md)。原笔记中的“双向链表”不是仓库实现；聊天室没有调用 `web.Run` 的入口，故只标为静态复核。重名检查、连接退出及广播错误处理列为独立实现待办。
- 重写 Go、MySQL、Redis、网络、系统五篇[面试问答](interview/README.md)。Go `hmap`/溢出桶描述不再代表 Go 1.25 的 map 实现；修正 nil 切片、map 比较和接口方法集等边界。MySQL 的索引与隔离级别说明按 MySQL 8.4 文档限定；Redis 的 RDB/AOF、复制和淘汰策略按官方文档校正，并明确仓库没有 Redis 实践代码。网络与系统问答接入本机 RPC、Gin、编译、GC 和 pprof 专题；未做抓包或硬件测量。
- 同步更新[框架索引](framework/README.md)、[面试索引](interview/README.md)、[测试说明](testing.md)和[待办](../TODO.md)。验证：Go 1.25.5 下 `go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...`、`go vet ./internal/goweb/gowebsockets/chat ./test/interview`、Markdown 本地链接检查及 `git diff --check` 均通过。
- 边界：本阶段是笔记复核与静态编译；没有运行 Beego 聊天室、Redis、真实 MySQL 查询计划或事务，也没有网络抓包。下一步按待办处理 Gin 旧处理器；Beego 聊天室实现缺口保留为后续事项。

## 第十三阶段：Gin 旧处理器与内部调用路径（2026-10-14）

- 修正已注册 JSON 处理器在绑定失败后仍写成功响应的问题；隔离的登录表单修正 `password` 字段标签，并为缺字段、认证失败和成功分别返回 400、401、200。旧上传示例改为只读取或返回文件元信息，不使用客户端文件名写入当前目录；使用 `http.MaxBytesReader` 限制 8 MiB 请求体。代理的 `POST /netac/base/upload` 只解析 multipart，不记录字段内容，错误格式与超限分别返回 400、413。
- 隔离的外部 HTTP 图片处理器改为接受调用方配置的 URL 与客户端，默认客户端有 5 秒超时；上游非 200 返回 503，成功和失败均关闭响应体。测试使用内存 HTTP 传输替身，未访问公网。原本记录表单内容及管理员用户名的打印已移除。
- [Gin 笔记](framework/go_gin.md)增加已注册与隔离处理器的状态、Gin v1.9.1 的路由注册、`Engine.ServeHTTP`、中间件与绑定调用路径，并指出 `MaxMultipartMemory` 不是请求体上限。同步更新[框架索引](framework/README.md)、[测试说明](testing.md)和[待办](../TODO.md)。没有为设备验收或已知边界新增待办项。
- 验证：Go 1.25.5 下 `go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...`、`go vet ./internal/goweb/gogin/... ./cmd`、Markdown 本地链接检查及 `git diff --check` 通过。完整测试中的 Gin 代理用例需要 loopback 端口，在允许本机监听的环境执行；旧处理器测试不需要外部网络。
- 边界：`formGroupInit`、`othersGroupInit` 和独立上传服务没有接入主入口；固定演示账号不应用于真实认证。SSH/终端仍需专用设备、凭据和人工交互，本轮没有执行真实会话。当前清单余项是 Beego 聊天室运行示例。
