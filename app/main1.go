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
	fmt.Println("###################### MAIN 1 ###########################\n")

	delay := []int{3, 3}
	kill := "127.0.0.1:5001"

	// bufferMsm := make(chan f.Message, n)
	bufferMsm := make(chan f.Message)

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
		Kill:  kill,
	}

	go f.ReceiveGroup(connect, n)
	go f.SendGroup(connect, bufferMsm)

	for i := range bufferMsm {
		fmt.Println(i)
	}

	<-time.After(time.Second * 20)

}
