//go:build integration

package gotool

import (
	"godemo/internal/gotool/sshdemo/device"
	"godemo/pkg"
	"testing"
)

func TestDeviceClient(t *testing.T) {

	gbCfg, err := pkg.GetGlobalConfig("")
	if err != nil {
		t.Fatal(err)
	}
	devices := gbCfg.Devices

	if len(devices) > 0 {
		dev := devices[0]
		devClient, err := device.NewDeviceClient(dev.Username, dev.Password, dev.Host, dev.Port)
		if err != nil {
			t.Fatal(err)
		}

		err = devClient.ShellSession()
		if err != nil {
			t.Fatal(err)
		}

		defer devClient.Close()

		devClient.SendCmd("enable\n", "\n")
		devClient.SendCmd("\n")
		devClient.SendCmd("show slb ssl cert 0807-123.crt detail\n")
	}
}
