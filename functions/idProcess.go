package functions

import (
	"net"
	"os"
	"strconv"
)

func IdProcess(n int, mode string) []string {
	var id string
	var ids []string

	if mode == "local" {
		for i := 1; i <= n; i++ {
			id = "127.0.1.1:500" + strconv.Itoa(i)
			ids = append(ids, id)
		}

	} else if mode == "remote" {
		for i := 1; i < 21; i++ {
			// id = "144.210.154." + strconv.Itoa(190 + i) + ":1400"
			id = "155.210.154." + strconv.Itoa(190+i) + ":1400"
			ids = append(ids, id)
		}

	}

	return ids
}

func IpAddress() string {
	var ip string
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip = ipv4.String()
		}
	}
	return ip
}
