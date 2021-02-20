package common

import (

	"net"
	"os"
	"strings"
	"fmt"
)

const interfacePrefix = "eth"

type HostIP struct {
	Interface string // like "eth0"
	IPv4      string
	IPv6      string
}

func GetHostIP() ([]HostIP, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Errorf("get host interfaces error[%v]", err)
		return nil, err
	}

	var hostIP HostIP
	hostIPs := make([]HostIP, 0)
	for _, inter := range interfaces {
		if strings.Contains(inter.Name, interfacePrefix) {
			hostIP.Interface = inter.Name
			addrs, err := inter.Addrs()
			if err != nil {
				fmt.Errorf("get addr error[%v]", err)
				return nil, err
			}

			for _, addr := range addrs {
				ip := strings.Split(addr.String(), "/")[0]
				if strings.Contains(ip, "::") {
					// ipv6
					ipAddr, err := net.ResolveIPAddr("ip6", ip)
					if err != nil {
						fmt.Errorf("resovle ipv6[%s] error[%v]", ip, err)
						return nil, err
					}
					hostIP.IPv6 = ipAddr.String()
				} else {
					// ipv4
					ipAddr, err := net.ResolveIPAddr("ip4", ip)
					if err != nil {
						fmt.Errorf("resovle ipv4[%s] error[%v]", ip, err)
						return nil, err
					}
					hostIP.IPv4 = ipAddr.String()
				}
			}
			hostIPs = append(hostIPs, hostIP)
		}
	}

	return hostIPs, nil
}

func GetHostName() string {
	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("get hostname error[%v]", err)
		return "zc-unknown-host"
	}
	return host
}

func GetShortHostName() string {
	hostname := GetHostName()
	return shortHostname(hostname)
}

// shortHostname returns its argument, truncating at the first period.
// For instance, given "www.google.com" it returns "www".
func shortHostname(hostname string) string {
	if i := strings.Index(hostname, "."); i >= 0 {
		return hostname[:i]
	}
	return hostname
}
