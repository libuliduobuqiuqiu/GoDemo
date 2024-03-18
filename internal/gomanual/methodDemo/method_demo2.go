package methoddemo

import (
	"fmt"
)

type DeviceHandler interface {
	LoadDevice() string
	LoadState() string
}

type Device struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (d Device) LoadDevice() string {
	d.Age = 100
	fmt.Printf("%p\n", &d)
	return "device"
}

func (d *Device) LoadState() string {
	return "up"
}

func showDevice(d interface{}) {
	if device, ok := d.(*Device); ok {
		device.LoadDevice()
		fmt.Println(device.Age)
	}

}

func ScanDevice() {
	d := Device{}
	fmt.Printf("%p\n", &d)
	showDevice(&d)
}
