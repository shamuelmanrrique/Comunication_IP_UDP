package functions

type Connection interface {
	GetId() string
    GetIp() string
    GetPort() string
    GetIds() []string
}

type Conn struct {
    Id, Ip , Port, Host string
    Ids []string
}

func (id Conn) GetId() string {
    return id.Id
}
func (ip Conn) GetIp() string {
    return ip.Ip
}

func (port  Conn) GetPort() string {
    return port.Port
}

func (ids Conn) GetIds() []string {    
    return ids.Ids
}


type Message struct{
    
}