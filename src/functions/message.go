package functions

import (
	"log"
	v "sd_paxos/src/vclock"
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
	log.Printf("###################### MAIN  %s ########################### \n", s)
}

func DistWall() {
	log.Println("############################################################################")
}

func DistUnic(s string) {
	log.Printf("#########################  %s ################################# \n", s)
}

func (p *Pack) GetMes() Message {
	return p.Mes
}

func (p *Pack) GetConfACK() Ack {
	return p.ConfACK
}

// CheckMsm function to add only new message in array
func CheckMsm(msms []Message, m Message) ([]Message, bool, Message) {
	for _, a := range msms {
		if m.GetFrom() == a.GetFrom() && m.GetVector().Compare(a.GetVector(), v.Equal) {
			return msms, true, m
		}
	}

	msms = append(msms, m)

	return msms, false, m

}
