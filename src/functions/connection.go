package functions

import (
	"net"
	v "sd_paxos/src/vclock"
	"time"
)

// Inf Remote IPs
var RemoteIP3 = map[string]string{"TestSubNet0": "155.210.154.199:1400", "TestSubNet1": "155.210.154.200:1400", "TestSubNet2": "155.210.154.204:1400"}
var RemoteIP3s = []string{"155.210.154.199:1400", "155.210.154.200:1400", "155.210.154.204:1400"}
var RemoteIP5 = map[string]string{"TestSubNetR50": "155.210.154.199:1400", "TestSubNetR51": "155.210.154.200:1400", "TestSubNetR52": "155.210.154.204:1400", "TestSubNetR53": "155.210.154.209:1400", "TestSubNetR54": "155.210.154.210:1400"}
var RemoteIP5T = map[string]string{"TestSubNetRD50": "155.210.154.199:1400", "TestSubNetRD51": "155.210.154.200:1400", "TestSubNetRD52": "155.210.154.204:1400", "TestSubNetRD53": "155.210.154.209:1400", "TestSubNetRD54": "155.210.154.210:1400"}
var RemoteIP5s = []string{"155.210.154.199:1400", "155.210.154.200:1400", "155.210.154.204:1400", "155.210.154.204:1400", "155.210.154.210:1400"}
var GoMainLog = "/usr/local/go/bin/go run /home/a802400/go/src/sd_petry_nets/src/cmd/distconssim/main.go"
var GoMain = "/usr/local/go/bin/go run /home/a802400/go/src/sd_petry_nets/src/cmd/distconssim/main.go -l=false"
var GoTest = "/usr/local/go/bin/go test -timeout 499s -v /home/a802400/go/src/sd_petry_nets/src/distconssim -run "

// Inf Local IPs
var LocalIP3 = map[string]string{"TestSubNetL0": "127.0.1.1:5000", "TestSubNetL1": "127.0.1.1:5001", "TestSubNetL2": "127.0.1.1:5002"}
var LocalIP3s = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002"}
var LocalIP5 = map[string]string{"TestSubNetL0": "127.0.1.1:5000", "TestSubNetL1": "127.0.1.1:5001", "TestSubNetL2": "127.0.1.1:5002", "TestSubNetL3": "127.0.1.1:5003", "TestSubNetL4": "127.0.1.1:5004"}
var LocalIP5s = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003", "127.0.1.1:5004"}
var GoLocalMainLog = "go run ~/go/src/sd_petry_nets/src/cmd/distconssim/main.go"
var GoLocalTest = "go test ~/go/src/sd_petry_nets/src/distconssim -timeout 499s -v -run "

// type Connections []Conn

const (
	MulticastAddress = "229.0.40.000:9999"
	MaxBufferSize    = 8192
)

type Connection interface {
	GetId() string
	GetIp() string
	GetPort() string
	GetKill() []string
	GetIds() []string
	GetDelays() []time.Duration
	GetDelay(n int) time.Duration
	GetTarget(n int) string //puedo elminar
	GetEnv(n int) string
	GetVector() v.VClock
	GetListe() net.Listener
	GetAccept() int
}

type Conn struct {
	Id, Ip, Port, Host, Env string
	Vector                  v.VClock
	Ids                     []string
	Kill                    []string
	Delays                  []time.Duration
	Liste                   net.Listener
	Accept                  int
}

func (c *Conn) SetKill() {
	n := len(c.Kill)
	if n > 0 {
		c.Kill = c.Kill[:n-1]
	}
}

func (c *Conn) SetDelay() {
	n := len(c.Kill)
	if n > 0 {
		c.Delays = c.Delays[:n-1]
	}
}

func (c *Conn) SetClock(v v.VClock) {
	c.Vector = v
}

func (c Conn) GetId() string {
	return c.Id
}

func (c Conn) GetIp() string {
	return c.Ip
}

func (c Conn) GetEnv(n int) string {
	for i, v := range c.GetIds() {
		if i == n {
			return v
		}
	}
	return ""
}

func (c Conn) GetDelay(n int) time.Duration {
	for i, v := range c.GetDelays() {
		if i == n {
			return v
		}
	}
	return 0
}

func (c Conn) GetPort() string {
	return c.Port
}

func (c Conn) GetAccept() int {
	return c.Accept
}

func (c Conn) GetIds() []string {
	return c.Ids
}

func (c Conn) GetKill() []string {
	return c.Kill
}

func (c Conn) GetDelays() []time.Duration {
	return c.Delays
}

func (c Conn) GetVector() v.VClock {
	return c.Vector
}

func (c Conn) GetTarget(n int) string {
	for i, v := range c.GetKill() {
		if i == n {
			return v
		}
	}
	return ""
}

func (c Conn) GetListe() net.Listener {
	return c.Liste
}

// // NewConnec will create slice of Connect
// func NewConnec(IPs []string) Connections {
// 	var connections Connections
// 	for _, val := range IPs {
// 		addr := strings.Split(val, ":")
// 		conn := Connect{}
// 		conn.IP = addr[0]
// 		conn.Port = addr[1]
// 		conn.Accept = false
// 		conn.IDSubRed = val
// 		connections = append(connections, conn)
// 	}
// 	return connections
// }

// // GetConnection return connection by Index in slices
// func (c Connections) GetConnection(n int) Connect {
// 	for i, connect := range c {
// 		if i == n {
// 			return connect
// 		}
// 	}
// 	return Connect{}
// }
