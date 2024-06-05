## 泛型

> 泛型就是为了解决执行逻辑和类型无关的问题

### 泛型结构

泛型写法：
- 类型形参
- 类型约束
- 类型实参


#### 备注

使用泛型的常见问题：
- 泛型不能作为基础类型；
- 泛型类型无法使用断言；
- 匿名结构不支持泛型；
- 匿名函数不支持自定义泛型，但是在闭包中可以使用已有的泛型类型；
- 方法不支持泛型方法，但是接收器可以拥有泛型形参；


### 类型集

> 1.18之后，接口的定义被改为类型集，类型集主要用于类型约束，不能用于类型声明。类型约束指定了类型形参可接受的类型集合；
```go
type SingedInt interface {
	int8 | int16 | int32 | int64
}

type UnSingedInt interface {
	uint8 | uint16 | uint32 | uint64
}

type Integer interface {
	SingedInt | UnSingedInt
}

type MyNumber interface {
	Integer
	SingedInt
}

type MyInteger interface {
	SingedInt
	UnSingedInt
}
```
类型集使用：
- Integer 是SingedInt和UnSingedInt的并集;
- MyNumber 是SingedInt和Integer的交集;
- MyInteger 是SingedInt和UnSingedInt的交集，但交集为空所以是空集;
- 空接口interface{}是所有类型集的集合，包含所有类型；

#### 特殊场景

当声明一个新类型，底层类型包含在类型集中，当这个新类型正常传入时无法编译通过；解决方法可以考虑类型集声明时加上波浪号。
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

#### 备注

使用类型集的常见问题：
- 带有方法集的接口无法并入类型集;
- 类型集无法当做类型实参使用;
- 类型集中类型并集，对于非接口类型不允许有交集，对于接口类型可以允许有交集；

