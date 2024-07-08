## 并发
常用并发控制：
- WaitGroup
- Context
- Channel
- Mutex,RWMutex

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
什么是CSP模型？
> 不要通过共享内存进行通信，而是试过通信实现内存共享。
Go依赖CSP模型，基于Channel实现，Go并发原则就是尽量使用Channel，把goroutine作为免费资源使用；

读取一个已关闭的Channel？
> 当Channel已关闭，如果是有缓冲的，如果里面存在数据依然可以正常读取，当里面没有数据，返回的ok标识则为false；

Channel操作总结:
|操作|nil channel|closed channel|not nil,not closed channel|
| ---- | ---- | ---- |---- |
|close| panic | panic | 正常关闭|
|读<-ch| panic | 读到对应类型的零值 | 阻塞或者正常读取数据，缓冲型channel为空时和非缓冲型channel没有等待发送者时会阻塞 |
|写ch<-|panic |panic | 阻塞或者正常写入数据，缓冲型channel buf满时和非缓冲型channel没有等待接收者是会阻塞 |

如何优雅关闭一个Channel？
> 已知关闭一个已经关闭的channel和写入一个已经关闭的channel会出现panic
主要根据receiver和sender数量，分下面几种情况：
1. 1个sender,1个receiver
2. 1个sender,N个receiver
3. N个sender,1个receiver
4. N个sender,N个receiver
解决方案：
- 1,2可以直接关闭sender
- 3可以增加额外的关闭信号通道，receiver通过信号通道发送关闭数据channel指令，sender接收信号通道发送的关闭指令，停止发送数据
- 4可以增加额外两个通道：关闭信号通道，中间人通道，receiver和sender需要关闭通道可以向中间人通道发送信息，让中间人关闭信号通道，sender和receiver接收到关闭信号通道指令然后退出；


Channel是什么？为什么安全？
- 发送和接收都是原子性的;
- Channel是一个管道，通过管道进行通信，数据是先进先出（FIFO)
- Go的并发设计思想就是通过通信来共享内存，而不是通过内存来通信（前者通过Channel、后者通过锁）

for-range读取Channel：
> 遍历获取Channel中管道中的数据, 对于无缓冲通道，for range会在每次接收操作时阻塞，直到其他协程向通道写入数据；对于有缓冲通道，会读取缓冲区的所有
数据，当缓冲区为空，会阻塞等待，直到有其他协程向通道写入数据；

select：
- select是一种管道多路复用的控制结构，通过同时监测多个管道是否可用；
- select一般都有多个case和一个default组成，每个case是一个管道操作；
- 当有多个case可用时，select会**伪随机**的选择一个case执行，当所有case不可用时，则执行default分支，如果没有default分支时，将会阻塞等待；
- select配合for，可实现循环无限监测管道，直到退出，配合time.After设置超时，可实现超时退出，通过break跳出for循环；
- select{}语句中什么case都没有，主协程会无限等待；

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
- 内部实现就是计数器+信号量，协程开始时Add初始化信号量，结束后调用Done，计数-1，直到减为0，主协程会调用Wait一直阻塞等待计数为0时，才会被唤醒；

### Context
Context使用场景?
> 需要统一对多个goroutine执行“取消”动作，常用于并发控制和超时控制；(也可用于传递共享数据)

#### Context接口
Context接口：
- Done() <- chan struct{}:当Context被取消或者到Dealine，返回一个channel
- Err() error: 当channel Done被关闭后，返回context取消原因
- Dealine() (deadline time.Time, ok bool)：返回context截止时间
- Value()：返回之前设置key的value

Context接口的几种实现：
- emptyCtx: Background(),TODO();
- cancelCtx: WithCancel();
- timerCtx: WithDeadline(),WithTimeOut;
- valueCtx: WithValue();

#### emptyCtx
> emptyCtx就是空的上下文，可以通过context.Background(),context.TODO()函数创建；emptyCtx一般当做最顶层的上下文，在创建其他上下文是作为父上下文传入。

#### valueCtx
> valueCtx比较简单，除了内嵌一个Context类型, 并且包含了一对key、value键值对；本身只实现了Value方法，用于寻找value，在当前上下文找不到就会去父上下文找，一直递归寻找;

Context.Value的查找过程是怎样？
- context.Value设置value，通过将context包起来，设置key，value；
- context查找过程中，每个context指向父context，根据指向循环遍历，然后判断比较是否存在对应的value(递归查找过程)；

#### cancelCtx
> cancelCtx和timerCtx都实现了canceler接口;WithCancel()返回一个cancelCtx，以及方法，返回的方法是创建上下文时，不对外暴露其中的cancel方法，通过闭包将其包装为返回值给外部调用;

```go
type canceler interface {
    // removeFromParent 表示是否从父上下文中删除自身
    // err 表示取消的原因
	cancel(removeFromParent bool, err error)
    // Done 返回一个管道，用于通知取消的原因
	Done() <-chan struct{}
}
```
context如何被取消？
- Done()返回一个只读的Channel，通过select，只有当chanel的关闭的时候才能读取到零值；
- Cancel()，关闭Channel，c.done；取消它的所有子节点；从父节点删除自己。达到的效果就是通过关闭channel，然后发送取消信号到它所有子节点；

done()
```go
func (c *cancelCtx) Done() <-chan struct{} {
	d := c.done.Load()
	if d != nil {
		return d.(chan struct{})
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	d = c.done.Load()
	if d == nil {
		d = make(chan struct{})
		c.done.Store(d)
	}
	return d.(chan struct{})
}
```
> 返回ctx上的done通道

cancel()
```go
  d, _ := c.done.Load().(chan struct{})
	if d == nil {
		c.done.Store(closedchan)
	} else {
		close(d)
	}
```
> 调用退出方法，关闭ctx上的通道；

实际使用
```go
func UseCancelCtx() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		case <-time.After(10 * time.Second):
			fmt.Println("TimeOut")
			return
		}
	}
}
```

#### timerCtx

> timerCtx是在cancelCtx的基础上增加了超时机制；提供了两个创建函数WithTimeOut和WithDeadline，前者是指定超时间隔，后者是具体指定超时时间；

```go
func UseTimerCtx() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer cancel()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("上下文取消：", ctx.Err())
				return
			default:
				fmt.Println("等待取消中。。。")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}
```
### sync
> sync标准库提供一些有关并发过程使用的各种工具，sync.Once：懒加载，sync.Pool：对象复用，sync.Map：并发安全的Map，sync/atomic：原子操作;


