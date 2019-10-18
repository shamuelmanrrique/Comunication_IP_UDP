package functions

type Msm interface {
	GetTo() string
	GetFrom() string
    GetData() string
    GetVector() []int
}

type Message struct {
	To, From, Data string
	Vector []int
}

func (m Message) GetTo() string {
    return m.From
}

func (m Message) GetFrom() string {
    return m.To
}

func (m  Message) GetData() string {
    return m.Data
}

func (m Message) GetVector() []int {    
    return m.Vector
}
