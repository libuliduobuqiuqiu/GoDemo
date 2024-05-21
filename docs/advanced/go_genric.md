## 泛型



类型约束接口
```go
type Number interface{
  int | int8 | int32 | int64
}
```
> 不使用波浪号，类型约束只接受精确匹配的类型，不接受这些类型的定义类型

```go
type Number interface {
  ~int | ~int8 | ~int32 | ~in64
}
```
> 使用波浪号，类型约束接受基础类型以及其所有的定义类型
