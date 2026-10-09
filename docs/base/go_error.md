# Go 错误处理

> 状态：已整理；验证环境：Go 1.25。关联代码：[错误链示例](../../internal/godemo/errors/errors_demo.go)、[堆栈示例](../../internal/godemo/errors/stack_demo.go)、[测试](../../test/godemo/errors_test.go)。

## 先记住这几条

- `error` 是普通返回值。产生错误的函数返回它，调用方决定重试、降级、记录或终止。
- 给错误补充操作上下文时，用 `fmt.Errorf("操作: %w", err)` 保留原始错误；`%v` 只保留文字。
- 调用方用 `errors.Is` 判断已知错误，用 `errors.As` 提取错误类型。不要依赖 `err.Error()` 的完整文本做分支。
- 自定义错误需要暴露原因时，实现 `Unwrap() error`。方法名和大小写必须准确。

## 保留错误链

```go
package main

import (
    "errors"
    "fmt"
    "os"
)

func readConfig(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("read config %q: %w", path, err)
    }
    return data, nil
}

func main() {
    _, err := readConfig("config.json")
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("配置不存在，可创建默认配置")
    }
}
```

每一层只补充对定位问题有用的上下文，例如文件名或操作名。错误在调用边界统一记录，避免每层都打印同一个错误。只有程序入口能够决定是否退出；库函数应将错误返回给调用方。

## 自定义错误

需要携带领域信息时可以定义类型。若同时包裹底层原因，就实现 `Unwrap`：

```go
type OperationError struct {
    Op    string
    Cause error
}

func (e *OperationError) Error() string { return e.Op + ": " + e.Cause.Error() }
func (e *OperationError) Unwrap() error { return e.Cause }
```

之后 `errors.Is` 和 `errors.As` 可以沿错误链查找。[项目中的 `MyError`](../../internal/godemo/errors/errors_demo.go) 展示了这一点；[对应测试](../../test/godemo/errors_test.go) 会检查链上的类型和原因。

## 堆栈与第三方库

需要记录错误发生位置时，在**创建错误时**捕获调用栈；在 `Error()` 中临时取栈，得到的是打印错误的位置。[`StackError`](../../internal/godemo/errors/stack_demo.go) 展示了创建时捕获。仓库也保留了 `github.com/pkg/errors` 的历史示例，用于理解已有项目；普通的包装、判断和类型提取可以直接使用标准库。

## 验证与待办

```bash
go test ./test/godemo -run '^TestHandleError$'
```

仓库中仍有部分配置和数据库演示使用 `log.Fatal` 处理库函数错误；后续整理这些模块时，应逐步改为返回错误，避免调用方无法处理。其他测试与手动演示的运行方式见[测试说明](../testing.md)。
