package main

import (
	"fmt"
	f "practice1/functions"
	v "practice1/vclock"
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
	// Determinamos el numero de procesos n
	var ids []string = f.IdProcess(n, "local")

	// Inicializo todos el reloj del proceso
	var vector = v.New()
	for _, v := range ids {
		vector[v] = 0
	}

	var connect f.Conn = f.Conn{
		Id:     ip + port,
		Ip:     ip,
		Port:   port,
		Ids:    ids,
		Delay:  delay,
		Kill:   kill,
		Vector: vector,
	}

	go f.ReceiveGroup(connect, n)
	// time.Sleep(time.Second * 5)
	// go f.SendGroup(connect)

	<-time.After(time.Second * 30)

}
