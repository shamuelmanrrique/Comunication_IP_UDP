package functions

import (
	"net"
	"os"
	"fmt"
)

func GetIp() string {
	var ip string = ""
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			// fmt.Println("IPv4: ", ipv4)
			ip =  addr.String() 
		}
	}
	return ip
}
