package main

import (
	"fmt"
	f "practice1/functions"
	"time"
)

// EStas constantes pasaran como flash en la consola
const (
	n    = 2           // Determinamos el numero de procesos n
	ip   = "127.0.0.1" //En este caso se define local
	port = ":5001"
)

func main() {
	fmt.Println("###################### MAIN ###########################\n")

	delay := []int{3, 3}
	kill := "127.0.0.1:5003"

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

	go f.SendGroup(connect, bufferMsm)
	// go f.receiveGroup(connect, bufferMsm)

	<-time.After(time.Second * 20)
}
