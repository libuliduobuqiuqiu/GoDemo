package ADops

import (
	"bytes"
	"fmt"
	"strings"
)

const Version2 = "Cron v.2.0"

var buf bytes.Buffer

type BasicErrorMsg struct {
	Code int    `json:"error"`
	Msg  string `json:"message"`
}

type requestError struct {
	BasicErrorMsg
	stacks []*Stack
	info   string
	err    error
}

type Stack struct {
	File string
	Line int
	Func string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Student struct {
	Person
	Grade string
}

func init() {
	s := Student{
		Person: Person{Name: "zhangsan", Age: 12},
		Grade:  "六年级",
	}
	fmt.Println(s)
	fmt.Println("Init function: ", Version2)

	words := "hello, world, my name is wangyangming, xiaoxiao"
	splitWords(words)
	fmt.Println("words: ", words)

}

func splitWords(words string) {
	wordList := strings.Split(words, ",")

	words = wordList[0]
	fmt.Println(words)
	for _, word := range wordList {
		fmt.Println(word)
	}

	buf.WriteString("Update dns_pool set ")
	buf.WriteString(" device_group_id = ?")
	sqlStr := buf.String()
	fmt.Println(sqlStr)
}
