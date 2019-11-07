package functions

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

func AddAcks(acks []Ack, a Ack) ([]Ack, bool) {
	for _, ac := range acks {
		if a == ac {
			return acks, true
		}
	}
	acks = append(acks, a)
	return acks, false

}

// La funcion toma el Ack y chequea
func CheckAcks(waitAks []string, acks []Ack) ([]string, bool) {
	aux := waitAks
	for _, a := range acks {
		for _, ip := range waitAks {
			if a.GetOrigen() == ip {
				aux = Remove(aux, ip)
			}
		}
	}

	if len(waitAks) == 0 {
		return aux, true
	}
	return aux, false
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
