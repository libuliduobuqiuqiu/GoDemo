package genericsdemo

import (
	"encoding/json"
	"fmt"
)

type Person interface {
	GetName() string
}

type Man[T any] struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address T      `json:"address"`
}

type successItemJson[itemClass any] struct {
	Code    int       `json:"error"`
	Message string    `json:"message"`
	Item    itemClass `json:"item,omitempty"`
}

func (m *Man[T]) GetName() string {
	return m.Name
}

func NewSuccessItem[itemClass any](code int, message string, item *itemClass) {
	tmpJson := successItemJson[itemClass]{0, message, *item}
	tmpStr, _ := json.Marshal(tmpJson)

	fmt.Println(string(tmpStr))
	return
}

func marshalMan() {
	a := Man[string]{"linshukai", 22, "zhangdsan"}
	NewSuccessItem[Man[string]](200, "success", &a)
}

func PrintMan() {
	a := Man[string]{"linshukai", 22, "zhangdsan"}
	fmt.Println(a.GetName())
}

// 类型约束接口
type Number interface {
	~int | ~int8 | ~int32 | ~int64
}

func SumInt[n Number](a, b n) n {
	return a + b
}

func Equal[n Number](a, b n) bool {
	return a == b
}
