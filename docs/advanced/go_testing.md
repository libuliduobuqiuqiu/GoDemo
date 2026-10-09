# Go 测试：从示例到自动验证

> 状态：已整理并验证；验证环境：Go 1.25。本文对应仓库中的 [测试示例](../../test/manual/) 与 [测试运行说明](../testing.md)。

## 四种入口

| 类型 | 函数名与用途 | 仓库示例 |
| --- | --- | --- |
| 单元测试 | `TestXxx(t *testing.T)`，验证行为与边界 | [`example_test.go`](../../test/manual/example_test.go) |
| 示例测试 | `ExampleXxx()`，用 `// Output:` 核对输出并生成文档 | [`example_test.go`](../../test/manual/example_test.go) |
| 基准测试 | `BenchmarkXxx(b *testing.B)`，使用 `b.N` 驱动测量次数 | [`BenchmarkEqual`](../../test/manual/example_test.go) |
| 模糊测试 | `FuzzXxx(f *testing.F)`，用种子和生成输入检查性质 | [`FuzzReverse`](../../test/manual/fuzz_reverse_test.go) |

测试文件以 `_test.go` 结尾。可以与被测代码同包，以便访问包内实现；也可以使用 `包名_test`，只通过公开接口测试。本仓库的部分早期演示集中在 `test/`，新增测试优先放在相关代码包中。

## 默认测试与手动演示

```bash
go test ./...
go test ./test/godemo -run '^TestHandleError$'
go test -tags=integration -run '^$' ./...
```

第一条运行默认测试。第二条按名称运行一个用例。第三条编译带 `integration` 标签的手动演示，但不运行匹配不到的测试。即使使用 `-run '^$'`，包的初始化和 `TestMain` 仍会执行，所以不要在包导入时连接数据库或读取私有配置。需要设备、数据库或固定文件的演示应单独选择，见[测试运行说明](../testing.md)。

`-short` 只是向测试传递“缩短运行”的信号；只有测试代码主动检查 `testing.Short()` 时才会改变行为，不能替代 `integration` 标签。

## 写出能说明结果的测试

- 验证输出或状态，而不是只打印日志。错误应使用 `t.Fatal`、`t.Errorf` 等明确报告。
- 文件操作使用 `t.TempDir()`；需要在测试结束时关闭资源，使用 `t.Cleanup()` 或合适的 `defer`。
- 有 goroutine 的测试要有可观察的完成条件。靠固定 `time.Sleep` 等待容易慢，也可能漏掉尚未退出的 goroutine。
- 表格测试适合覆盖多个输入；`t.Run` 给每个案例独立的名称。

例如 [`filecopy_test.go`](../../test/golib/filecopy_test.go) 在临时目录写入源文件，并比较复制后的内容；[`cond_demo_test.go`](../../internal/goconcurrency/gosync/cond_demo_test.go) 验证所有等待者都被唤醒。

## TestMain

只有整个测试包确实需要共享准备工作时才使用 `TestMain`。它必须传递 `m.Run()` 的退出码：

```go
func TestMain(m *testing.M) {
    code := m.Run()
    os.Exit(code)
}
```

仓库中的[实际示例](../../test/manual/example_test.go)会在运行前后打印提示，并正确传递退出码。普通测试的资源准备优先放在测试函数内；不要让 `TestMain` 使包在新机器上无法运行。

## 基准测试

`b.N` 由测试框架决定；不要写死迭代次数。先排除普通测试，再用较短测量时间验证基准入口：

```bash
go test ./test/manual -run '^$' -bench '^BenchmarkEqual$' -benchtime=100ms -benchmem
```

输出中的 `ns/op` 表示每次操作耗时，`B/op` 和 `allocs/op` 分别表示每次操作的分配字节数与分配次数。结果受机器、Go 版本和负载影响；跨环境比较时要记录这些条件。

## 模糊测试

模糊测试应检查一个可以对大量输入成立的性质。例如有效 UTF-8 文本反转两次应得到原文，[`FuzzReverse`](../../test/manual/fuzz_reverse_test.go)先加入种子，再检查这个性质：

```bash
go test ./test/manual -run '^$' -fuzz '^FuzzReverse$' -fuzztime=3s
```

不带 `-fuzz` 的默认 `go test` 仍会运行模糊测试的种子语料。随机生成的输入可能超出函数的有效范围；应在测试中明确前提，例如该示例只处理有效 UTF-8 文本。

## 并发测试

```bash
CGO_ENABLED=1 go test -race ./internal/goconcurrency/gochannel ./internal/goconcurrency/gocontext ./internal/goconcurrency/gosync
```

`-race` 需要 CGO 和可用的 C 编译器；当前环境默认关闭 CGO，因此命令显式启用。它可发现执行路径上的数据竞争，但一次通过不能证明所有调度顺序都正确。并发测试仍要验证退出、关闭和取消条件；并发样例见[专题笔记](../base/go_concurrency_practice.md)。
