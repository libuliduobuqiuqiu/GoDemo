package netdemo

import (
	"fmt"
	"net"
)

func IPToBinary(ip net.IP) string {
	if ip.To4() != nil {
		ip = ip.To4()
	} else {
		ip = ip.To16()
	}
	binaryIP := ""
	for _, b := range ip {
		binaryIP += fmt.Sprintf("%08b", b)
	}
	return binaryIP
}

func IPToString(ip net.IP) string {
	ipBytes := ip.To16()
	if ip.To4() != nil {
		return fmt.Sprintf("%d.%d.%d.%d", ipBytes[12], ipBytes[13], ipBytes[14], ipBytes[15])
	} else {
		return fmt.Sprintf("%04x:%04x:%04x:%04x:%04x:%04x:%04x:%04x",
			uint16(ipBytes[0])<<8|uint16(ipBytes[1]), uint16(ipBytes[2])<<8|uint16(ipBytes[3]),
			uint16(ipBytes[4])<<8|uint16(ipBytes[5]), uint16(ipBytes[6])<<8|uint16(ipBytes[7]),
			uint16(ipBytes[8])<<8|uint16(ipBytes[9]), uint16(ipBytes[10])<<8|uint16(ipBytes[11]),
			uint16(ipBytes[12])<<8|uint16(ipBytes[13]), uint16(ipBytes[14])<<8|uint16(ipBytes[15]))
	}
}
