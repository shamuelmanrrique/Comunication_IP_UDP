package functions

import (
	"fmt"
	v "practice1/vclock"
)

type Msm interface {
	GetTo() string
	GetFrom() string
	GetData() string
	GetIgnor() string
	GetVector() v.VClock
}

type Message struct {
	To, From, Data, Ignor string
	Vector                v.VClock
}

func (m Message) GetTo() string {
	return m.To
}

func (m Message) GetFrom() string {
	return m.From
}

func (m Message) GetData() string {
	return m.Data
}

func (m Message) GetIgnor() string {
	return m.Ignor
}

func (m Message) GetVector() v.VClock {
	return m.Vector
}

func DistMsm(s string) {
	fmt.Printf("###################### MAIN  %s ########################### \n", s)
}
