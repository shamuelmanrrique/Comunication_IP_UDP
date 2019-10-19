package functions

type Connection interface {
	GetId() string
	GetIp() string
	GetPort() string
	GetIds() []string
	GetEnv(n int) string
	GetDelay() []int
}

type Conn struct {
	Id, Ip, Port, Host, Env string
	Ids                     []string
	Delay                   []int
}

func (c Conn) GetId() string {
	return c.Id
}

func (c Conn) GetIp() string {
	return c.Ip
}

func (c Conn) GetEnv(n int) string {
	for i, v := range c.GetIds() {
		if i == n {
			return v
		}
	}
	return ""
}

func (c Conn) GetPort() string {
	return c.Port
}

func (c Conn) GetIds() []string {
	return c.Ids
}

func (c Conn) GetDelay() []int {
	return c.Delay
}
