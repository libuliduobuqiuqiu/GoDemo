package ADops

import "fmt"

const Version = "v2.0.1"

func init() {
	fmt.Println("Init function: ", Version)
}

func Hello() {
	fmt.Println("Hello, World")
}
