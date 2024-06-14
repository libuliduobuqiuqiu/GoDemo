package reflectdemo

import (
	"fmt"
	"reflect"
)

func CompareMax(a []int, b []int, params ...[]int) int {
	a = append(a, b...)

	for i := 0; i < len(params); i++ {
		a = append(a, params[i]...)
	}

	if len(a) == 0 {
		return 0
	}

	maxNum := a[0]
	for _, v := range a {
		if maxNum < v {
			maxNum = v
		}
	}
	return maxNum
}

func BaseUseReflectFunction() {

	a := []int{1, 2, 3, 4, 1, 5, 19, 22, 311, 332, 11, 2, 33333, 11}
	b := []int{3993, 3, 322, 22, 2211, 23123324, 7567655, 854234, 23}
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)

	rType := reflect.TypeOf(CompareMax)
	fmt.Println(rType.Name())             // 打印函数名称
	fmt.Println(rType.NumIn())            // 打印函数传入参数数量
	fmt.Println(rType.In(0), rType.In(1)) // 打印函数传入参数1,2类型
	fmt.Println(rType.NumOut())           // 打印函数返回值数量
	fmt.Println(rType.Out(0))             // 打印函数第一个返回值类型
	fmt.Println(rType.String())

	rValue := reflect.ValueOf(CompareMax)
	resValue := rValue.Call([]reflect.Value{aValue, bValue})
	for _, v := range resValue {
		fmt.Println(v.Interface())
	}
}
