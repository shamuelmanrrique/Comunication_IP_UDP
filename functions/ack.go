package functions

type AckInterface interface {
	GetCode() string
}

type Ack struct {
	Code string
}

func (a Ack) GetCode() string {
	return a.Code
}
