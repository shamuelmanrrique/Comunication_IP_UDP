package main

import (
	"fmt"
	f "practice1/functions"
	"time"
)

const (
	n    = 2           // Determinamos el numero de procesos n
	ip   = "127.0.0.1" //En este caso se define local
	port = ":5002"
)

func main() {
	fmt.Println("####### MAIN 1 #####################")

	delay := []int{0, 0, 0, 2}
	bufferMsm := make(chan f.Message, n)

	// Determinamos el numero de procesos n
	var ids []string = f.IdProcess(n, "local")
	// var ids []string = f.IdProcess(n, "remote")

	// fmt.Println(ids[1:])

	var connect f.Conn = f.Conn{
		Id:    ip + port,
		Ip:    ip,
		Port:  port,
		Ids:   ids,
		Delay: delay,
	}

	m := "[MSM] => To: " + connect.GetId() + " From: " + connect.GetEnv(1) + " He disparado a " + connect.GetId()
	var msm f.Message = f.Message{
		To:   connect.GetId(),
		From: connect.GetEnv(1),
		Data: m,
	}

	fmt.Println(connect)

	go f.R(connect, bufferMsm)

	go f.S(connect, msm, bufferMsm)

	<-time.After(time.Second * 20)

}
