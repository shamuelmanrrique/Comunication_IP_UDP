package functions

import (
	v "practice1/vclock"
)

type Connection interface {
	GetId() string
	GetIp() string
	GetPort() string
	GetKill() string //puedo elminar
	GetIds() []string
	GetDelays() []int
	GetDelay(n int) int
	GetEnv(n int) string
	// GetValues(ip string) (string, string)
	GetVector() v.VClock
}

type Conn struct {
	Id, Ip, Port, Host, Env, Kill string
	Vector                        v.VClock
	Ids                           []string
	Delay                         []int
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

func (c Conn) GetDelay(n int) int {
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

func (c Conn) GetIds() []string {
	return c.Ids
}

func (c Conn) GetKill() string {
	return c.Kill
}

func (c Conn) GetDelays() []int {
	return c.Delay
}

func (c Conn) GetVector() v.VClock {
	return c.Vector
}

// func (c Conn) GetValues(ip string) (string,string) {
// 	for i, v := range c.GetIds() {
// 		s := strings.Split(v, ":")
// 		if ip == s[:1] {
// 			return
// 		}
// 	}
// 	return 0
// }
