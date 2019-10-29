package functions

import (
	"fmt"
	v "practice1/vclock"
	"time"
)

type PackInt interface {
	GetMes() Message
	GetConfACK() Ack
}

type Pack struct {
	Mes     Message
	ConfACK Ack
}

type Msm interface {
	GetTo() string
	GetFrom() string
	GetData() string
	GetTarg() string
	GetDelay() time.Duration
	GetVector() v.VClock
}

type Message struct {
	To, From, Data, Targ string
	Delay                time.Duration
	Vector               v.VClock
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

func (m Message) GetTarg() string {
	return m.Targ
}

func (m Message) GetVector() v.VClock {
	return m.Vector
}

func (m *Message) SetDelay(t time.Duration) {
	m.Delay = t
}

func (m *Message) GetDelay() time.Duration {
	return m.Delay
}

func DistMsm(s string) {
	fmt.Printf("###################### MAIN  %s ########################### \n", s)
}

func (p *Pack) GetMes() Message {
	return p.Mes
}

func (p *Pack) GetConfACK() Ack {
	return p.ConfACK
}
