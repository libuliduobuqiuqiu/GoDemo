package reflectdemo

import (
	"fmt"
	"reflect"
)

type Device struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	Port    int    `json:"port" db:"port"`
}

type PrintInfo interface {
	GetName() string
	GetAddress() string
	GetPort() int
	SetName(string)
	SetAddress(string)
	SetPort(int)
}

func (d *Device) GetName() string {
	return d.Name
}

func (d *Device) GetAddress() string {
	return d.Address
}

func (d *Device) GetPort() int {
	return d.Port
}

func (d *Device) SetName(name string) {
	d.Name = name
}

func (d *Device) SetAddress(address string) {
	d.Address = address
}

func (d *Device) SetPort(port int) {
	d.Port = port
}

func ReflectMysqlVar(table interface{}) {

	tableType := reflect.TypeOf(table)
	fmt.Println(tableType)

}
