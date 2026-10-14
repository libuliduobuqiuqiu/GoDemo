# 并发控制：Channel、Context 与 sync.Cond

> 状态：已整理并验证；验证环境：Go 1.25。对应代码与测试分别在 [`gochannel`](../../internal/goconcurrency/gochannel/channel_select.go)、[`gocontext`](../../internal/goconcurrency/gocontext/timerctx_demo.go)、[`gosync`](../../internal/goconcurrency/gosync/cond_demo.go) 中。

这篇记录三个常见问题：等待多个事件、让任务在取消时退出、让多个等待者在条件满足后继续运行。运行时和并发概览见[并发与运行时](go_concurrency.md)。

## select：等待多个事件

`select` 会执行一个已就绪的 case；若多个 case 同时就绪，不能依赖固定顺序。没有就绪 case 和 `default` 时会等待。nil channel 的发送与接收永远不会就绪，因此可以用它禁用某个 case。关闭的 channel 可继续接收缓冲数据，读完后得到零值和 `ok == false`。

[`SelectFirst`](../../internal/goconcurrency/gochannel/channel_select.go) 同时等待数字、文本或 `ctx.Done()`。它区分了“收到零值”和“通道已关闭”，也能用 nil channel 只启用一个输入。验证见[`TestSelectFirst`](../../internal/goconcurrency/gochannel/channel_select_test.go)。

使用时不要让所有输入都保持 nil 且没有取消来源；那会一直等待。发送到已关闭的 channel 会 panic。

## Context：取消和超时

把 `context.Context` 传给可能阻塞的操作；收到 `<-ctx.Done()` 后读取 `ctx.Err()`，区分主动取消与期限已到。创建 `WithCancel`、`WithTimeout` 或 `WithDeadline` 后，在拥有该 context 的函数中调用返回的 `cancel`，释放相关资源。

[`WaitForCancellation`](../../internal/goconcurrency/gocontext/timerctx_demo.go) 展示如何等待取消原因；[`GenNum`](../../internal/goconcurrency/gocontext/timerctx_demo.go) 展示生产者在取消时关闭输出 channel。消费者提前停止读取时必须取消 context，否则生产者可能永久卡在发送操作。对应验证见[`timerctx_demo_test.go`](../../internal/goconcurrency/gocontext/timerctx_demo_test.go)。

不要靠固定 `time.Sleep` 等待 goroutine 退出；根据场景使用 channel、`sync.WaitGroup` 或其它明确的完成信号。

## sync.Cond：等待共享条件

`sync.Cond` 适合多个 goroutine 等待同一个受锁保护的条件。检查条件和调用 `Wait` 时必须持锁；`Wait` 在等待期间释放锁，唤醒后重新加锁。醒来仍要重新检查条件，因此使用 `for !condition { cond.Wait() }`，不能只用 `if`。

[`BroadcastToListeners`](../../internal/goconcurrency/gosync/cond_demo.go) 用就绪信号确保监听者进入等待阶段，然后更新条件并 `Broadcast`，最后等待所有监听者结束。它不依赖 Ctrl+C 或全局状态；对应验证见[`cond_demo_test.go`](../../internal/goconcurrency/gosync/cond_demo_test.go)。

## 验证

```bash
go test -count=10 ./internal/goconcurrency/gochannel ./internal/goconcurrency/gocontext ./internal/goconcurrency/gosync
CGO_ENABLED=1 go test -race ./internal/goconcurrency/gochannel ./internal/goconcurrency/gocontext ./internal/goconcurrency/gosync
```

`-count=10` 用来重复观察并发测试的稳定性，`-race` 检查执行过程中可检测的数据竞争。当前环境默认关闭 CGO，所以第二条命令显式启用它；运行机器还需有可用的 C 编译器。两个命令的用途不同，都不能代替对退出条件和资源释放的检查。
