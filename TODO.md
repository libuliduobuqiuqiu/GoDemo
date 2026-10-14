# Go 知识库待办

目标：让本仓库持续作为个人 Go 知识库使用。每个主题都应能从[知识库首页](docs/README.md)找到，说明结论、适用边界、对应代码和验证方式；需要外部服务的演示单独标明运行前提。完成一项后，更新相关主题索引、本文勾选状态和[整理记录](docs/maintenance-log.md)。文章格式见[收纳与写作约定](docs/knowledge-guide.md)。

## 当前进度（2026-10-14）

- 已完成：默认构建与测试基线、知识库导航、配置和固定路径整理、并发与测试专题，以及 gRPC Hello World 的本机通信验证。RPC chat 已明确标为未实现；“明确状态”这一待办已完成，聊天功能本身没有实现。
- 数据库部分按当前约定完成静态验收：检查配置、sqlx SQL 拼接、GORM 模型关联、查询和事务代码；修正批量插入与公司关联示例中的明确逻辑错误，并加入不连接数据库的测试。`TestStorageRoundTrip` 保留作可选手动演示，未执行，不将实际读写或回滚结果标为已验证。
- 下一步：整理 Gin 与代理，再逐篇复核旧笔记；无需为本阶段准备 MySQL 服务。
- 本阶段验证：`go build ./...`、`go test -timeout 20s ./...`、`go test -tags=integration -run '^$' ./...`、`go vet ./internal/gostorage/... ./pkg`、`go vet -tags=integration ./test/gostorage`。带标签的命令只编译手动演示。

## 优先级与执行顺序

排序依据：先消除会影响其他示例的配置和环境问题，再完成可独立验收的示例，最后复核旧笔记。同一优先级内按编号执行。数据库部分采用静态检查和无数据库测试作为本阶段验收标准。`P0` 优先处理，`P1` 随后处理，`P2` 持续维护。

| 顺序 | 优先级 | 事项 | 当前状态 |
| --- | --- | --- | --- |
| 1 | P0 | 配置读取与错误处理 | 已完成 |
| 2 | P0 | 固定路径演示与 Makefile | 已完成 |
| 3 | P1 | gRPC Hello World | 已完成并做过本机通信验证 |
| 4 | P1 | RPC chat 状态 | 已标明未实现 |
| 5 | P1 | sqlx、GORM 数据库示例 | 静态检查与无数据库测试已完成 |
| 6 | P1 | GORM 笔记 | 已按静态验收范围整理 |
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
- [x] 梳理 [`internal/gostorage`](internal/gostorage/) 和 [`test/gostorage`](test/gostorage/) 中的 sqlx、GORM 示例：修正批量插入占位符与参数、空条件列表、公司关联字段和空公司列表处理；用无数据库测试检查 SQL 生成及模型映射，用 `go vet` 和带标签编译检查其余演示。完成标准：默认测试不连接数据库，静态检查可重复执行。真实数据库读写不在本阶段验收范围；[往返测试](test/gostorage/roundtrip_test.go)仅保留为可选手动演示，见[数据库说明](docs/advanced/go_database.md)。
- [x] 复核 [GORM 笔记](docs/advanced/go_gorm.md) 的核心 API 与示例，把查询、事务和错误处理连接到仓库代码；写明版本、无数据库验证命令及未验证的运行行为，并同步[进阶索引](docs/advanced/README.md)。旧业务表、GORM Gen、实际事务回滚结果没有经过数据库验证。
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
