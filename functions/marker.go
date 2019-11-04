package functions

type MarkerInterface interface {
	GetCounter()
	GetHeader()
	GetChannel()
}

type Marker struct {
	Counter int
	Header  []Message
	Channel map[string]Message
}

func (m *Marker) GetCounter() int {
	return m.Counter
}

func (m *Marker) GetHeader() []Message {
	return m.Header
}

func (m *Marker) GetChannel() map[string]Message {
	return m.Channel
}

func (m *Marker) SetHeader(msms []Message) {
	m.Header = msms
}

func (m *Marker) SetChannel(val map[string]Message) {
	m.Channel = val
}
