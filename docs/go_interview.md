常见的并发模型有哪些？
- 无缓冲Channel
- Sync.WaitGroup 同步组
- Context 上下文

## 并发

Go的并发机制以及它使用的CSP并发模型?
Go GPM 调度模型?
- G: Goroutine 协程
- P: Processer 处理器
- M: Machine 线程

GMP 调度:
> 并发过程中如何进行调度，不同场景下的调度方式是？

并发和并行之间的区别：
- 并发指的是在一个处理器上，一段时间内执行多个任务，更加注重的是任务之间交替执行。（多个事件在同一时间间隔内交替执行）
- 并行指的是多个处理器上同时处理多个任务（不同事件在多个实体是同时执行）

Channel是什么？为什么安全？
- 发送和接收都是原子性的;
- Channel是一个管道，通过管道进行通信；
- Go的并发设计思想就是通过通信来共享内存，而不是通过内存来通信（前者通过Channel、后者通过锁）

Go的几种锁（使用场景）：
- 互斥锁(sync.Mutex)
- 读写锁(sync.RWMutex)-写时不可读、读时不可写、不可并发写、可以并发读
- sync.Map(并发安全的底层原理)


Go语言栈空间管理?
数据竞争如何解决？
进程、线程、协程之间的区别？
Goroutine 数量怎么限制？能在多少个线程上运行？
> Channel sync.WaitGroup

Context使用场景?
原子操作，CAS算法
Go Channel的实现？


并发场景：
- 限制主协程在所有协程完成后才能执行；(sync.WaitGroup)


## 垃圾回收

垃圾回收?

Go 逃逸分析？

## 语法

Go的defer语句?
- 可以理解为栈，先进后出执行顺序
- return在函数中不是原子操作：1. 返回值赋值 2. 调用执行defer语句 3. 返回返回值给调用函数

Go的Selec语句？Select机制？
> Select监听Channel，每个case是一个事件，如果所有case事件阻塞会执行default语句逻辑

Goroutine 退出：
- for-range 检测通道是否关闭
- select-case

nil slice和空slice有什么区别？
> nil slice赋值的时候会出现越界错误，因为只声明了slice，没有实例化对象；

Go slice扩容策略？
slice和map传递过程中有什么区别？
Go 值接收器和指针接收器？
new和make的区别:
> new是传入一个类型，申请内存空间，并初始化为对应的零值，返回该内存空间的指针（主要初始化对象为值类型）make只用来为引用类型对象slice、chan、map的内存创建，返回的是类型本身；

Map的底层实现: 哈希函数、扩容策略、查找性能、碰撞？

接口值和nil进行比较时，会比较接口值的类型T和值V是否都是unset状态

## 网络

Go的http包实现原理？

