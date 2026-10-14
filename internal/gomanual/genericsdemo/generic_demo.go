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

func NewSuccessItem[itemClass any](code int, message string, item itemClass) {
	tmpJson := successItemJson[itemClass]{code, message, item}
	tmpStr, _ := json.Marshal(tmpJson)

	fmt.Println(string(tmpStr))
	return
}

func MarshalMan() {
	a := Man[*string]{Name: "linshukai", Age: 22}
	fmt.Println(a.Address, a.Address == nil)
	NewSuccessItem[Man[*string]](200, "success", a)
}

func PrintMan() {
	a := Man[string]{"linshukai", 22, "zhangdsan"}
	fmt.Println(a.GetName())
}

func UnMarshalMan() error {
	a := `{"name": "zhangsan"}`
	m := Man[string]{}

	if err := json.Unmarshal([]byte(a), &m); err != nil {
		return err
	}
	fmt.Println(m)
	return nil
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
