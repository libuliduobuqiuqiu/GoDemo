package main

import "fmt"

func testHello1(name string) error {

	if len(name) > 10 {
		return fmt.Errorf("输入的名字长度超过10")
	}
	return nil
}

func testHello2(age int) error {

	if age > 200 {
		return fmt.Errorf("正常人年龄不能超过200岁")
	}
	return nil
}

func testInput(name string) (text string, err error) {
	if len(name) > 10 {
		err = fmt.Errorf("输入的名字长度超过10")
	}

	text = "Hello, My name is " + name
	return
}

func main() {
	err := testHello1("zhangsansdjfiuosjodifjaoisd")
	if err != nil {
		fmt.Println(err)
	}

	if err = testHello2(300); err != nil {
		fmt.Println(err)
	}

	text, err := testInput("Tom")
	fmt.Println(text)
}
