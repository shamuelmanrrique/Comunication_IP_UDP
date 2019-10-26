package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	c "practice1/communication"
	f "practice1/functions"
	v "practice1/vclock"
)

var flags f.Coordinates

func init() {
	flag.IntVar(&flags.Process, "n", 4, "numero de procesos que vas a crear")
	flag.StringVar(&flags.Run, "r", "local", "Se va correr local o remote")
	flag.StringVar(&flags.Port, "p", ":1400", "puerto que usara el proceso :XXXX")
	flag.BoolVar(&flags.Master, "m", false, "pppo")
	flag.Var(&flags.TimeDelay, "d", "Lista de flags separados por coma")
	flag.Var(&flags.Target, "t", "listas de ip objectivos")
}

func main() {
	flag.Parse()
	var val bool = len(flags.TimeDelay) != len(flags.Target)
	if val {
		panic("El tamaño del arreglo Targets debe ser igual al de Delays")
		os.Exit(1)
	}

	ip := f.IpAddress()
	port := flags.GetPort()
	n := flags.GetProcess()
	f.DistMsm(ip + port)

	var ids []string = f.IdProcess(n, flags.GetRun())

	// Inicializo todos el reloj del proceso
	var vector = v.New()
	for _, v := range ids {
		vector[v] = 0
	}

	msmreceive := len(ids) - len(flags.GetTarget()) - 1
	fmt.Println("ESTOY EN EL MAIN n ", msmreceive)

	connect := &f.Conn{
		Id:     ip + port,
		Ip:     ip,
		Port:   port,
		Ids:    ids,
		Delays: flags.GetTimeDelay(),
		Kill:   flags.GetTarget(),
		Accept: msmreceive,
		Vector: vector,
	}

	go c.ReceiveGroup(connect)
	if flags.Master {
		fmt.Println("Llamo sendGroup MAIN", *connect)
		time.Sleep(time.Second * 1)
		go c.SendGroup(connect)
	}

	for i := 0; i < 30; i = i + 5 {
		time.Sleep(time.Second * 5)
		fmt.Println("Fin del main, contando...", i, "segundos...")
	}
}
