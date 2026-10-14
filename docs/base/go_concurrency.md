# Go 并发与运行时

> 状态：已整理并验证（公开语义与 Go 1.25 调度概览）；验证环境：Go 1.25.5，模块语言版本 1.25。可运行代码：[Channel](../../internal/goconcurrency/gochannel/)、[Context](../../internal/goconcurrency/gocontext/)、[同步原语](../../internal/goconcurrency/gosync/)。

## 并发、并行和调度

并发是管理多个可独立推进的任务，并行是同一时刻利用多个执行单元运行任务。Go 运行时把 goroutine 调度到操作系统线程上；常用的 G、M、P 概念分别表示 goroutine、线程和运行 Go 代码所需的处理器资源。`GOMAXPROCS` 限制同时执行 Go 代码的处理器数量，不限制 goroutine 总数，也不是所有操作系统线程的硬上限。

在 Go 1.25，未显式设置 `GOMAXPROCS` 时，默认值可能受逻辑 CPU 数、进程 CPU 亲和性和 Linux 容器 CPU 配额影响，运行时还可能更新默认值。旧笔记把默认值固定说成“CPU 核心数”已失效。检查当前值可用 `runtime.GOMAXPROCS(0)`；具体规则见[Go 1.25 发行说明](https://go.dev/doc/go1.25)和[`runtime.GOMAXPROCS` 文档](https://pkg.go.dev/runtime#GOMAXPROCS)。

调度器实现会随 Go 版本改变。系统调用、网络轮询、抢占和工作窃取都可能影响何时运行一个 goroutine；业务代码不能依赖固定运行顺序、固定抢占时间或“每个 goroutine 对应一个线程”。旧笔记中精确到某个函数、队列及 10ms 阈值的流程不作为本页结论。

## Channel 与 select

| 操作 | nil 通道 | 已关闭通道 |
| --- | --- | --- |
| 发送 | 永远等待 | panic |
| 接收 | 永远等待 | 先读完缓冲数据，之后得到零值且 `ok == false` |
| 关闭 | panic | 再次关闭会 panic |

无缓冲通道的发送需与接收配对；缓冲通道的发送在缓冲满时等待。关闭通常由发送方负责，并需保证只关闭一次。`select` 在多个已就绪 case 中随机选择一个；没有已就绪 case 时执行 `default`，若没有 `default` 就等待。`select{}` 永远等待。代码见[`SelectFirst`](../../internal/goconcurrency/gochannel/channel_select.go)及[对应测试](../../internal/goconcurrency/gochannel/channel_select_test.go)。

通道发送与相应接收、通道关闭与观察到关闭的接收之间有同步关系；这不意味着任意共享变量访问都安全。读写共享数据应通过通道所有权、互斥锁或原子操作建立同步关系。详见[Go 内存模型](https://go.dev/ref/mem)。

## Context、WaitGroup 与退出

`context.Context` 传递取消信号、截止时间和请求范围的数据。`WithCancel`、`WithTimeout`、`WithDeadline` 返回的取消函数应由持有者调用；工作 goroutine 需监听 `Done` 并退出，取消不是强制中断。代码与断言见[`WaitForCancellation`](../../internal/goconcurrency/gocontext/timerctx_demo.go)和[测试](../../internal/goconcurrency/gocontext/timerctx_demo_test.go)。

`sync.WaitGroup` 等待一组任务结束，不传递错误，也不负责取消。`sync.Mutex` 保护共享状态；`sync.Cond` 用于在锁保护下等待条件，醒来后仍要在循环中重查条件。详细可运行用法见[并发控制实践](go_concurrency_practice.md)。goroutine 若一直卡在发送、接收或其他等待点，调用方应提供取消或关闭路径，避免资源长期滞留。

## 验证与边界

```bash
go test -count=10 ./internal/goconcurrency/gochannel ./internal/goconcurrency/gocontext ./internal/goconcurrency/gosync
CGO_ENABLED=1 go test -race ./internal/goconcurrency/gochannel ./internal/goconcurrency/gocontext ./internal/goconcurrency/gosync
```

首条重复验证退出与结果；第二条检查执行到的路径上是否有数据竞争，需可用的 C 编译器。它们都不能证明所有调度交错正确。调度器内部数据结构和策略只按本机 Go 1.25 源码核对，不应当作跨版本 API。
