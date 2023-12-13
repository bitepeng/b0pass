package nets

import (
	"fmt"
	"net"
	"strings"
)

// GetOutBoundIP 使用UDP获取IP地址
func GetOutBoundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Printf("【错误】您的网络不通，请检查！ err: %s", err.Error())
		return "127.0.0.1"
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.String(), ":")[0]
}

// GetLocalIP 获取内网IP地址字符串
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetIPArr 用来获取有效的内网IP
// 使用sclice存储
func GetIPArr() ([]string, error) {
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

// GetIPMap 用来获取有效的内网IP
// 使用map存储
func GetIPMap() (map[int]string, error) {
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
