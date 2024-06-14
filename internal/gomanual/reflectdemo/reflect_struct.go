package reflectdemo

import (
	"errors"
	"fmt"
	"reflect"
)

type PersonModel struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	money   int
}

func (p PersonModel) Talk(msg string) string {
	return p.Name + ":" + msg
}

func handleStruct(i interface{}) error {

	rType := reflect.TypeOf(i)
	if rType.Kind() == reflect.Pointer {
		s := rType.Elem()
		if s.Kind() == reflect.Struct {
			// 输出结构体字段
			for i := 0; i < s.NumField(); i++ {
				field := s.Field(i)
				// 打印基础信息
				fmt.Println(field.Name, field.Type, field.Tag)
				// 打印标签
				fmt.Println(field.Tag.Get("json"))
			}

			// 修改年龄
			rValue := reflect.ValueOf(i).Elem()

			method := rValue.MethodByName("Talk")
			msg := reflect.ValueOf("hello,world")
			resValues := method.Call([]reflect.Value{msg})
			for _, v := range resValues {
				fmt.Println(v)
			}

			age := rValue.FieldByName("Age")
			if age.IsValid() && age.CanSet() {
				tmpAge := age.Interface()
				age.SetInt(int64(tmpAge.(int) + 1))

			}

			// 修改地址
			address := rValue.FieldByName("Address")
			if address.IsValid() && address.CanSet() {
				tmpAddress := address.Interface()
				address.SetString(tmpAddress.(string) + " Guangzhou city")
			}

		}
	}
	return errors.New("not support reflect type: " + rType.Kind().String())
}

func BaseUseReflectStruct() {
	p := PersonModel{
		Name:    "zhangsan",
		Age:     23,
		Address: "guangdong provience",
		money:   33000,
	}
	fmt.Println(p)
	handleStruct(&p)
	fmt.Println(p)
}
