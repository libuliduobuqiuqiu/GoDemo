## 语法

如何查看当前汇编代码？如何定位代码中指定的源码
```bash
go build -gcflags="-S" main.go
```

new和make的区别:
- new是传入一个类型，申请内存空间，并初始化为对应的零值，返回该内存空间的指针（主要初始化对象为值类型）
- make只用来为引用类型对象slice、chan、map的内存创建，返回的是类型本身；


Go的defer语句?
- 可以理解为栈，先进后出执行顺序
- return在函数中不是原子操作：1. 返回值赋值 2. 调用执行defer语句 3. 返回返回值给调用函数

闭包操作

### slice and array
nil slice和空slice有什么区别？
> nil slice赋值的时候会出现越界错误，因为只声明了slice，没有实例化对象；但是如果对nil slice和空slice添加元素时，也是申请内存，然后再赋给
nil slice和空 slice，返回一个新的slice。

slice和array有什么区别？
> 数组是一片连续的内存，固定长度是类型的一部分。切片底层数据是数组，切片是对数组的封装，描述一个数组的片段。切片的类型和长度无关，可以
动态扩容。

Go slice扩容策略？
> Go slice底层数据就是数组，向slice中添加元素，则是向数组添加元素。当数组容量不足容纳新的元素时，则会触发切片的扩容，切片会申请一个更大
的数组，并且将旧slice的元素复制到新申请的数组中，返回一个新的切片，指向新的底层数组。

slice作为函数参数传递：
- Go函数传递，只有值传递，没有引用传递。
- 无论传递的是slice还是slice指针，修改slice底层数组的数据，都会反应到实参slice底层数据。

### map
> map的任务就是设计一种数据结构用来维护一个集合的数据，并且可以对数据进行增删改查；最主要的数据结构：哈希查找表（Hash Map），搜索树（Search Tree）。

几个问题：什么是哈希查找表？什么是哈希冲突？怎么解决哈希冲突？
- 哈希查找表就是通过哈希函数将不同键映射到不同索引上。这要求哈希函数输出范围大于输入范围，但在现实，键的数量远远大于映射的范围；
- 哈希冲突指的就是当不同的键通过哈希函数，映射到同一个bucket上。解决哈希冲突的主要方式：开放寻址法、链表法；
- Go采用的是哈希查找表，通过链表法解决冲突，链表法就是将一个bucket实现成一个链表，落入同一个bucket中的key都会插入这个链表；

关于带comma，和不带comma的map读取方式：
- 带comma的map读取，会返回一个对应value类型的零值和一个bool型变量提示是否在map中。
- 不带comma的map读取，只会返回一个对应value类型的零值。
- 之所以go支持这两语法，因为编译器会在分析代码之后，将两种语法映射到不同的函数;另外根据key的不同类型，编译器也会将查找、插入、删除替换成不同函数处理，优化处理效率。

go语言使用多个数据结构组合表示map，map中最核心的结构就是runtime.hmap
```go
type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32

	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr

	extra *mapextra
}

type mapextra struct {
	overflow    *[]*bmap
	oldoverflow *[]*bmap
	nextOverflow *bmap
}
```
备注：
- count 表示哈希表的中的元素数量
- B 表示持有buckets的数量，但是因为哈希表中的桶的数量是2的倍数，所以该字段会存储对数，也就是len(buckets) == 2^B;
- hash0是哈希种子，引入随机性，创建函数时确定，调用哈希函数作为参数传入；
- oldbuckets是哈希扩容时用于保存之前buckets字段，大小是当前buckets的一半；

哈希表runtime.hmap的桶是runtime.bmap，buckets是一个指针，指向bmap结构体：
```go
type bmap struct {
	tophash [bucketCnt]uint8
}
```
这部分是表面结构，编译期间会加料，动态的创建一个新的结构体：
```go
type bmap struct {
    topbits  [8]uint8
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}
```
备注：
- 桶中最多装8个key;
- 根据key经过哈希计算之后，如果哈希结果是“一类”的，就会落入一个桶中，然后根据计算出来的hash值的高八位，决定key落入桶的哪个位置；


#### map 遍历
> 简单来说，map的遍历过程，就是遍历所有的bucket，以及后面挂的overflow bucket，遍历所有bucket中的cell，将含有key的cell去除key和value。

当map发生扩容时？
> 当map出现扩容的情况，由于扩容过程不是原子操作，最多搬运两次，此时map处于中间态，有可能存在部分元素还在旧bucket，部分元素在新bucket。此时golang遍历应该如何处理这种情况？

当map出现正在扩容的情况，遍历从新的bucket序号顺序开始进行，碰到老的bucket未搬迁的情况，会在老bucket中找到将要搬迁的新key。

#### map 赋值
> 通过对key计算出hash值，根据hash值计算出赋值的位置（插入新key或者更新老key），对相应位置进行赋值。流程核心在于双循环，外层循环遍历bucket和overflow bucket，内存层循环遍历整个bucekt的各个cell。(赋值过程还会考虑是否出现并发写的情况，以及是否正在扩容情况)

#### map 删除

删除流程:
- 检查flags标志，判断写标志是否为1，表示其他协程正在对这个map进行写操作，直接panic。
- 计算key的哈希值，找到落入的bucket。检查此map如果正在扩容操作，直接进行一次搬迁操作。
- 同样是两层循环，核心是找到key的具体位置，找到对应位置，对key或者value进行“清零”操作。
- 最后将hmap的count减一，将对应位置的tophash值置成nil。

#### map 扩容
> 使用哈希表的意义就是为了快速定位查找到key，向map中添加越来越多key之后，发生碰撞的概率也就越来越大，bucket中的8个cell逐渐被填满，查找、插入、删除的效率会降低。

装载因子：
```
loadFactor := count / 2^B
```
count: map的元素个数，2^B表示bucket数量

触发扩容条件：
- 装载因子超过默认阈值6.5
- overflow的buckets数量过多：当B小于15，bucket总数小于2^15时，overflow bucket数量超过2^B;
当B大于等于15，bucket总数大于等于2^15时，如果overflow bucket数量超过2^15；

针对两种扩容条件不同的扩容机制：
- 装载因子超过阈值，本质是装不下了，所以需要将桶的数量翻倍，然后进行迁移。（迁移过程中是在map进行写操作增量进行，减少性能的瞬时抖动；
- overflow bucket过多，本质是上为了重新整理，减少链表。等量扩容完成之后，key都集中到了一个bucket，更加紧凑，提高了查找效率；

#### 额外问题

1. 为什么map是无序的？
> map中的key可能因为扩容操作，导致不在原来的位置，当按照原来的顺序遍历bucket，按照顺序遍历key时，就会出现无序的情况。而在Go中，Go直接在遍历过程中，
并不是在固定序号的bucket开始遍历，是从一个随机值序号的bucket进行遍历，并且从bucket一个随机序号的cell开始遍历，所以就算hard code写死map，每次遍历过程
也会出现不一样顺序的结果。

2. float类型可以作为map的key吗？
> Go中只要是可以比较的类型都能作为map的key。float也可以作为map的key，但是由于精度的问题，可能会出现一些比较诡异的问题，所以使用需要谨慎。

3. 可以边遍历map边删除么？
> Go中map并不是一个线程安全的结构，所以不支持并发同步时读写，会直接触发panic。但是如果在同一个协程中边遍历边删除，并不会检测到同时读写。但是由于删除key时间不同，可能导致遍历的结果也出现差异。
一般这种问题可以通过读写锁sync.RWMuetex，或者sync.Map。

4. 可以对map的元素进行取址？
> 通过正常方式是没办法无法对key、value进行取址，没办法编译运行。但是可以通过unsafe.Pointer获取，但是这种获取方式一旦map发生扩容，map的key和value的地址都会发生改变失效。

5. 如何判断两个map是相等的？
- 都为nil；
- 非空、长度相等，指向同一个map实体对象；
- 相应的key指向value“深度”相等；

6. map是否线程安全？
> map不是线程安全，map结构体中有个字段是写标志位，在查找、赋值、删除时，写标志位为1时，直接panic。当写标志复位之后才能继续之后的操作。

### interface
动态语言的特点
> 变量绑定的类型是不确定的，在运行期间才能确定。函数和方法可以接受任意类型的参数，且调用时不用检查参数类型，不用实现接口。

关于鸭子类型
> 鸭子类型是动态语言一种对象推断策略，关注的是对象的行为，而不是对象类型本身。Go作为一种静态语言通过接口完美支持鸭子类型，实际上是Go编译器作了隐匿的转换工作。

关于接口
> 接口是定义一种规范，描述类的行为和功能，不做具体实现。

关于方法
> 方法是在函数的基础上添加一个接收者，接收者可以是值类型也可以是指针类型。调用方法时，值类型调用者和指针类型调用者，无需固定接收者类型的方法，都可以调用。

值接收者和指针接收者
| - | 值接收者 | 指针调用者 |
| ---- | ----| ----|
| 值类型调用者 | 方法相当于传递调用者的一个副本，类似"传值" | 使用值的引用调用方法|
| 指针类型调用者 | 指针被解引用为值 | 实际上也是“传值”，方法里的操作会影响调用者 |

两者分别在何时使用?
- 值接收器方法，无论调用者是指针还是值都只是修改其副本，不影响调用者，比较适用一些内置的类型；
- 指针接收器方法，当出现大型结构体和数组时，指针接收器更加高效，避免每次调用方法时的值拷贝；

iface和eface?
> iface和eface都是描述接口的底层结构体，iface描述的接口包含方法，eface则是不包含任何方法的空接口(interface)；

接口的动态类型和动态值？
> iface结构体中有两个字段，tab是接口表指针，指向类型信息，data是数据指针，指向具体的数据。这两个分别成为动态类型和动态值。接口值包括动态类型和动态值。
所以当接口值的零值指的是动态类型和动态值为nil，只有两部分都为nil时，才会被认为接口值==nil；

编译器检测类型是否实现了接口
```go
var d types.DeviceHandler = (*DevPHandler)(nil)
```
> 上述赋值语句过程会出现隐式转换，编译器会检测等号右边的类型是否实现了左边等号规定的函数

类型断言和类型转换有什么区别？
- Go类型转换不允许隐式转换，赋值=两边不允许出现类型不一致的变量。类型转换和类型断言本质上是把一个类型转换成另外一个类型，区别在于类型断言是对接口变量进行操作
- 类型转换是针对类型转换前后的变量要相互兼容
- 由于空接口interface{}里面没有定义任何函数，因此所有类型都实现了空接口，所以当一个形参是interface{}时，需要对形参进行断言，得到真实的类型

接口转换的原理？
1. 具体类型转空接口时，_type 字段直接复制源类型的 _type；调用 mallocgc 获得一块新内存，把值复制进去，data 再指向这块新内存。
2. 具体类型转非空接口时，入参 tab 是编译器在编译阶段预先生成好的，新接口 tab 字段直接指向入参 tab 指向的 itab；调用 mallocgc 获得一块新内存，把值复制进去，data 再指向这块新内存。
3. 而对于接口转接口，itab 调用 getitab 函数获取。只用生成一次，之后直接从 hash 表中获取。

通过interface实现多态？
> 多态指的是同一个接口，使用不同的实例执行不同的操作
- 一种类型具有多种类型的能力；
- 允许不同的对象对同一消息进行灵活的反应；
- 非动态语言必须使用继承或者接口来实现；


接口值和nil进行比较时，会比较接口值的类型T和值V是否都是unset状态

### reflect
> Golang提供的反射机制，就是在编译过程中无法确定变量类型，需要在运行时动态对变量更新、访问、调用它们的方法；

常用场景：
1. 不能明确接口调用哪个函数，需要更具传入的参数在运行时决定；
2. 不能明确传入函数时的参数类型，需要运行时处理任意对象；

反射的缺点：
1. 反射代码的可阅读性；
2. 反射代码不能再编译期间就发现相应的问题，有可能运行很久才会发现；
3. 反射性能差;

利用反射机制的DeepEqual深度比较

### unsafe

Go指针的限制：
1. Go指针不能进行数学运算；
2. Go指针不同类型不能相互转换；
3. Go指针不同类型之间不能通过!= 和 == 比较;
4. Go指针不同类型之间不能相互赋值；


## 网络

Go的http包实现原理？

