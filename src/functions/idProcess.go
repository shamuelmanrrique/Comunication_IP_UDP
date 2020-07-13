package functions

import (
	"net"
	"os"
	"strconv"
)

var RemoteIPs = []string{"155.210.154.199", "155.210.154.197", "155.210.154.196"}
var RemoteFlags = []string{"-r=\"proof\" -i=\"155.210.154.199\" -t=\"155.210.154.197\" -d=\"5s\" -n=3 -m=true -p=\":1400\"", "-i=\"155.210.154.197\" -r=\"proof\" -n=3 -p=\":1400\"", "-i=\"155.210.154.196\" -r=\"proof\" -n=3 -p=\":1400\""}
var LocalFlags = []string{"-r=\"local\" -t=\"127.0.1.1:5002\" -d=\"5s\" -n=3 -m=true -p=\":5001\"", "-r=\"local\" -n=3 -p=\":5002\"", " -r=\"local\" -n=3 -p=\":5003\""}
var Command = make(map[string]string)

// -ip="155.210.154.199", -ip="155.210.154.197", -ip="155.210.154.196"

// type  Command make(map[string]string)
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
			id = "155.210.154." + strconv.Itoa(190+i) + ":1400"
			ids = append(ids, id)
		}

	} else if mode == "proof" {
		// Ip con las que voy hacer la prueba
		ids = RemoteIPs
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

func NewCommand(ips []string, name string) map[string]string {
	if name == "proof" {
		for i, ip := range ips {
			Command[ip] = GetString(i, RemoteFlags)
		}
	} else if "local" == name {
		for i, ip := range ips {
			Command[ip] = GetString(i, LocalFlags)
		}
	}

	return Command
}

func GetString(n int, value []string) string {
	for i, v := range value {
		if i == n {
			return v
		}
	}
	return ""
}

func FlagsExec(p map[string]string, ip string) string {
	if v, found := p[ip]; found {
		return v
	}

	return ""
}
