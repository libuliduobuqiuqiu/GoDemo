package ADops

import (
	"fmt"
)

const Version = "v2.0.1"

type DeviceNodeInfo struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

type Device struct {
	DeviceNodeInfo
	ID      string `json:"id"`
	Address string `json:"address"`
}

func init() {
	fmt.Println("Init function: ", Version)
}

func Hello() {
	d := Device{
		ID:             "zssjdoifjsoi-sdfs-dzzxcv",
		Address:        "Guangzhou",
		DeviceNodeInfo: DeviceNodeInfo{Version: "3.0", Name: "zhangsna"},
	}

	d.Version = "2.0"

	fmt.Println("Hello, World")
	fmt.Println(d)
}
