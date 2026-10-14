# Go 垃圾回收

> 状态：已整理并验证（Go 1.25 的公开行为及写屏障概览）；验证环境：Go 1.25.5、Linux/amd64。代码入口：[有界分配示例](../../internal/godemo/gc_demo.go)、[结果断言](../../test/godemo/gc_test.go)。

## 核心机制

Go 当前的垃圾回收器是并发、非移动的追踪式标记清除回收器。它从根（如仍可达的栈和全局引用）扫描对象图，标记可达对象，之后回收未标记的堆内存；局部值若被放在栈上，无需由堆 GC 单独回收。三色模型中的白、灰、黑分别表示尚未确认可达、已发现待扫描、已扫描。颜色是理解标记过程的抽象，不是程序可读的对象属性。参见[Go GC 指南](https://go.dev/doc/gc-guide)和本机 Go 1.25 `runtime/mgc.go`。

标记主要与程序并发执行，但阶段切换和部分根扫描仍有短暂停顿。旧笔记把 GC 描述为必须在整个标记清除期间停止所有用户代码，已不符合当前实现。回收器只管理 Go 运行时纳入的内存；`GOMEMLIMIT` 也不是整个进程内存的硬上限。

## 写屏障、删除屏障与混合屏障

并发标记时，程序仍会更新指针；若只扫描一次对象图，可能漏掉仍可达的对象。写屏障是编译器和运行时对某些指针写入执行的附加标记动作。Go 1.25 的 `runtime/mbarrier.go` 说明混合屏障结合两种思路：

1. 删除屏障在旧引用被覆盖前标记其指向的对象，防止唯一可见引用从堆移到尚未扫描的栈而丢失。
2. 插入屏障在当前 goroutine 的栈仍为灰色时标记新引用指向的对象，防止该引用从栈写入已扫描的堆对象后被漏掉。

可把堆指针写入概括为 `shade(old); if stackIsGrey { shade(new) }; store(new)`。这是解释正确性的伪代码；真实实现含缓冲、汇编和编译器优化，不能逐行当作 Go API。写屏障主要用于可能位于堆中的指针更新，普通栈局部写入不必一律执行屏障。实现依据：[Go 1.25.5 的 `mbarrier.go`](https://github.com/golang/go/blob/go1.25.5/src/runtime/mbarrier.go)和[GC 指南](https://go.dev/doc/gc-guide)。

## 调优与观察

`GOGC` 或 `runtime/debug.SetGCPercent` 控制下一轮堆目标；Go 1.25 的目标计算考虑上次存活堆及扫描根的规模，默认值为 100。提高它通常减少 GC 频率但增加内存占用。`GOMEMLIMIT` 或 `debug.SetMemoryLimit` 为 Go 运行时管理的内存设置软限制；限制过低可能使 GC 频繁运行，并不能保证进程总内存不超过该值。优先减少不必要的分配，再用基准测试、pprof 和生产环境指标判断调优收益。详见[Go GC 指南的调优部分](https://go.dev/doc/gc-guide)和[性能分析专题](../advanced/go_pprof.md)。

仓库原来的分配函数把未使用的切片丢给空白标识符，编译器可能消除这些分配；现在[`UseAllocate`](../../internal/godemo/gc_demo.go)保留并读取一组有界缓冲区，测试断言校验和。它验证示例逻辑，不保证触发固定次数的 GC。

```bash
go test ./test/godemo -run '^TestUseAllocate$'
GODEBUG=gctrace=1 go test ./test/godemo -run '^TestUseAllocate$' -count=1
```

第二条可观察 GC 日志；输出受机器、缓存、运行时和其他分配影响，不应断言固定数字。更可靠的性能比较见[测试专题](../advanced/go_testing.md)的基准测试与[pprof 笔记](../advanced/go_pprof.md)。
