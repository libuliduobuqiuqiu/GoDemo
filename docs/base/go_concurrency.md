## 并发
常用并发控制：
- WaitGroup
- Context
- Channel
- Mutex,RWMutex

### Go Scheduler
> Go程序主要包含Go Program、Go Runtime，即用户程序和运行时。他们之间通过函数调用实现内存分配，垃圾回收，并发调度等功能。用户程序进行的
系统调用都会被Runtime拦截，以此帮助它进行调度以及垃圾回收工作。一般用户程序无法直接和系统内核交换，都是通过Runtim间接交互。
Go Scheduler是Runtime最重要的一部分，Runtime维护所有的Goroutines，并通过scheduler进行调度，goroutines和threads是独立分开，但是goroutines
要依赖threads才能执行。

并发和并行之间的区别：
- 并发指的是在一个处理器上，一段时间内执行多个任务，更加注重的是任务之间交替执行。（多个事件在同一时间间隔内交替执行）
- 并行指的是多个处理器上同时处理多个任务（不同事件在多个实体是同时执行）

Goroutine和Thread之间的区别？
- 内存占用：创建Goroutine的栈内存消耗2KB，栈空间不够用，会自动扩容；创建Thread的栈内存消耗1MB。
- 创建和销毁：Goroutine创建和销毁是Go runtime进行管理消耗资源小，属于用户级；Thread创建和销毁需要和操作系统进行交互，属于内核级。
- 切换：Goroutine切换是由Go runtime进行管理，切换只需要少量的上下文信息，开销非常小；Thread切换需要保存和回复线程的上下文信息，需要较多的时间和资源。

#### 设计原理
GMP 调度模型:
- G(Goroutine): 表示Goroutine，它是一个待执行的任务。
- M(Machine)：表示操作系统的线程，它由操作系统调度器调度和管理。
- P(Processor)：表示处理器，负责调度和管理它的Goroutine队列，将Goroutine分配给对应的M执行。(Processor的数量是GOMAXPROCS，默认为CPU核心数量)
- GRQ：存储全局可运行的Goroutines。
- LRQ：存储本地（P）上可运行的Goroutines。

goroutine可能发生调度：
- go关键字,创建一个新的goroutine，go scheduler会考虑调度。
- GC，由于进行GC的Goroutine也需要在M上进行，因此肯定会发生调度。
- 系统调用，当一个goroutine进行系统调用，会阻塞M，所以它会被调度走，其他的Goroutine会被调度上来。
- 内存同步访问，atomic、channel、mutex操作会使goroutine阻塞，因此会被调度走。

什么是M:N模型？
> Go runtime会负责goroutine的创建和销毁，Runtime会在程序启动时，会启动M个线程（CPU执行调度单位），之后创建的N个Goroutines会在线程上
执行。
在同一个时刻，在一个线程上只能执行一个Goroutine，当Goroutine执行过程阻塞了，调度器会将当前的Goroutine调走，让其他Goroutine执行。

什么是工作窃取？
> Go scheduler需要保证runnable goroutines均匀分布在P上运行的M。
M需要绑定P，获取P的LRQ上的Goroutine，然后才能执行。当M上出现系统调用阻塞，P会释放和M的绑定，找到其他可用的M执行P上LRQ的Goroutine。
当P的LRQ的Goroutine为空之后，并且这是GRQ上也没有Goroutine，P会从其他P的LRQ上“偷取”一半的G。
(当系统调用阻塞时M才会解绑P；而当非阻塞I/O，比如一些网络I/O，会有Net Poller处理，会将I/O操作注册到事件通知机制中，Net Poller会异步等待
这些事件完成，不会阻塞当前Goroutine或M，当网络事件完成后Net poller会将Goroutine注入到GRQ，P会调度一个空闲的M执行。)

协作式调度和抢占式调度？
> 协作式调度依靠调度方主动弃权；抢占式调度依靠调度器强制将被调方被动中断。

基于协作的抢占式调度：
1. 编译器会在被调函数前插入runtime.morestack(抢占标志检测)；
2. Go语言会在运行时，会在垃圾回收暂停程序、系统监控发现Goroutine运行超过10ms，会设置Goroutine的stackguard0为StackPreempt(抢占标志位);
3. 在发生函数调用时，可能会执行runtime.morestack，它调用的runtime.newstack会检测Goroutine的stackguard0字段是否为StackPreempt；
4. 如果是StackPreempt则触发抢占让出线程；
(这里的函数检测是编译器插入的，但是需要函数调用作为入口触发抢占，所以这算是一个协作式的抢占调度？)

#### 基于信号的抢占式调度
> 在以往的基于协作的抢占式调度只有在函数调用离设置抢占标志位，对于不是函数就没有办法。如果是一个纯算法循环计算，Go调度器就没办法，可能还是
会出现饥饿的问题。
```go
func main() {
	var x = 0
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		fmt.Println(i)
		go func() {
			for {
				x++
			}
		}()
	}

	fmt.Println("hello,world")
}
```
按照协作式调度的方法，启动和CPU核心相等的goroutine，goroutine无限循环。这样可能导致M和P都被占满了，一直无限循环。
其他的Goroutine想要执行，会一直等待，出现饥饿。

基于信号的抢占式调度流程：
1. M注册一个SIGURG信号处理的函数：sighandler。
2. sysmon启动间隔性性能监控，发现超过Goroutine运行超过10ms，向M发送抢占信号。
3. M接收到信号之后，内核中断器执行的代码，执行注册信号处理函数，将当前Goroutine的状态从_Grunning改成Grunnable，把抢占的
Goroutine放到全局队列中，M继续从P中找其他的Goroutine执行。
4. 被抢占的P再次调过来会继续原来的执行流。

#### 源码阅读
调度启动（初始化）：
- 设置maxmcount=10000,Go能够创建的最大线程，获取GOMAXPROCS环境变量。
- 根据GOMAXPROCS环境变量，调用runtime.procsize更新程序中的处理器数量。
- 调用runtime.procsize是初始化的最后一步，调度器在这之后完成响应数量处理器启动，等待用户创建的
Goroutine并为Goroutine调度处理器资源。

runtime.procsize:
1. 如果全局变量 allp 切片中的处理器数量少于期望数量，会对切片进行扩容；
2. 使用 new 创建新的处理器结构体并调用 runtime.p.init 初始化刚刚扩容的处理器；
3. 通过指针将线程 m0 和处理器 allp[0] 绑定到一起；
4. 调用 runtime.p.destroy 释放不再使用的处理器结构；
5. 通过截断改变全局变量 allp 的长度保证与期望处理器数量相等；
6. 将除 allp[0] 之外的处理器 P 全部设置成 _Pidle 并加入到全局的空闲队列中；
(初始化处理器，保证处理器队列和期望处理器数量相等，绑定m0和处理器allp[0]，将除allp[0]之外的处理器放入全局的空闲队列中)


#### 额外问题
Goroutine 数量怎么限制？能在多少个线程上运行？
> Channel sync.WaitGroup

Go的Selec语句？Select机制？
> Select监听Channel，每个case是一个事件，如果所有case事件阻塞会执行default语句逻辑

### Channel
什么是CSP模型？
> 不要通过共享内存进行通信，而是试过通信实现内存共享。
Go依赖CSP模型，基于Channel实现，Go并发原则就是尽量使用Channel，把goroutine作为免费资源使用；

Channel是什么？为什么安全？
- 发送和接收都是原子性的;
- Channel是一个管道，通过管道进行通信，数据是先进先出（FIFO)
- Go的并发设计思想就是通过通信来共享内存，而不是通过内存来通信（前者通过Channel、后者通过锁）

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
- 4可以增加额外两个通道：关闭信号通道、中间人通道，receiver和sender需要关闭通道可以向中间人通道发送信息，让中间人关闭信号通道，sender和receiver接收到关闭信号通道指令然后退出；

Channel发送和接收的本质是什么？
> All transfer of value on the go channels happens with the copy of value.
发送和接收的本质上都是“值的拷贝”

Channel什么情况下发生泄漏？
> 泄露的原因一般为goroutine操作channel后，处于发送或者接收阻塞状态，而channel处于满或空的状态一直不会改变。垃圾回收器不会回收此类资源。

for-range读取Channel：
> 遍历获取Channel中管道中的数据, 对于无缓冲通道，for range会在每次接收操作时阻塞，直到其他协程向通道写入数据；对于有缓冲通道，会读取缓冲区的所有
数据，当缓冲区为空，会阻塞等待，直到有其他协程向通道写入数据；

select：
- select是一种管道多路复用的控制结构，通过同时监测多个管道是否可用；
- select一般都有多个case和一个default组成，每个case是一个管道操作；
- 当有多个case可用时，select会**伪随机**的选择一个case执行，当所有case不可用时，则执行default分支，如果没有default分支时，将会阻塞等待；
- select配合for，可实现循环无限监测管道，直到退出，配合time.After设置超时，可实现超时退出，通过break跳出for循环；
- select{}语句中什么case都没有，主协程会无限等待；

Channel的几种使用场景：
- 停止信号
- 超时控制、定时执行某个任务
- goroutine并发数控制
- 生产者消费者模型

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

### 同步原语与锁
> Go在sync包里面提供了基本的原语sync.Mutex:锁,sync.RWMutex:读写锁,sync.WaitGroup:等待组,sync.Cond:同步条件通知,sync.Once:懒加载，
sync.Pool:对象复用，sync.Map:并发安全的Map，sync/atomic:原子操作;

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

