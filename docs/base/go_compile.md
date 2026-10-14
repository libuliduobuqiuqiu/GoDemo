# Go 构建与逃逸分析

> 状态：已整理并验证；验证环境：Go 1.25.5、Linux/amd64。代码入口：[逃逸示例](../../internal/godemo/escape_demo.go)、[项目模块](../../go.mod)。

## 构建命令

`go build ./...` 编译仓库中可构建的包，`go test ./...` 额外运行默认测试；`go run . --help` 临时构建并运行根命令。编译器、汇编器和链接器共同完成构建，工具链可按包使用缓存；“每个源文件先生成一个独立的 `.s` 再链接”不是 `go build` 的固定流程。`GOROOT` 指向 Go 安装位置，模块项目的依赖版本由 `go.mod` 管理；常规构建不要求源代码位于 `GOPATH/src`。参见[Go 命令文档](https://pkg.go.dev/cmd/go)。

## 逃逸分析

编译器根据值的生命周期与引用关系决定存储方式。对局部变量取地址并不必然导致堆分配；返回指针也不能被当作所有调用点都必然堆分配的语言规则，因为内联等优化可能改变最终结果。是否逃逸是当前编译结果，不是 API 承诺。示例用 `//go:noinline` 保持诊断易读：

```go
func NonEscapingAddress() int {
    x := 1
    p := &x
    return *p
}

func ReturningAddress() *int {
    x := 1
    return &x
}
```

实际代码见[`escape_demo.go`](../../internal/godemo/escape_demo.go)。本机运行：

```bash
go build -gcflags=-m ./internal/godemo 2>&1 | rg 'escape_demo.go'
```

Go 1.25.5 对 `ReturningAddress` 的 `x` 报告 `moved to heap: x`，对 `NonEscapingAddress` 未报告该变量逃逸。编译器诊断文字及优化结果可能随版本、架构和上下文改变。减少不必要的分配有时能减轻 GC 负担，但是否值得修改代码应由基准测试和分析数据决定。参见[Go FAQ 的栈与堆说明](https://go.dev/doc/faq#stack_or_heap)及[GC 专题](go_gc.md)。

## 验证

```bash
go version
go build ./...
go test -timeout 20s ./test/godemo
```

构建和测试检查项目可运行；`-gcflags=-m` 只观察当前编译器对示例的分析，不等于运行时性能测量。
