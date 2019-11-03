package main

import (
	"encoding/gob"
	"flag"
	"log"
	"os"
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

	gob.Register(f.Message{})
	gob.Register(f.Ack{})

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

	log.Println("[MAIN] Fin contando...vector", ids)
	log.Println("[MAIN] Fin contando...vector", vector)

	msmreceive := len(ids) - len(flags.GetTarget()) - 1
	f.DistMsm("UDP " + ip + port)

	connectM := &f.Conn{
		Id:     ip + port,
		Ip:     ip,
		Port:   port,
		Ids:    ids,
		Delays: flags.GetTimeDelay(),
		Kill:   flags.GetTarget(),
		Accept: msmreceive,
		Vector: vector,
	}

	chanAck := make(chan f.Ack, len(connectM.GetIds())-1)
	// chanAck := make(chan f.Ack)
	defer close(chanAck)
	chanMessage := make(chan f.Message, len(connectM.GetIds()))
	// chanMessage := make(chan f.Message)
	defer close(chanMessage)

	go u.ReceiveM(chanAck, chanMessage, connectM.GetPort())

	go u.ReceiveGroupM(chanMessage, chanAck, connectM)
	if flags.GetMaster() {
		go u.SendGroupM(chanAck, connectM)
	}

	for i := 0; i < 35; i = i + 5 {
		time.Sleep(time.Second * 5)
		log.Println("[MAIN] Fin contando...", i, "segundos...")
	}
}
