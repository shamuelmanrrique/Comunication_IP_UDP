package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"time"

	c "practice1/chandylamport"
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

	gob.Register(f.Message{})
	gob.Register(f.Marker{})
	// Comentados para pruebas con UDP
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

	chanMarker := make(chan f.Marker, n)
	defer close(chanMarker)
	chanMessage := make(chan f.Message, n)
	defer close(chanMessage)

	fmt.Println("ESTOY FUERA")
	go c.ReceiveGroup(chanMessage, chanMarker, connect)
	if flags.Master {
		// fmt.Println("Llamo sendGroup MAIN", *connect)
		time.Sleep(time.Second * 1)
		go c.SendGroup(chanMessage, chanMarker, connect)
	}

	// target := ""
	// inf := "am dead"
	// id := connect.GetId()

	// msm := &f.Message{
	// 	To:   id,
	// 	From: id,
	// 	Targ: target,
	// 	Data: inf,
	// }

	// marker := &f.Marker{
	// 	Counter: 5,
	// }

	// go c.Receive(chanMarker, chanMessage, connect.GetPort())
	// <-time.After(time.Second * 1)
	// go c.Send(msm, connect.GetId())
	// go c.Send(marker, connect.GetId())

	// x1 := <-chanMessage
	// x2 := <-chanMarker

	// fmt.Println(x1, x2)
	// fmt.Println(x1)

	<-time.After(time.Second * 30)

}
