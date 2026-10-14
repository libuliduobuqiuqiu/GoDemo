# 基础主题

这里收纳 Go 语言、标准库与运行时知识。基础旧笔记已按 Go 1.25 重审；各页说明代码入口、验证方式和仍未覆盖的边界。

| 主题 | 文章 | 状态 | 相关代码 |
| --- | --- | --- | --- |
| 语言基础总览 | [Go 基础总览](go_manual.md) | 已整理并验证（核心规则） | [`internal/godemo`](../../internal/godemo/)、[`internal/gomanual`](../../internal/gomanual/) |
| 泛型 | [泛型与类型集](go_generics.md) | 已整理并验证 | [`internal/gomanual/genericsdemo`](../../internal/gomanual/genericsdemo/) |
| 错误处理 | [Go 错误处理](go_error.md) | 已整理并验证 | [`internal/godemo/errors`](../../internal/godemo/errors/)、[`test/godemo/errors_test.go`](../../test/godemo/errors_test.go) |
| 文件与路径 | [文件读写与路径](go_file.md) | 已整理并验证（普通文件） | [`internal/golib/osdemo`](../../internal/golib/osdemo/)、[`test/golib/filecopy_test.go`](../../test/golib/filecopy_test.go) |
| 并发控制实践 | [Channel、Context 与 sync.Cond](go_concurrency_practice.md) | 已整理并验证 | [`internal/goconcurrency`](../../internal/goconcurrency/) |
| 并发与运行时 | [并发与运行时](go_concurrency.md) | 已整理并验证（公开语义与调度概览） | [`internal/goconcurrency`](../../internal/goconcurrency/) |
| 编译与逃逸 | [构建与逃逸分析](go_compile.md) | 已整理并验证 | [`internal/godemo/escape_demo.go`](../../internal/godemo/escape_demo.go) |
| 垃圾回收 | [GC 笔记](go_gc.md) | 已整理并验证（Go 1.25） | [`internal/godemo/gc_demo.go`](../../internal/godemo/gc_demo.go) |

[返回知识库首页](../README.md)
