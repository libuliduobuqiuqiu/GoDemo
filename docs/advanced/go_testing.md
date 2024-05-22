## 测试

Go tes工具支持几种测试类型：
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

### 示例测试

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
### 单元测试

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
### 基准测试

> 基准测试又叫做性能测试，常用于对程序的内存占用、CPU使用情况、执行耗时等性能指标；

基础使用：
- BenchMark作为测试函数的前缀，传入参数必须为b *testing.B
