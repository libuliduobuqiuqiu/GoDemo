# Go 泛型与类型集

> 状态：已整理并验证；验证环境：Go 1.25.5，模块语言版本 1.25。代码入口：[`Number` 与 `SumInt`](../../internal/gomanual/genericsdemo/generic_demo.go)、[映射求和](../../internal/gomanual/genericsdemo/generic_demo2.go)、[断言测试](../../test/manual/generics_test.go)。

## 何时使用

当相同操作能安全地用于多个类型，并希望调用方保持具体类型时，使用类型形参。泛型函数在名称后写形参及约束，调用时编译器可推断类型实参。下面的仓库示例只接受底层类型为几种有符号整数的类型：

```go
type Number interface {
    ~int | ~int8 | ~int32 | ~int64
}

func SumInt[T Number](a, b T) T { return a + b }
```

`~int` 包含底层类型为 `int` 的自定义类型；没有 `~` 的 `int` 只表示精确的 `int` 类型。测试中的 `type Score int` 能传给 `SumInt`，验证了这一差异。`comparable` 适合映射键等需要 `==` 的类型；对类型参数能执行的操作必须对约束类型集中的所有类型成立。详见[语言规范的类型集](https://go.dev/ref/spec#General_interfaces)。

## 并集、交集与空类型集

`A | B` 表示类型项的并集；在同一接口内分行嵌入约束取交集。例如仓库中的 `Integer` 联合有符号与无符号整数，`MyNumber` 取 `Integer` 与 `SignedInt` 的交集；`MyInteger` 同时要求有符号和无符号整数，因此没有可用的普通类型实参。交集为空的约束本身可以声明，但实例化需要满足约束的类型。

包含类型项的非基本接口只能用作类型约束，不能把 `Number` 当作普通变量类型。接口的方法要求也可与类型项取交集；旧笔记关于“带方法的接口无法参与类型集”的笼统说法不准确。

## Go 1.25 的边界

泛型类型可以有方法，接收器声明该类型的形参；本项目使用的 Go 1.25 不支持给方法额外声明一套独立类型形参，需改为泛型函数。不能用类型形参直接作为新定义类型的底层类型。普通函数中的类型断言作用于接口值，不能直接对类型参数变量做断言；必要时先转为 `any`，但应优先让约束表达所需操作。这些限制按本仓库 Go 版本说明，不代表未来版本永久不变。

映射求和示例已按真实约束命名为 `SumIntegers`；它不包含浮点类型。运行：

```bash
go test ./test/manual -run '^TestGenericNumberConstraints$'
go build ./internal/gomanual/genericsdemo
```

这些命令检查可编译的类型集用法和数值结果；非法实例化的编译错误由[语言规范](https://go.dev/ref/spec)说明，未作为仓库默认测试加入不能编译的代码。
