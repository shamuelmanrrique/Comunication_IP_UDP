package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	// "net"
	f "practice1/functions"
	v "practice1/vclock"
)

// Estas constantes pasaran como flash en la consola
// go run main.go -r "local" -t "127.0.1.1:5003" -d "10s" -n 3 -m=true -p=":5001"
// go run main.go -r "local" -t "127.0.1.1:5001" -d "20s" -n 3 -p=":5002"
// go run main.go -r "local" -t "127.0.1.1:5002" -d "30s" -n 3 -p=":5003"

// go run main.go -r "local" -t "127.0.1.1:5002" -d "10ms" -n 2 -m=true -p=":5001"
// go run main.go -r "local" -t "127.0.1.1:5001" -d "20ms" -n 2 -p=":5002"

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
		panic("El tamaño del arreglo Targets debe ser igual al de Delay")
		os.Exit(1)
	}

	var ip string = f.IpAddress()
	var port string = flags.GetPort()
	n := flags.GetProcess()
	f.DistMsm(ip + port)

	var ids []string = f.IdProcess(n, flags.GetRun())

	// Inicializo todos el reloj del proceso
	var vector = v.New()
	for _, v := range ids {
		vector[v] = 0
	}

	// connect := New(f.Conn)
	var connect f.Conn = f.Conn{
		Id:     ip + port,
		Ip:     ip,
		Port:   port,
		Ids:    ids,
		Delay:  flags.GetTimeDelay(),
		Kill:   flags.GetTarget(),
		Vector: vector,
	}

	go f.ReceiveGroup(connect)
	if flags.Master {
		fmt.Println("Llamo sendGroup MAIN")
		time.Sleep(time.Second * 1)
		go f.SendGroup(connect)
	}

	<-time.After(time.Second * 40)
}
