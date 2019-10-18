package main

import (
	"fmt"
	f "practice1/functions"
	"time"
)

const (
	n    = 2           // Determinamos el numero de procesos n
	ip   = "127.0.0.1" //En este caso se define local
	port = ":5001"
)

func main() {
	// i :=  conn.RemoteAddr().String()

	i := f.GetIp()
	// Creo canal para comunicar send and receive con un buffer de tamaño n cantidad de procesos
	bufferMsm := make(chan f.Message, n)

	// Determinamos el numero de procesos n
	var ids []string = f.IdProcess(n, "local") // var ids []string = f.IdProcess(n, "remote")
	fmt.Println(ids[1:])
	fmt.Println(bufferMsm)
	fmt.Println(i)

	var connect f.Conn = f.Conn{
		Id:   ip + port,
		Ip:   ip,
		Port: port,
		Ids:  ids,
	}

	var msm f.Message = f.Message{
		To:   connect.GetId(),
		From: connect.GetEnv(1),
		Data: "ja wueno",
	}

	fmt.Println(connect)

	go f.R(connect, bufferMsm)

	go f.S(connect, msm, bufferMsm)

	<-time.After(time.Second * 20)
	//LLAMO A FUNCION QUE VA A TENER SEND RECEIVE

}
