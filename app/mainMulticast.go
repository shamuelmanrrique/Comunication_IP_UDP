package main

import (
	"fmt"
	"net"
	"time"
)

const (
	defaultMulticastAddress = "229.0.040.000:9999"
)

func mainMul() {
	// ######################################################
	// ################### MULTICAST	#####################
	// ######################################################
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

	// Define connection to udp
	redUDP, err := net.ResolveUDPAddr("udp", connectM.GetId())
	f.Error(err, "Send connection error \n")

	connectionUDP, _ := net.DialUDP("udp", nil, redUDP)
	f.Error(err, "Send connection error \n")
	defer connectionUDP.Close()

	msm := &f.Message{
		To:   f.MulticastAddress,
		From: connectM.GetId(),
		Targ: connectM.GetId(),
		Data: "inf",
	}
	ackID := &f.Ack{Code: connectM.GetId() + "," + msm.GetFrom()}

	pack := &f.Pack{
		Mes:     *msm,
		ConfACK: *ackID,
	}

	var canalPacks chan f.Pack

	go u.ReceivePack(canalPacks, connectionUDP, connectM.GetId())
	time.Sleep(time.Second * 3)
	go u.SendPack(pack, connectM.GetId())

	recpack, _ := <-canalPacks

	fmt.Println("[Main] envio el msm que recibi: ", recpack)

	// // inicio ReceiveGroupM
	// go u.ReceiveGroupM(connectM)
	// time.Sleep(time.Second * 2)

	// // Si soy master llamo SendGroupM msm
	// if flags.Master {

	// 	target := ""
	// 	delay, _ := time.ParseDuration("0s")
	// 	inf := "Me mataron"
	// 	id := connectM.GetId()

	// 	// Actualizo el reloj
	// 	vector := connectM.GetVector()

	// 	if len(connectM.GetKill()) > 0 && len(connectM.GetDelays()) > 0 {
	// 		target = connectM.GetTarget(0)
	// 		delay = connectM.GetDelay(0)
	// 		inf = "He disparado"
	// 		connectM.SetKill()
	// 		connectM.SetDelay()
	// 	}

	// 	// Incremento el reloj
	// 	vector.Tick(id)
	// 	connectM.SetClock(vector)

	// 	// TODO CREATE SNAPSHOP RELOJ []VCLOCK
	// 	// Copio el vector
	// 	copyVector := vector.Copy()

	// 	// IMprimo TODO
	// 	// fmt.Println("[Main] ", copyVector, target, delay, inf)

	// 	// En este caso tomo el target para enviar el delay
	// 	var msm f.Message = f.Message{
	// 		To:     f.MulticastAddress,
	// 		From:   id,
	// 		Targ:   target,
	// 		Data:   inf,
	// 		Vector: copyVector,
	// 		Delay:  delay,
	// 	}

	// 	fmt.Println("Llamo sendGroup MAIN", *connectM)
	// 	time.Sleep(time.Second * 2)
	// 	go u.SendGroupM(&msm, connectM)
	// }

	for i := 0; i < 10; i = i + 1 {
		time.Sleep(time.Second * 5)
		// fmt.Println("Fin del main, contando...", i, "segundos...", msm)
	}

}
