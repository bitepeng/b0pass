package ipaddress

import (
	"testing"
)

func TestGetIP(t *testing.T) {
	ip, err := GetIP()
	if err != nil {
		t.Error(err)
	}
	t.Log(ip)
}

func TestGetIP2(t *testing.T) {
	ip, err := GetIP2()
	if err != nil {
		t.Error(err)
	}
	t.Log(ip)
}
