# 反射

状态：已按仓库 Go 1.25 示例复核。入口：[反射示例](../../internal/gomanual/reflectdemo/)；可验证用例在同包的 `reflect_struct_test.go` 和 [手动演示测试](../../test/manual/reflect_test.go)。

## 类型和值

`reflect.TypeOf(x)` 返回接口中具体值的动态类型；`x == nil` 时返回 `nil`。`Kind()` 是底层类别（如 `Struct`、`Pointer`），而 `Type.String()` 可保留命名类型信息。`Elem()` 只适用于数组、通道、映射、指针、切片等支持元素类型的 `Type`，不能对任意类型调用。见 [TypeOf 示例](../../internal/gomanual/reflectdemo/reflect_typeof.go)。

`reflect.ValueOf(x)` 取得值；需要修改结构体字段时，通常传入非 nil 指针，再调用 `Elem()`。使用字段前检查 `IsValid()`、`CanSet()` 和字段 `Kind()`；`CanSet()` 为假时调用 `Set*` 会 panic。访问字段原值前检查 `CanInterface()`；未导出字段即使能看到元信息，也不能随意调用 `Interface()`。见 [ValueOf 示例](../../internal/gomanual/reflectdemo/reflect_valueof.go)及 [结构体示例](../../internal/gomanual/reflectdemo/reflect_struct.go)。

`handleStruct` 现在会拒绝 nil、非指针、nil 指针和非结构体指针；处理有效指针后返回 `nil`。测试同时检查年龄和地址确实被修改。示例中的方法调用仅在 `Talk(string)` 存在时执行；生产代码还应根据实际 API 校验返回值和错误，避免把反射用作未经检查的通用分发器。

## 函数、方法与接口

[函数示例](../../internal/gomanual/reflectdemo/reflect_function.go)演示 `Type.NumIn/NumOut` 与 `Value.Call`。调用前要确认函数有效、实参与形参数量和类型匹配；可变参数展开需区别 `Call` 与 `CallSlice`。`Type.Name()` 对函数类型通常是空字符串，不能把它当作函数名。[接口实现示例](../../internal/gomanual/reflectdemo/reflect_implements.go)使用 `Type.Implements` 检查动态类型的方法集；指针和值的方法集可能不同。

## DeepEqual 的边界

`reflect.DeepEqual` 要求两边类型相同；不同类型的 `int(1)` 和 `int64(1)` 返回假。nil 切片与非 nil 空切片也不相等；指针可递归比较所指值，因此两个不同地址的 `*int` 仍可能相等。函数值只有都为 nil 时才深度相等；含 NaN 的值也有特殊边界。原笔记“不同类型变量如何相等”的 TODO 已核对：`DeepEqual` 不做跨类型转换，如需按业务规则比较应先显式转换或逐字段比较。详见 [DeepEqual 文档](https://pkg.go.dev/reflect#DeepEqual)和[示例](../../internal/gomanual/reflectdemo/reflect_deepequal.go)。

## 验证与范围

在仓库根目录运行：

```bash
go test ./internal/gomanual/reflectdemo ./test/manual -run 'TestHandleStruct|TestReflectionBoundaries|TestReflect'
```

同包测试断言可设置性、未导出字段、`DeepEqual` 边界与结构体修改；`test/manual` 中其余演示主要打印观察结果。旧的 MySQL SQL 拼接反射示例依赖结构体字段和标签，未把它作为通用 SQL 构造器验证。反射规则参考 [标准库 reflect](https://pkg.go.dev/reflect)与 [Go 官方反射文章](https://go.dev/blog/laws-of-reflection)；本文按仓库 Go 1.25 的行为写作。
