package functions

import (
	"net"
	"os"
)

func GetIp() string {
	var ip string = ""
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip = addr.String()
		}
	}
	return ip
}
