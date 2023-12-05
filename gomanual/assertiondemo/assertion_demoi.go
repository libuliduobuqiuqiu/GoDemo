package main

import "fmt"

type usb interface {
	start()
	close()
}

type phone struct {
	name       string
	deviceType string
}

func (p phone) start() {
	fmt.Printf("Phone name is %s, Device Type is %s, Phone start working.....\n", p.name, p.deviceType)
}

func (p phone) close() {
	fmt.Printf("Phone name is %s, Device Type is %s, Phone stop working......\n", p.name, p.deviceType)
}

func (p phone) call() {
	fmt.Printf("Phone calling now......\n")
}

type mouse struct {
	name       string
	deviceType string
	dpi        int
}

func (m mouse) start() {
	fmt.Printf("Mouse name is %s, Device Type is %s, Dpi is %d, Phone start working.....\n", m.name, m.deviceType,
		m.dpi)
}

func (m mouse) close() {
	fmt.Printf("Mouse name is %s, Device Type is %s, Dpi is %d, Phone stop working.....\n", m.name, m.deviceType,
		m.dpi)
}

func (m mouse) click() {
	fmt.Printf("Mouse clicking now......\n")
}

func computerWorking(u usb) {
	u.start()

	if p, ok := u.(phone); ok {
		p.call()
	}
	if m, ok := u.(mouse); ok {
		m.click()
	}

	u.close()
}

func judgeType(dList ...interface{}) {
	for _, device := range dList {
		switch device.(type) {
		case mouse:
			fmt.Println("The device is mouse.")
		case phone:
			fmt.Println("The device is phone.")
		default:
			fmt.Println("The device is not found.")
		}

	}
}

func main() {
	p1 := phone{name: "my phone1", deviceType: "iPhone"}
	p2 := phone{name: "my phone2", deviceType: "XIAOMI"}
	p3 := phone{name: "my phone3", deviceType: "HUAWEI"}

	m1 := mouse{name: "my mouse1", deviceType: "LOGI", dpi: 2000}
	m2 := mouse{name: "my mouse2", deviceType: "SNAKE", dpi: 30000}
	m3 := mouse{name: "my mouse3", deviceType: "HUAWEI", dpi: 500}
	uList := []usb{p1, p2, p3, m1, m2, m3}

	for _, v := range uList {
		computerWorking(v)
	}

	a := 1
	judgeType(m1, p2, m3, a)
}
