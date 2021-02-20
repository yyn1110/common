package common

import (
	"fmt"
	"testing"
)

func TestGetHostIP(t *testing.T) {
	ips, err := GetHostIP()
	if err != nil {
		t.Errorf("get ip error[%v]", err)
	}

	if len(ips) <= 0 {
		t.Error("get ip failed")
	} else {
		for _, ip := range ips {
			fmt.Printf("interface: %s ipv4: %s\n", ip.Interface, ip.IPv4)
			fmt.Printf("interface: %s ipv6: %s\n", ip.Interface, ip.IPv6)
		}
	}
}
