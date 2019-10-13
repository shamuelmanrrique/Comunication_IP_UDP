package functions

import ( 
    // "fmt"
    // "strings"
    "strconv"
)

var Info []string

type p struct {
    a Data
}

type Data struct {
	num  int
	Inf []string
	id, ip ,host string
	// el msm debe ser una variable local
}

type Conection interface {
	GetId() string
    GetIp() string
    GetHost() string
    GetNum() int
    GetInfo() []string
}

func (id Data) getId() string {
    return id.id
}
func (ip Data) getIp() string {
    return ip.ip
}

func (host Data) getHost() string {
    return host.host
}

func (num Data) getNumber() int {    
    return num.num
}

func (inf Data) getInfo() []string {    
    return inf.Inf
}

func (d Data) Getdata() { 
    aux := d.Inf
    // primes := [1]string{"2:127.0.0.1:1402"}
    // reg := ["2:127.0.0.1:1402"]
    // a := strings.Join(d.data[:], " ")
    // d = strings.Split(a, ":")
    d.id =aux[1] + aux[2]
    d.ip =aux[1]
    d.host =aux[2]
    n, err := strconv.Atoi(aux[0])
    Error(err)
    d.num = n
    
}

