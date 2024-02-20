package main

import "fmt"

type Pointer struct {
	X int
}

func (p *Pointer) AddX() {
	p.X += 1
	fmt.Println(p.X)
}

func main() {
	p := Pointer{X: 2}
	p.AddX()
	fmt.Println(p.X)
}
