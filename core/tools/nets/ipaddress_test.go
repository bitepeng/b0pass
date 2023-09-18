package nets

import (
	"testing"
)

func TestIpaddress(t *testing.T) {
	ip1 := GetOutBoundIP()
	t.Logf("%v %T", ip1, ip1) //192.168.2.220

	ip2 := GetLocalIP()
	t.Log(ip2) //192.168.56.1

	ip3, _ := GetIPArr()
	t.Log(ip3) //[192.168.56.1 192.168.2.220]

	ip4, _ := GetIPMap()
	t.Log(ip4) //map[0:192.168.56.1 1:192.168.2.220]
}
