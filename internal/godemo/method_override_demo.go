package godemo

import "fmt"

type Model interface {
	Use() error
}

type Car struct {
	Name string
}

func (c *Car) Use() error {

	fmt.Println(c.Name)
	return nil

}

type Byd struct {
	*Car
	Address string
}

func (b *Byd) Use() error {
	fmt.Println(b.Address)
	return b.Car.Use()
}

func StartByd() {
	b := Byd{
		Car:     &Car{Name: "byd"},
		Address: "Hangzhou",
	}
	b.Use()
}
