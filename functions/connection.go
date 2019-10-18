package functions

type Connection interface {
	GetId() string
    GetIp() string
    GetPort() string
    GetEnv(n int) string
    GetIds() []string

}

type Conn struct {
    Id, Ip , Port, Host, Env string
    Ids []string
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

func (c  Conn) GetPort() string {
    return c.Port
}

func (c Conn) GetIds() []string {    
    return c.Ids
}

