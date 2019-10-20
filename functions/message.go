package functions

import (
	v "practice1/vclock"
)

type Msm interface {
	GetTo() string
	GetFrom() string
	GetData() string
	GetVector() v.VClock
}

type Message struct {
	To, From, Data string
	Vector         v.VClock
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

func (m Message) GetVector() v.VClock {
	return m.Vector
}
