## 并发

### Goroutine

进程、线程、协程之间的区别？

CSP模型
> CSP模型核心思想：不要通过共享内存进行通信，而是要通过通信共享内存；

Go GMP 调度模型?
- G(Goroutine): 协程，计算任务，需要执行的代码和上下文
- M(Machine)：系统线程，执行实体，想要在CPU执行代码必须通过线程
- P(Processer): 处理器, 为Goroutine和Machine的调度器

GMP 调度:
> 并发过程中如何进行调度，不同场景下的调度方式是？


并发和并行之间的区别：
- 并发指的是在一个处理器上，一段时间内执行多个任务，更加注重的是任务之间交替执行。（多个事件在同一时间间隔内交替执行）
- 并行指的是多个处理器上同时处理多个任务（不同事件在多个实体是同时执行）

Goroutine 数量怎么限制？能在多少个线程上运行？
> Channel sync.WaitGroup

Go的Selec语句？Select机制？
> Select监听Channel，每个case是一个事件，如果所有case事件阻塞会执行default语句逻辑

Goroutine 退出：
- for-range 检测通道是否关闭
- select-case


### Channel
Channel是什么？为什么安全？
- 发送和接收都是原子性的;
- Channel是一个管道，通过管道进行通信；
- Go的并发设计思想就是通过通信来共享内存，而不是通过内存来通信（前者通过Channel、后者通过锁）


Go Channel的实现？

### 锁

Go的几种锁（使用场景）：
- 互斥锁(sync.Mutex)
- 读写锁(sync.RWMutex)-写时不可读、读时不可写、不可并发写、可以并发读
- sync.Map(并发安全的底层原理)

数据竞争如何解决？
> 锁、Channel、CAS操作

原子操作，CAS算法

### WaitGroup

并发场景：
- 限制主协程在所有协程完成后才能执行；(sync.WaitGroup)


### Context
Context使用场景?
> 需要统一对多个goroutine执行“取消”动作，常用于并发控制和超时控制；(也可用于传递共享数据)

Context.Value的查找过程是怎样？
- context.Value设置value，通过将context包起来，设置key，value；
- context查找过程中，每个context指向父context，根据指向循环遍历，然后判断比较是否存在对应的value(递归查找过程)；

context如何被取消？
- Done()返回一个只读的Channel，通过select，只有当chanel的关闭的时候才能读取到零值；
- Cancel()，关闭Channel，c.done；递归取消它的所有子节点；从父节点删除自己。达到的效果就是通过关闭channel，然后递归发送取消信号到它所有子节点；

Context接口：
- Done() <- chan struct{}:当Context被取消或者到Dealine，返回一个channel
- Err() error: 当channel Done被关闭后，返回context取消原因
- Dealine() (deadline time.Time, ok bool)：返回context截止时间
- Value()：返回之前设置key的value

Go语言栈空间管理?

## 垃圾回收

垃圾回收?

Go 逃逸分析？

## 语法

new和make的区别:
- new是传入一个类型，申请内存空间，并初始化为对应的零值，返回该内存空间的指针（主要初始化对象为值类型）
- make只用来为引用类型对象slice、chan、map的内存创建，返回的是类型本身；


Go的defer语句?
- 可以理解为栈，先进后出执行顺序
- return在函数中不是原子操作：1. 返回值赋值 2. 调用执行defer语句 3. 返回返回值给调用函数

### slice

nil slice和空slice有什么区别？
> nil slice赋值的时候会出现越界错误，因为只声明了slice，没有实例化对象；

Go slice扩容策略？

slice和map传递过程中有什么区别？

### map

Map的底层实现? 
哈希函数?
扩容策略?
查找性能?
碰撞？

### interface
Go 值接收器和指针接收器？

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



## 网络

Go的http包实现原理？

