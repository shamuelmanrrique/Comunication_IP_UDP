package functions

import (
	"log"
	"strings"
)

type AckInterface interface {
	GetCode() string
	GetOrigen() string
}

type Ack struct {
	Origen string
	Code   string
}

func (a Ack) GetCode() string {
	return a.Code
}

func (a Ack) GetOrigen() string {
	return a.Origen
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

func AddAcks(acks []Ack, a Ack) ([]Ack, bool) {
	for _, ac := range acks {
		if a == ac {
			log.Println(" [SendGroupM] = [CheckMsm] Ya lento ese ACK")
			return acks, true
		}
	}

	log.Println("[SendGroupM] = [CheckMsm] Agrego el ACK ")
	acks = append(acks, a)

	return acks, false

}
