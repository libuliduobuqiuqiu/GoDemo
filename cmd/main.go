package main

func foo() *int {
	a := 1
	return &a
}

func main() {
	tmp := foo()
	print(*tmp)
}
