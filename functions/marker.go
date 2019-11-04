package functions

import (
	"fmt"
)

var BufferMsm = make(map[string]Message)

type MarkerInterface interface {
	GetCounter()
	GetHeader()
	GetChannel()
	GetRecoder()
	GetCheckPoints()
	SetRecoder(b bool)
	SetCounter(val int)
	SetHeader(msms []Message)
	SetChannel(val []Message)
}

type Marker struct {
	Recoder     bool
	Counter     int
	Header      []Message
	Channel     []Message
	CheckPoints []string
}

func (m *Marker) GetRecoder() bool {
	return m.Recoder
}

func (m *Marker) GetCounter() int {
	return m.Counter
}

func (m *Marker) GetCheckPoints() []string {
	return m.CheckPoints
}

func (m *Marker) GetHeader() []Message {
	return m.Header
}

func (m *Marker) GetChannel() []Message {
	return m.Channel
}

func (m *Marker) SetCounter(val int) {
	m.Counter = val
}

func (m *Marker) SetHeader(msms []Message) {
	m.Header = msms
}

func (m *Marker) SetCheckPoints(ch string) {
	m.CheckPoints = append(m.GetCheckPoints(), ch)
}

func (m *Marker) SetChannel(val []Message) {
	for _, c := range val {
		m.Channel = append(m.Channel, c)
	}
}

func (m *Marker) SetRecoder(b bool) {
	m.Recoder = b
}

func (m *Marker) PrintMarker(ip string) {
	fmt.Println("################### SNAPSHOT", ip, "#####################")
	fmt.Println("Init state:", m.GetHeader())
	fmt.Println("Recording:", m.GetHeader())
	fmt.Println("CheckPoints:", m.GetCheckPoints())
	fmt.Println("#################################################################")
}
