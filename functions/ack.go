package functions

import (
	"strings"
)

type AckInterface interface {
	GetCode() string
}

type Ack struct {
	Code string
}

func (a Ack) GetCode() string {
	return a.Code
}

func CheckAcks(acks []Ack, c *Conn) (bool, []string) {
	ips := c.GetIds()
	ips = Remove(ips, c.GetId())

	if len(acks) == len(ips) {
		return true, []string{}
	}

	for _, a := range acks {
		code := strings.Split(a.GetCode(), ",")
		ips = Remove(ips, code[1])
	}

	return false, ips

}

func Remove(l []string, item string) []string {
	var aux []string
	for i, other := range l {
		if other == item {
			aux = append(l[:i], l[i+1:]...)
		}
	}
	return aux
}
