常见的并发模型有哪些？
- 无缓冲Channel
- Sync.WaitGroup 同步组
- Context 上下文

Channel是什么？为什么安全？
- 发送和接收都是原子性的;
- Channel是一个管道，通过管道进行通信；
- Go的并发设计思想就是通过通信来共享内存，而不是通过内存来通信（前者通过Channel、后者通过锁）

nil slice和空slice有什么区别？
> nil slice赋值的时候会出现越界错误，因为只声明了slice，没有实例化对象；


Go的并发机制以及它使用的CSP并发模型?
Go GPM 调度模型?
进程、线程、协程之间的区别？
数据竞争如何解决？
垃圾回收?

Go的defer语句?
- 可以理解为栈，先进后出执行顺序
- return在函数中不是原子操作：1. 返回值赋值 2. 调用执行defer语句 3. 返回返回值给调用函数

Go的Select语句？ 
