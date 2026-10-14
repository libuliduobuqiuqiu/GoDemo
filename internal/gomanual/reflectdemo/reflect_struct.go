package reflectdemo

import (
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
	if rType == nil || rType.Kind() != reflect.Pointer || reflect.ValueOf(i).IsNil() || rType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a non-nil pointer to struct, got %T", i)
	}
	s := rType.Elem()
	for n := 0; n < s.NumField(); n++ {
		field := s.Field(n)
		fmt.Println(field.Name, field.Type, field.Tag)
		fmt.Println(field.Tag.Get("json"))
	}

	rValue := reflect.ValueOf(i).Elem()
	method := rValue.MethodByName("Talk")
	if method.IsValid() && method.Type().NumIn() == 1 && method.Type().In(0) == reflect.TypeOf("") {
		for _, v := range method.Call([]reflect.Value{reflect.ValueOf("hello,world")}) {
			fmt.Println(v)
		}
	}

	age := rValue.FieldByName("Age")
	if age.IsValid() && age.CanSet() && age.Kind() == reflect.Int {
		age.SetInt(age.Int() + 1)
	}
	address := rValue.FieldByName("Address")
	if address.IsValid() && address.CanSet() && address.Kind() == reflect.String {
		address.SetString(address.String() + " Guangzhou city")
	}
	return nil
}

func BaseUseReflectStruct() {
	p := PersonModel{
		Name:    "zhangsan",
		Age:     23,
		Address: "guangdong provience",
		money:   33000,
	}
	fmt.Println(p)
	if err := handleStruct(&p); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
}
