# Go 知识库

这里按知识主题收纳笔记，并把每篇笔记连接到仓库中的代码或测试。建议从基础主题开始，遇到具体工程问题时再查进阶目录。各页的“状态”区分已整理内容与尚待复核的旧笔记。

| 目录 | 收纳内容 | 入口 |
| --- | --- | --- |
| 基础 `base/` | 语言、标准库、并发、编译和运行时 | [基础主题](base/README.md) |
| 进阶 `advanced/` | 测试、反射、性能分析、RPC、数据库实践 | [进阶主题](advanced/README.md) |
| 框架 `framework/` | Web 框架与项目示例 | [框架主题](framework/README.md) |
| 面试 `interview/` | 按问题复习 Go、网络、存储和系统知识 | [面试主题](interview/README.md) |

## 建议的阅读顺序

1. [Go 基础提纲](base/go_manual.md)了解已有主题范围，再读[错误处理](base/go_error.md)与[并发控制实践](base/go_concurrency_practice.md)。
2. 用[测试笔记](advanced/go_testing.md)和[测试运行说明](testing.md)把概念与可运行样例对上。
3. 按工作中遇到的问题查[文件操作](base/go_file.md)、[GORM](advanced/go_gorm.md)、[RPC](advanced/go_rpc.md)等专题。
4. 面试目录用于复盘问题；遇到需要实践的知识，以专题笔记和代码为主。

## 代码从哪里找

| 想看什么 | 代码与验证位置 |
| --- | --- |
| 语言特性和标准库 | [`internal/godemo`](../internal/godemo/)、[`internal/gomanual`](../internal/gomanual/)、[`internal/golib`](../internal/golib/)；测试见 [`test/godemo`](../test/godemo/)、[`test/manual`](../test/manual/)、[`test/golib`](../test/golib/) |
| 并发与上下文 | [`internal/goconcurrency`](../internal/goconcurrency/)；演示见 [`test/concurrency`](../test/concurrency/) |
| 数据库与 RPC | [`internal/gostorage`](../internal/gostorage/)、[`internal/gorpc`](../internal/gorpc/)；部分演示需要外部环境 |
| Web 与框架 | [`internal/goweb`](../internal/goweb/)；运行前先看[测试运行说明](testing.md) |

新增或整理笔记时，按[收纳与写作约定](knowledge-guide.md)填写结论、示例、边界和验证方式。项目改动与阶段结果记录在[整理日志](maintenance-log.md)。
