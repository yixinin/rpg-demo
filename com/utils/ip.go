package utils

import (
	"net"
	"strings"
)

func IPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip := ipnet.IP; ip != nil && ip.IsGlobalUnicast() && strings.Contains(ipnet.String(), "/24") {
				return ip.To4().String()
			}
		}
	}
	return ""
}
