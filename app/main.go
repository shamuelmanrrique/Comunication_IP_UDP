package main

import (
	"flag"
	"fmt"
	"net"

	// "net"

	f "practice1/functions"
	v "practice1/vclock"
)

// Estas constantes pasaran como flash en la consola
// go run main.go -r "remote" -t "12.3.2.3,3223.323" -d "10s,20s,20s" -n 5 -m true
// go run main.go -r "local" -t "12.3.2.3,3223.323" -d "10s,20s,20s" -n 5 -m true

const ( //En este caso se define local
	port = ":1400"
)

var flags f.Coordinates

func init() {
	flag.IntVar(&flags.Process, "n", 4, "pppo")
	flag.StringVar(&flags.Run, "r", "local", "pppo")
	flag.BoolVar(&flags.Master, "m", false, "pppo")
	flag.Var(&flags.TimeDelay, "d", "Lista de flags separados por coma")
	flag.Var(&flags.Target, "t", "listas de ip objectivos")
}

func main() {
	var ip string = f.IpAddress()
	var err error
	netInterfaceAddresses, err := net.InterfaceAddrs()
	f.Error(err, "MAIN ERROR")
	flag.Parse()

	fmt.Println(flags)
	fmt.Println(netInterfaceAddresses)

	fmt.Printf("###################### MAIN  %s ########################### \n", ip)
	var ids []string = f.IdProcess(flags.Process, flags.Run)

	fmt.Println(ids)
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
		Delay:  flags.TimeDelay,
		Kill:   flags.Target,
		Vector: vector,
	}

	// Proceso maestro llama el send y receive de una vez
	// go f.ReceiveGroup(connect, n)
	// time.Sleep(time.Second * 2)
	// go f.SendGroup(connect)
	//<-time.After(time.Second * 20)
}
