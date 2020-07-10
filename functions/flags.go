package functions

type Coordinates struct {
	Machine string
	Mode    string
}

type CoordinatesInterface interface {
	GetMachine()
	GetMode()
}

func (c Coordinates) GetMachine() string {
	return c.Machine
}

func (c Coordinates) GetMode() string {
	return c.Mode
}
