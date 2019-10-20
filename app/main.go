package main

import (
	"fmt"
	// "net"
	f "practice1/functions"
	v "practice1/vclock"
	"time"
)

// EStas constantes pasaran como flash en la consola
const (
	n    = 2           // Determinamos el numero de procesos n
	ip   = "127.0.0.1" //En este caso se define local
	port = ":5001"
)

func main() {
	fmt.Println("###################### MAIN ###########################")

	delay := []int{5, 8}
	kill := "127.0.0.1:5002"
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

	// Proceso maestro llama el send y receive de una vez
	go f.ReceiveGroup(connect, n)
	time.Sleep(time.Second * 2)
	go f.SendGroup(connect)

	<-time.After(time.Second * 20)
}
