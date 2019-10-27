package main

import (
	"flag"
	"fmt"
	"time"

	f "practice1/functions"
	u "practice1/multicast"
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
	// Comentados para pruebas con UDP
	// var val bool = len(flags.TimeDelay) != len(flags.Target)
	// if val {
	// 	panic("El tamaño del arreglo Targets debe ser igual al de Delays")
	// 	os.Exit(1)
	// }

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
	fmt.Println("ESTOY EN EL MAIN port: ", port, " ip : ", ip)

	// connect := &f.Conn{
	// 	Id:     ip + port,
	// 	Ip:     ip,
	// 	Port:   port,
	// 	Ids:    ids,
	// 	Delays: flags.GetTimeDelay(),
	// 	Kill:   flags.GetTarget(),
	// 	Accept: msmreceive,
	// 	Vector: vector,
	// }

	// go c.ReceiveGroup(connect)
	// if flags.Master {
	// 	fmt.Println("Llamo sendGroup MAIN", *connect)
	// 	time.Sleep(time.Second * 1)
	// 	go c.SendGroup(connect)
	// }

	println("##################   UDP ", ip, "   ###############################")

	connect := &f.Conn{
		Id:     "229.0.040.000:9999",
		Ip:     "229.0.040.000",
		Port:   ":9999",
		Ids:    ids,
		Delays: flags.GetTimeDelay(),
		Kill:   flags.GetTarget(),
		Accept: msmreceive,
		Vector: vector,
	}

	var c chan f.Message

	msm := &f.Message{
		To:   connect.GetId(),
		From: "4253647586970",
		Data: "Hola",
	}

	// log.Println(*connect)
	// go u.ReceiveMulticast(c, connect)
	// time.Sleep(time.Second * 3)
	// go u.SendMulticast(msm, connect)

	go u.ReceiveGroupM(c, connect)
	// time.Sleep(time.Second * 2)
	// go u.SendGroupM(msm, connect)

	// go u.SendGroupM(msm, connect)

	// go u.SendGroupM(msm, connect)

	for i := 0; i < 50000; i = i + 1 {
		time.Sleep(time.Second * 5)
		// fmt.Println("Fin del main, contando...", i, "segundos...", msm)
	}

}
