package functions

import (
	"net"
	v "practice1/vclock"
	"time"
)

type Connection interface {
	GetId() string
	GetIp() string
	GetPort() string
	GetKill() []string //puedo elminar
	GetIds() []string
	GetDelays() Delays
	GetDelay(n int) time.Duration
	GetTarget(n int) string //puedo elminar
	GetEnv(n int) string
	GetVector() v.VClock
	GetListe() net.Listener
	Conn([]string)
	// SetKill(s []string) //puedo elminar
}

type Conn struct {
	Id, Ip, Port, Host, Env string
	Vector                  v.VClock
	Ids, Kill               []string
	Delay                   Delays
	Liste                   net.Listener
}

func (c *Conn) SetKill(s []string) {
	c.Kill = s
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

func (c Conn) GetIds() []string {
	return c.Ids
}

func (c Conn) GetKill() []string {
	return c.Kill
}


func (c Conn) GetDelays() Delays {
	return c.Delay
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
