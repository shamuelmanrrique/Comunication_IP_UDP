package functions

type Connection interface {
	GetId() string
    GetIp() string
    GetHost() string
    GetNum() int
    GetInf() string
}

type Data struct {
	Num  int
	Inf string
	Id, Ip ,Host string
}

func (id Data) GetId() string {
    return id.Id
}
func (ip Data) GetIp() string {
    return ip.Ip
}

func (host Data) GetHost() string {
    return host.Host
}

func (num Data) GetNum() int {    
    return num.Num
}

func (inf Data) GetInf() string {    
    return inf.Inf
}


