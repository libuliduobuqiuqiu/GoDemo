package main

import (
	"fmt"
	"reflect"
)

func handleUser(u interface{}, newName string) {
	tmp := reflect.ValueOf(u)
	if tmp.Kind() == reflect.Ptr {
		var newUser interface{}
		user := tmp.Elem()
		name := user.FieldByName("Name")
		if name.IsValid() && name.CanSet() {
			newUser = user.Interface()
			name.SetString(newName)
		}
		fmt.Println(newUser)
	}
}

func main() {
	n := struct {
		Name string `json:"name"`
	}{Name: "linshukai"}
	handleUser(&n, "new_linshukai")
	fmt.Println(n.Name)
}
