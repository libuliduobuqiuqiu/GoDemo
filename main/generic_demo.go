package main

import (
	"encoding/json"
	"fmt"
)

type Man struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type successItemJson[itemClass any] struct {
	Code    int       `json:"error"`
	Message string    `json:"message"`
	Item    itemClass `json:"item,omitempty"`
}

func NewSuccessItem[itemClass any](code int, message string, item *itemClass) {
	tmpJson := successItemJson[itemClass]{0, message, *item}
	tmpStr, _ := json.Marshal(tmpJson)

	fmt.Println(string(tmpStr))
	return
}

func marshalMan() {

	a := Man{"linshukai", 22}
	NewSuccessItem[Man](200, "success", &a)

}
