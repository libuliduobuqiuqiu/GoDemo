package netdemo

import (
	"fmt"
	"net"
)

// 扫描指定网段的IP地址
func GetAllIP(subnet, startIP, endIP string) ([]string, error) {
	var count int
	var ipAddrs []string
	var startIPAddr, endIPAddr net.IP

	_, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}

	if startIP != "" {
		startIPAddr = net.ParseIP(startIP)
		if !ipNet.Contains(startIPAddr) {
			return nil, fmt.Errorf("%s 不属于 %s 网段", startIP, subnet)
		}

	} else {
		startIPAddr = net.IP(ipNet.Mask)
	}

	if endIP != "" {
		endIPAddr = net.ParseIP(endIP)
		if !ipNet.Contains(endIPAddr) {
			return nil, fmt.Errorf("%s 不属于 %s 网段", endIP, subnet)
		}
	}

	for ip := startIPAddr; ipNet.Contains(ip); inc(ip) {
		if count > 99999 || ip.String() == endIPAddr.String() {
			break
		}
		ipAddrs = append(ipAddrs, ip.String())
		count++
	}

	return ipAddrs, nil
}

func inc(ip net.IP) {
	fmt.Println(ip)
	for j := len(ip) - 1; j >= 0; j-- {
		fmt.Println(ip[j])
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func Increment(ip net.IP) {
	fmt.Println(ip)
	for j := len(ip) - 1; j >= 0; j-- {
		fmt.Println(j, ip[j])
	}

	for j := len(ip) - 1; j >= 0; j-- {
		fmt.Println(j, ip[j])
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
