## 反射

什么是反射？为什么需要用到反射？
> 反射是程序能够检测自身结构的机制，程序可以利用反射这个机制动态修改变量的值、调用函数和方法、甚至可以创建新的数据类型和结构。
但是反射带来最大的弊端就是性能问题，其次就是代码脆弱，问题可能在运行过程中才会被发现，以及代码可读性，可维护性。

反射三个经典定律：
1. 反射能把interface{}转换成反射对象；
2. 反射能把反射对象还原成interface{}类型变量；
3. 要修改反射对象，其值必须可设置；
> 反射需要访问类型相关信息时，需要用到reflect.TypeOf；需要修改反射值，需要用到reflect.ValueOf。


### reflect.TypeOf
> reflect.TypeOf函数将变量转换成reflect.Type接口,reflect.Type接口定义了一系列方法，通过这些方法可以获取类型的相关信息;

reflect.Type:
- Kind()：获取变量在Go中的基础类型;
- Elem()：可以判断类型为any的数据结构所存储的元素类型，可接收底层参数必须是指针、切片、数组、通道、映射表;
- Size()：获取对应类型所占的字节大小;
- Comparable()：判断一个类型是否可以被比较;
- Implements()：判断一个类型是否实现某一个接口;
- ConvertibleTo()：判断一个类型是否可以转换成另外一个指定类型；

```go
type rFace interface {
	hello()
}

type myStruct struct{}

func (m myStruct) hello() {
	fmt.Println("hello,world")
}

func BaseUseReflectType() {
	tmpMap := map[string]int{}
	rType := reflect.TypeOf(tmpMap)
	fmt.Println(rType.Kind())
	fmt.Println(rType.Elem())
	fmt.Println(rType.Size())

	tmpStruct := myStruct{}
	sType := reflect.TypeOf(tmpStruct)

	var tmpFace = new(rFace)
	rIface := reflect.TypeOf(tmpFace).Elem()
	fmt.Println(sType.Comparable())
	fmt.Println(sType.Implements(rIface))
}
```
### reflect.ValueOf

> reflect.ValueOf函数将变量转换成reflect.Value结构体变量，包含类型信息和实际值，并且这个结构体定义了很多方法，通过这些
方法可以直接操作Value字段ptr所指向的实际数据；

reflect.Value:
- Type()：获取一个反射值的类型；
- Elem()：获取一个接口或者指针指向的值，如果不是这两种类型则会panic；
- Set()：如果需要通过反射来修改反射值，则必须这个反射值必须是可寻址的；
- Interface()：获取反射值原本的值；

```go
type myStruct struct {
	Name string `json:"Name,omiempty"`
}

func (m myStruct) hello() {
	fmt.Println("hello,world")
}

func BaseUserReflectValue() {
	tmpstruct := &myStruct{Name: "linshukai"}

	rValue := reflect.ValueOf(tmpstruct)
	fmt.Println(rValue.Type())
	fmt.Println(rValue.Kind())

	s := rValue.Elem()
	fmt.Println(s)
	nameField := s.FieldByName("Name")
	nameField.SetString("zhangsan")
	fmt.Println(s.Interface())
}
```
### 函数

> 可以通过反射来获取函数的所有信息，也可以通过反射调用函数；

通过反射操作函数：
- 获取函数信息：TypeOf(),NumIn(),NumOut(),In(),Out(),Name(),String()
- 反射调用函数：ValueOf(),Call()

```go
func CompareMax(a []int, b []int, params ...[]int) int {
	a = append(a, b...)

	for i := 0; i < len(params); i++ {
		a = append(a, params[i]...)
	}

	if len(a) == 0 {
		return 0
	}

	maxNum := a[0]
	for _, v := range a {
		if maxNum < v {
			maxNum = v
		}
	}
	return maxNum
}

func BaseUseReflectFunction() {

	a := []int{1, 2, 3, 4, 1, 5, 19, 22, 311, 332, 11, 2, 33333, 11}
	b := []int{3993, 3, 322, 22, 2211, 23123324, 7567655, 854234, 23}
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)

	rType := reflect.TypeOf(CompareMax)
	fmt.Println(rType.Name())             // 打印函数名称
	fmt.Println(rType.NumIn())            // 打印函数传入参数数量
	fmt.Println(rType.In(0), rType.In(1)) // 打印函数传入参数1,2类型
	fmt.Println(rType.NumOut())           // 打印函数返回值数量
	fmt.Println(rType.Out(0))             // 打印函数第一个返回值类型
	fmt.Println(rType.String())

	rValue := reflect.ValueOf(CompareMax)
	resValue := rValue.Call([]reflect.Value{aValue, bValue})
	for _, v := range resValue {
		fmt.Println(v.Interface())
	}
}
```

### 结构体

> 通过反射访问结构体中的字段、访问字段的Tag、修改字段、调用结构体的方法等，常见用于orm引擎;

通过反射操作结构体：
- 访问字段：NumField(),Field(),FieldByName();
- 修改字段：SetString(),SetInt();
- 访问Tag：structTag.Lookup(),structTag.Get();
- 调用方法：MethodByName(),Call();


```go
type PersonModel struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	money   int
}

func (p PersonModel) Talk(msg string) string {
	return p.Name + ":" + msg
}

func handleStruct(i interface{}) error {

	rType := reflect.TypeOf(i)
	if rType.Kind() == reflect.Pointer {
		s := rType.Elem()
		if s.Kind() == reflect.Struct {
			// 输出结构体字段
			for i := 0; i < s.NumField(); i++ {
				field := s.Field(i)
				// 打印基础信息
				fmt.Println(field.Name, field.Type, field.Tag)
				// 打印标签
				fmt.Println(field.Tag.Get("json"))
			}

			// 修改年龄
			rValue := reflect.ValueOf(i).Elem()

			method := rValue.MethodByName("Talk")
			msg := reflect.ValueOf("hello,world")
			resValues := method.Call([]reflect.Value{msg})
			for _, v := range resValues {
				fmt.Println(v)
			}

			age := rValue.FieldByName("Age")
			if age.IsValid() && age.CanSet() {
				tmpAge := age.Interface()
				age.SetInt(int64(tmpAge.(int) + 1))

			}

			// 修改地址
			address := rValue.FieldByName("Address")
			if address.IsValid() && address.CanSet() {
				tmpAddress := address.Interface()
				address.SetString(tmpAddress.(string) + " Guangzhou city")
			}

		}
	}
	return errors.New("not support reflect type: " + rType.Kind().String())
}

func BaseUseReflectStruct() {
	p := PersonModel{
		Name:    "zhangsan",
		Age:     23,
		Address: "guangdong provience",
		money:   33000,
	}
	fmt.Println(p)
	handleStruct(&p)
	fmt.Println(p)
}
```

### reflect.DeepEqual

> 判断两个变量是否完全相等的函数

Todo: 解析如何判断不同类型变量是否相等(源码)；

```go
func BaseUseReflectDeepEqual() {
	a := make([]int, 100)
	b := make([]int, 100)
	fmt.Println(reflect.DeepEqual(a, b))

	a = append(a, 10)
	fmt.Println(reflect.DeepEqual(a, b))
}
```
### 备注

> https://golang.design/go-questions/stdlib/reflect/how/ 
> https://golang.halfiisland.com/essential/senior/105.reflect.html
