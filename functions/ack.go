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

// La funcion toma el Ack y chequea
func CheckAcks(acks []Ack, c *Conn) ([]string, bool) {
	ips := c.GetIds()
	ips = Remove(ips, c.GetId())

	if len(acks) == len(ips) {
		return []string{}, true
	}

	for _, a := range acks {
		code := strings.Split(a.GetCode(), ",")
		ips = Remove(ips, code[0])
	}

	return ips, false

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
