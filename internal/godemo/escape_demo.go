package godemo

//go:noinline
func NonEscapingAddress() int {
	x := 1
	p := &x
	return *p
}

//go:noinline
func ReturningAddress() *int {
	x := 1
	return &x
}
