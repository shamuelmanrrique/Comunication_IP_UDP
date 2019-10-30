package main

import (
	"encoding/gob"
	"flag"
	"fmt"
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

	msm := &f.Message{
		To:   f.MulticastAddress,
		From: connectM.GetId(),
		Targ: connectM.GetId(),
		Data: "inf",
	}

	chanInterf := make(chan interface{})
	ackID := &f.Ack{Code: " id" + "," + msm.GetFrom()}
	vr := "6"
	// red, _ := net.ResolveUDPAddr("udp", connectM.GetId())
	// log.Println("[RM]             localhostAddress ", red)

	// // printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	// listener, err := net.ListenUDP("udp", red)
	// f.Error(err, "[RM] ListenUDP Error")
	// defer listener.Close()

	go u.ReceivePack(chanInterf, connectM.GetPort())
	// go u.ReceiveM(chanInterf, connectM.GetPort())
	time.Sleep(1 * time.Second)
	go u.SendM(ackID, connectM.GetId())

	data := <-chanInterf
	fmt.Println("[MAIN]  --> ", ackID, data, vr, msm)

	// ve, ok := data.(f.Message)
	vEe, ok1 := data.(f.Ack)
	// fmt.Println("[Main] Soy Message", d)
	fmt.Println("[Main] Soy Message", vEe, ok1)

	ve, ok := data.(f.Message)
	// fmt.Println("[Main] Soy Message", d)
	fmt.Println("[Main] Soy Message", ve, ok)
	switch packNew := data.(type) {
	case f.Message:
		fmt.Println("[Main] Soy Message 11", data, t)
	case f.Ack:
		fmt.Println("[Main] Soy ACK", data, t)
	}

	if false {

		// inicio ReceiveGroupM
		go u.ReceiveGroupM(connectM)
		time.Sleep(time.Second * 2)

		// Si soy master llamo SendGroupM msm
		if flags.Master {

			target := ""
			delay, _ := time.ParseDuration("0s")
			inf := "Me mataron"
			id := connectM.GetId()

			// Actualizo el reloj
			vector := connectM.GetVector()

			if len(connectM.GetKill()) > 0 && len(connectM.GetDelays()) > 0 {
				target = connectM.GetTarget(0)
				delay = connectM.GetDelay(0)
				inf = "He disparado"
				connectM.SetKill()
				connectM.SetDelay()
			}

			// Incremento el reloj
			vector.Tick(id)
			connectM.SetClock(vector)

			// TODO CREATE SNAPSHOP RELOJ []VCLOCK
			// Copio el vector
			copyVector := vector.Copy()

			// IMprimo TODO
			// fmt.Println("[Main] ", copyVector, target, delay, inf)

			// En este caso tomo el target para enviar el delay
			var msm f.Message = f.Message{
				To:     f.MulticastAddress,
				From:   id,
				Targ:   target,
				Data:   inf,
				Vector: copyVector,
				Delay:  delay,
			}

			fmt.Println("Llamo sendGroup MAIN", *connectM)
			time.Sleep(time.Second * 2)
			go u.SendGroupM(&msm, connectM)
		}

		for i := 0; i < 10; i = i + 3 {
			time.Sleep(time.Second * 5)
			// fmt.Println("Fin del main, contando...", i, "segundos...", msm)
		}
	}
}
