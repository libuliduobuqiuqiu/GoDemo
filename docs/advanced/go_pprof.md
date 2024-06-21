## pprof

> pprof是Golang自带的性能分析工具，可以分析程序的CPU使用情况、内存使用情况、阻塞情况等；

使用方式：
- rutime/pprof，使用StartCPUProfile和StopCPUProfile生成分析样本;
- net/http/pprof，采集http server运行时数据进行性能分析，通过http服务调用Profile分析样本，底层还是调用runtime/pprof;

### runtime/pprof

#### 收集数据

1. go test
> go内置的testing框架中内置支持了pprof

testing源码：
```go
func (m *M) before() {
	if *memProfileRate > 0 {
		runtime.MemProfileRate = *memProfileRate
	}
	if *cpuProfile != "" {
		f, err := os.Create(toOutputDir(*cpuProfile))
		if err != nil {
			fmt.Fprintf(os.Stderr, "testing: %s\n", err)
			return
		}
		if err := m.deps.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "testing: can't start cpu profile: %s\n", err)
			f.Close()
			return
		}
	}
  // 通过执行命令时传入标识cpuprofile开启性能检测
}
```
本质上还是调用了pprof:
```go
func (TestDeps) StartCPUProfile(w io.Writer) error {
	return pprof.StartCPUProfile(w)
}
```
详细使用：
```bash
go test -v pprof_test.go -bench . -cpuprofile cpu.profile -memprofile mem.profile
```

2. 手动开启pprof
> pprof支持内存、CPU性能情况
```go
func AnalysisFibByPprof() error {
	// f, err := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0666)
	f, err := os.Create("cpu.profile")
	if err != nil {
		return err
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var n = 10
	for i := 0; i <= 5; i++ {
		fmt.Printf("Fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
	return nil
}
```
#### 展示数据

展示数据的两种方式：
- 命令行：go tool pprof 查看分析数据；
- 浏览器：go tool pprof -http 查看可视化数据；
（浏览器查看可视化数据需要安装graphviz）
```bash
go tool pprof cpu.profile
go tool pprof -http=0.0.0.0:8090 cpu.profile
```
### net/http/pprof

采集http服务分析数据，只需要直接导入net/http/pprof包即可
```go
import _ "net/http/pprof"
```
只需要在启动的http服务后面，加上路径/debug/pprof/,直接在浏览器访问；同时还可以通过pprof命令在命令行访问
```bash
go tool pprof http://172.23.227.139:8989/debug/pprof/allocs
```
浏览器可视化分析
```bash
go tool pprof -http=0.0.0.0:8090 http://172.23.227.139:8989/debug/pprof/allocs
```

## Trace

> 主要跟踪程序运行过程的工具，帮助分析程序的事件流、Goroutine的创建和销毁、垃圾回收、系统调用等；


#### 收集数据

1. 项目调用runtime/trace包
```bash
func AnalysisFibByTrace() error {
	f, err := os.Create("cpu.trace")
	if err != nil {
		return err
	}

	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	var n = 10
	for i := 0; i <= 5; i++ {
		fmt.Printf("Fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
	return nil
}
```
2. http服务引入net/http/pprof包
```bash
curl -o trace.out http://127.0.0.1:6060/debug/pprof/trace?seconds=30
```
3. 单元测试内置
```bash
go test -v pprof_test.go -trace trace.out -bench .
```

#### 分析数据
```bash
go tool trace -http=:8090 trace.out
```

## TODO

1. 为什么直接导入net/http/pprof包之后就能够通过url访问到对应分析数据

## 参考链接

