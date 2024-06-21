## Go testing

> 简单介绍Go官方支持的测试，以及日常开发过程中使用到的测试场景；

Go test工具支持几种测试类型：
1. 示例测试
2. 单元测试
3. 基准测试
4. 模糊测试

### 规范

测试规范：
- 测试包，测试文件通常单独放在一个包中，包通常命名为test
- 测试文件，测试文件通常以_test.go结尾，如果需要更仔细的划分，可根据测试文件中的测试类型（例如：example_、benchmark_、fuzzy_为前缀）
- 测试函数，不同的测试类型，测试函数命名前缀不同（示例测试：Example、单元测试：Test、基准测试：BenchMark、模糊测试：Fuzzy）

### 执行测试

go test 命令
- 执行当前目录下的所有测试用例(.)；后面还可以接指定测试文件名称，单独测试该测试文件下的测试用例；（可支持正则表达式）
- 执行当前测试文件下指定的测试用例（-run 测试用例名称，可支持正则表达式匹配）
- 执行当前的基准测试（-bench 测试用例名称，可支持正则表达式）

```bash
go test -v sum_test.go -bench .
```
测试当前sum_test.go文件下的所有基准测试、示例测试、单元测试、模糊测试，并且详细输出具体测试信息
```bash
go test -v sum_test.go -bench . -run ^$
```
只测试sum_test.go文件下的所有基准测试，相当于-run限制了其余测试用例的名称

执行测试命令常用参数：
- -v：输出更详细的测试日志；
- -count n：运行测试n次，默认是1次；
- -bench regexp：选中regexp匹配的基准测试；
- -benchmem：统计基准测试的内存分配；
- -benchtime t: 基准测试运行足够的迭代（b.N)满足时间t（默认1s)；也可以通过Nx的方式控制b.N迭代次数；
- -run regexp：选中的regexp匹配的测试用例;(单元测试、示例测试)
- -fuzz regexp：选中regexp匹配的模糊测试；
- -cpuprofile cpu.out：统计测试过程中的CPU使用情况并写入文件；
- -memprofile mem.out：统计测试过程中的内存使用情况并写入文件；
- -trace trace.out：将执行追踪情况写入文件；

### 示例测试(Example)

> 一般用于展示功能的使用方法，起到文档作用

使用规范：
- Example为测试函数前缀
- 使用Output注释来检测输出，如果没有Output注释则不会视为示例测试，不会被go执行
- Output输出支持，单行、多行、无序（Unorderd output)

```go
func ExampleSayHello() {
	SayHello()
	// Output:
	// hello
}
```
### 单元测试(Test)

> 对程序中最小可测单元进行测试（单元大小取决开发者：结构体、包、函数、类型）

基础使用：
- Test作为测试函数的前缀，传入参数必须为t *testing.T
- t.Cleanup注册一个收尾函数结束测试
- t.Helper()可以标记当前函数为帮助函数，测试打印过程只会显示帮助函数调用者的位置
- t.Run()可以在测试用例中调用其他的测试用例，这种嵌套的测试用例称为子测试
- 测试数据可以通过结构体切片的形式（表格风格）声明，更加方便直观

Cleanup函数、Helper函数
```go
func CleanUpHepler(t *testing.T) {
	t.Helper()
	t.Log("test finished")
}

func TestEqual(t *testing.T) {
	t.Cleanup(func() {
		CleanUpHepler(t)
	})
	a, b := myInt(101), myInt(101)

	if !genericsdemo.Equal[myInt](a, b) {
		t.Errorf("equal(%d, %d) error", a, b)
	}
}
```
### 基准测试(BenchMark)

> 基准测试又叫做性能测试，常用于对程序的内存占用、CPU使用情况、执行耗时等性能指标；

基础使用：
- BenchMark作为测试函数的前缀，传入参数必须为b *testing.B

```go
func BenchmarkConcatDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringDirct(longString)
	}
}
```

运行命令
```bash
go test bench_concat_test.go -v -bench . -count=2 -cpu=2,4,8 -v -benchmem
```
运行结果
```
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkConcatDirect
BenchmarkConcatDirect-2        	       3	 432794775 ns/op	4040056592 B/op	    9999 allocs/op
BenchmarkConcatDirect-2        	       3	 414416797 ns/op	4040056592 B/op	    9999 allocs/op
BenchmarkConcatDirect-4        	       3	 456195014 ns/op	4040067536 B/op	   10113 allocs/op
BenchmarkConcatDirect-4        	       3	 464858539 ns/op	4040068560 B/op	   10123 allocs/op
BenchmarkConcatDirect-8        	       3	 477814948 ns/op	4040084624 B/op	   10291 allocs/op
BenchmarkConcatDirect-8        	       3	 459436423 ns/op	4040085168 B/op	   10296 allocs/op
BenchmarkConcatWithBuilder
BenchmarkConcatWithBuilder-2   	    1620	    755694 ns/op	 4128176 B/op	      29 allocs/op
BenchmarkConcatWithBuilder-2   	    1903	    747043 ns/op	 4128176 B/op	      29 allocs/op
BenchmarkConcatWithBuilder-4   	    1628	    831750 ns/op	 4128187 B/op	      29 allocs/op
BenchmarkConcatWithBuilder-4   	    1412	    842513 ns/op	 4128188 B/op	      29 allocs/op
BenchmarkConcatWithBuilder-8   	    1423	    855130 ns/op	 4128207 B/op	      29 allocs/op
BenchmarkConcatWithBuilder-8   	    1448	    821266 ns/op	 4128206 B/op	      29 allocs/op
PASS
ok  	command-line-arguments	25.414s
```
- goos: 运行操作系统
- goarch：CPU架构
- cpu：CPU型号信息
- 第一列：-2、-4、-8分别代表测试使用的最大CPU数量；
- 第二列：代表循环迭代的次数b.N;
- 第三列：每一次循环消耗的时间；
- 第四列：每一次循环分配的内存大小；
- 第五列：表示每一次循环内存分配的次数；

### 模糊测试

> 模糊测试，可以通过语料库生成随机的测试数据；（随机测试更好的测试程序的边界条件）

使用基础：
- 函数前缀需要使用Fuzz，传入参数必须为f *testing.F
- -fuzz 执行模糊测试是通过语料库，生成随机测试数据进行测试


模糊测试字符串反转函数
```go
package manual

import "testing"

func Reverse(s string) string {
	tmp := []rune(s)
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return string(tmp)
}

func FuzzReverse(f *testing.F) {
	data := []string{
		"hello,world",
		"what",
		"fuzzy",
		"zhangsan",
	}

	for _, v := range data {
		f.Add(v)
	}

	f.Fuzz(func(t *testing.T, a string) {

		second := Reverse(a)
		third := Reverse(second)

		if a != third {
			t.Fatalf("before: %q, after: %q", a, third)
		}
	})
}
```
- f.Add 往语料库添加数据
- f.Fuzz 添加单元测试函数

### 参考

testing库：
> https://pkg.go.dev/testing@go1.22.3

模糊测试:
> https://go.dev/doc/security/fuzz/

参考文章：
> https://golang.halfiisland.com/essential/senior/120.test.html
