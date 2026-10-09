package manual

import (
	"fmt"
	"godemo/internal/gomanual/genericsdemo"
	"os"
	"testing"
)

func SayHello() {
	fmt.Println("hello")
}

func SayBye() {
	fmt.Println("bye")
}

func ExampleSayHello() {
	SayHello()
	// Output:
	// hello
}

func ExampleSayBye() {
	SayBye()
	// Output:
	// bye
}

type myInt int

var benchmarkEqualResult bool

func TestSumInt(t *testing.T) {
	t.Cleanup(func() {
		CleanUpHepler(t)
	})
	a, b := 1, 101
	expected := 102

	actual := genericsdemo.SumInt(a, b)
	if expected != actual {
		t.Errorf("SumInt(%d, %d) expected %d, actual is %d\n", a, b, expected, actual)
	}
}

// 嵌套子测试
func TestGenerics(t *testing.T) {
	t.Run("genericsdemo.SumInt('1','101')", TestSumInt)
	t.Run("genericsdemo.Equal('101','101')", TestEqual)
}

// Helper()帮助函数，测试忽略答应帮助函数中log位置，只会打印测试函数调用者位置
func CleanUpHepler(t *testing.T) {
	t.Helper()
	t.Log("test finished")
}

func TestEqual(t *testing.T) {
	t.Cleanup(func() {
		CleanUpHepler(t)
	})
	a, b := myInt(101), myInt(101)

	if !genericsdemo.Equal(a, b) {
		t.Errorf("equal(%d, %d) error", a, b)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("testing start...")
	code := m.Run()
	fmt.Println("testing end...")
	os.Exit(code)
}

func BenchmarkEqual(b *testing.B) {
	values := []myInt{100, 101}
	var result bool

	for i := 0; i < b.N; i++ {
		result = genericsdemo.Equal[myInt](values[i%len(values)], 101)
	}
	benchmarkEqualResult = result
}
