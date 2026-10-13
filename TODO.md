# Go 知识库待办

目标：让本仓库持续作为个人 Go 知识库使用。每个主题都应能从[知识库首页](docs/README.md)找到，说明结论、适用边界、对应代码和验证方式；需要外部服务的演示单独标明运行前提。完成一项后，更新相关主题索引、本文勾选状态和[整理记录](docs/maintenance-log.md)。文章格式见[收纳与写作约定](docs/knowledge-guide.md)。

## 当前进度（2026-10-13）

- 已完成：默认构建与测试基线、知识库导航、配置和固定路径整理、并发与测试专题，以及 gRPC Hello World 的本机通信验证。RPC chat 已明确标为未实现；“明确状态”这一待办已完成，聊天功能本身没有实现。
- 进行中：已写好独立的 sqlx/GORM 数据库往返测试、专用数据库准备说明和 GORM 核心 API 笔记；`integration` 测试已编译，但尚未连接真实 MySQL 运行，因此数据库与 GORM 两项仍未验收。
- 下一步：先在专用 MySQL 数据库运行 `TestStorageRoundTrip` 并修正问题，再同步数据库文章、GORM 索引和本待办的状态；之后处理 Gin 与代理，最后复核旧笔记。
- 本阶段验证：Go 1.25.0 下 `go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...` 和 `git diff --check` 通过。带标签的命令只编译手动演示，不代表真实数据库或设备验证通过。

## 优先级与执行顺序

排序依据：先消除会影响其他示例的配置和环境问题，再完成可独立验收的示例，最后复核旧笔记。同一优先级内按编号执行；涉及外部数据库的步骤以准备好测试环境为前提。`P0` 优先处理，`P1` 随后处理，`P2` 持续维护。

| 顺序 | 优先级 | 事项 | 当前状态 |
| --- | --- | --- | --- |
| 1 | P0 | 配置读取与错误处理 | 已完成 |
| 2 | P0 | 固定路径演示与 Makefile | 已完成 |
| 3 | P1 | gRPC Hello World | 已完成并做过本机通信验证 |
| 4 | P1 | RPC chat 状态 | 已标明未实现 |
| 5 | P1 | sqlx、GORM 数据库示例 | 最小路径已写，待真实 MySQL 验证 |
| 6 | P1 | GORM 笔记 | 核心 API 已复核，待真实 MySQL 验证 |
| 7 | P1 | Gin 与代理示例 | 未整理 |
| 8 | P2 | 基础旧笔记 | 待逐篇复核 |
| 9 | P2 | 进阶旧笔记 | 待逐篇复核 |
| 10 | P2 | 框架与面试旧笔记 | 待逐篇复核 |

## P0：先消除环境依赖

- [x] 移除 [`pkg/conf_init.go`](pkg/conf_init.go) 对仓库绝对路径的依赖，改为可明确传入或通过环境变量指定配置；缺少配置时返回可处理的错误。修正配置字段标签，并为配置读取补充无需本机私有文件的测试。完成标准：新检出的仓库无需修改源码即可运行默认构建与测试；数据库演示需要的配置单独说明。
- [x] 整理 [`test/golib`](test/golib/) 等仍使用固定本机路径的演示；适合自动验证的改用 `t.TempDir()`，其余保留 `integration` 标签并写明前提。检查 [`Makefile`](Makefile) 中会修改全局 `go env` 的目标，改为仅对当前命令生效的配置。完成标准：默认验证不依赖个人环境，也不会修改使用者的全局 Go 设置。

## P1：完成可复现示例

- [x] 完成 [`internal/gorpc/helloworld`](internal/gorpc/helloworld/) 的最小服务端和客户端演示：客户端使用超时上下文、正确处理连接与 RPC 错误；[RPC 笔记](docs/advanced/go_rpc.md) 给出启动顺序和预期输出。完成标准：可在本机按文档复现，且默认测试无需启动服务。
- [x] 为 [`internal/gorpc/chat`](internal/gorpc/chat/) 明确范围：补成可运行示例，或在目录和 RPC 索引中标注尚未实现。完成标准：入口状态与文档一致，不让读者误以为占位代码可直接运行。
- [ ] 梳理 [`internal/gostorage`](internal/gostorage/) 和 [`test/gostorage`](test/gostorage/) 中的 sqlx、GORM 示例，明确最小数据库准备、启动和清理步骤；将真实数据库验证留在 `integration` 标签内。完成标准：默认 `go test ./...` 不连接数据库，按文档配置后能单独运行至少一个数据库示例。
  - [x] 增加独立的 [`TestStorageRoundTrip`](test/gostorage/roundtrip_test.go) 和[数据库准备、清理说明](docs/advanced/go_database.md)；默认测试不连接数据库，带标签的演示已通过编译检查。
  - [ ] 在专用 MySQL 8 数据库按文档实际运行往返测试，确认创建、读写、回滚和自动删表；记录结果并修正发现的问题。旧业务表演示不在这条最小验收路径内。
- [ ] 复核 [GORM 笔记](docs/advanced/go_gorm.md) 的 API 与示例，把查询、事务和错误处理连接到仓库代码。完成标准：文章写明所用版本、可运行命令及外部条件，并在[进阶索引](docs/advanced/README.md)更新状态。
  - [x] 已核对核心 API、修正模型字段标签，并写明版本、代码入口和运行条件。
  - [ ] 结合真实数据库运行结果复核文章中的查询、更新和回滚示例，再将状态改为“已整理并验证”；旧业务表关联和 GORM Gen 仍需另行验证。
- [ ] 整理 [`internal/goweb/gogin`](internal/goweb/gogin/) 的启动入口与 [Gin 笔记](docs/framework/go_gin.md)，逐个列出依赖和可访问路由。先修复代理路由缺少处理函数、请求失败直接退出进程等问题；对任意目标代理、跳过 TLS 验证、WebSocket 来源和 SSH 主机密钥校验给出安全边界或移除不安全默认值。完成标准：至少一个基础路由可按文档运行，代理与终端功能的状态清楚且无未说明的不安全默认行为。

## P2：复核旧笔记

- [ ] 按[基础索引](docs/base/README.md)逐篇核对语言总览、泛型、文件、并发运行时、编译与 GC；补齐 [GC 笔记](docs/base/go_gc.md) 中的写屏障等未完成项。完成标准：每篇文章标注当前状态、适用 Go 版本、代码入口和验证依据；已复核后才改为“已整理并验证”。
- [ ] 按[进阶索引](docs/advanced/README.md)复核反射与 pprof，并处理 [pprof 笔记](docs/advanced/go_pprof.md) 的遗留 TODO。完成标准同上。
- [ ] 复核[框架索引](docs/framework/README.md)中的 Beego 笔记，以及[面试索引](docs/interview/README.md)里的问答；为 Redis 问答补充可靠的实践入口或明确只有知识笔记。完成标准：可从问题跳到对应专题或示例，缺少示例时如实标注。

## 已完成的基础

- [x] 修复构建与默认测试基线，区分自动测试和需要外部环境的演示。
- [x] 建立知识库首页、四类主题索引和写作约定。
- [x] 整理并验证错误处理、并发控制实践和测试专题。

各阶段的改动、验证命令与局限记录在[整理记录](docs/maintenance-log.md)。
