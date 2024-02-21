package main

import "fmt"

type PointerInterface interface {
	AddX()
	PrintX()
}

type SubInterface interface {
	SubX()
}

type Pointer struct {
	X int
	Y int
}

func (p *Pointer) AddX() {
	p.X += 1
}

func (p *Pointer) SubX() {
	p.X -= 1
}

func (p *Pointer) PrintX() {
	fmt.Println(p.X)
}

func UsePointer(p PointerInterface) {
	if tmp, ok := p.(*Pointer); ok {
		tmp.Y += 1
	}

	if tmp, ok := p.(SubInterface); ok {
		tmp.SubX()
	}

	switch p.(type) {
	case Pointer:
		p.PrintX()
	case SubInterface:
		p.SubX()
	}
}

func main() {
	p := &Pointer{X: 2}
	UsePointer(p)
	fmt.Println(p.Y)
	fmt.Println(p.X)
}
