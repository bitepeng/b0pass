package ipaddress

import (
	"fmt"
	"net"
)

// GetIP 用来获取有效的内网IP
// 使用sclice存储
func GetIP() ([]string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("net.Interfaces failed, err: %s", err.Error())
	}
	var ip = make([]string, 0)
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil && ipnet.IP.String() != "" {
						ip = append(ip, ipnet.IP.String())
					}
				}
			}
		}
	}
	return ip, nil
}

// GetIP 用来获取有效的内网IP
// 使用map存储
func GetIP2() (map[int]string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("net.Interfaces failed, err: %s", err.Error())
	}
	var ip = make(map[int]string)
	var j = 0
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil && ipnet.IP.String() != "" {
						ip[j] = ipnet.IP.String()
						j++
					}
				}
			}
		}
	}
	return ip, nil
}
