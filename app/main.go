package main

import (
	"flag"
	"fmt"

	// "net"

	f "practice1/functions"
	v "practice1/vclock"
	"time"
)

// Estas constantes pasaran como flash en la consola
const (
	n    = 2           // Determinamos el numero de procesos n
	ip   = "127.0.0.1" //En este caso se define local
	port = ":5001"
)

// var DelaysFlag Delays

// func init() {
// 	// Tie the command-line flag to the DelaysFlag variable and
// 	// set a usage message.
// 	flag.Var(&DelaysFlag, "deltaT", "comma-separated list of Delayss to use between events")
// }
func main() {
	var DelaysFlag f.Delays
	flag.Var(&DelaysFlag, "deltaT", "comma-separated list of Delayss to use between events")

	// var intervalFlag interval
	//Estoy coloacando cuatro procesos por defoult
	// var process = flag.Int("process", 4, "numero de procesos que quieres crear")

	// t := make([]time.Duration, process)

	// var intervalFlag interval
	// p := flag.Var(&intervalFlag, "p", "comma-separated list of intervals to use between events")
	flag.Parse()

	//var delay = flag.Duration("delay",, "arreglo de retardos")
	//var procesos = flag.Int("procesos", 4, "numero de procesos que quieres crear")
	fmt.Println(DelaysFlag[1:])
	// fmt.Println("count value ", *process)
	fmt.Printf("###################### MAIN  %s ########################### \n", ip+port)
	d := []int{5, 8}
	kill := "127.0.0.1:5																																																																																	002"
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
		Delay:  d,
		Kill:   kill,
		Vector: vector,
	}

	// Proceso maestro llama el send y receive de una vez
	//go f.ReceiveGroup(connect, n)
	time.Sleep(time.Second * 2)
	go f.SendGroup(connect)

	//<-time.After(time.Second * 20)
}
