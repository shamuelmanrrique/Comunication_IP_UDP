package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
	"time"
)

// Receive TODO ELIMINAR CALLER
func ReceiveM(chanAc chan<- f.Ack, chanMes chan<- f.Message, caller string) error {
	// func ReceiveM(canal chan<- interface{}, caller string) error {
	var pack interface{}
	var err error

	log.Println("[ReceiveM]       Estoy escuchando UDP ", caller)
	red, _ := net.ResolveUDPAddr("udp", caller)

	listener, err := net.ListenUDP("udp", red)
	f.Error(err, "[ReceiveM] ListenUDP Error")
	defer listener.Close()

	timeoutDuration := 10 * time.Second

	for {
		// log.Println("[ReceiveM]+++++++++++++++FOR++++++++")
		listener.SetReadDeadline(time.Now().Add(timeoutDuration))

		buffer := make([]byte, f.MaxBufferSize)
		nRead, src, err := listener.ReadFrom(buffer)
		if err != nil {
			return err
		}

		dataBuffer := bytes.NewBuffer(buffer[:nRead])
		decode := gob.NewDecoder(dataBuffer)
		// log.Println("[ReceiveM] buffer: ", nRead)

		err = decode.Decode(&pack)
		f.Error(err, "Receive error  \n")

		// ack, oka := pack.(f.Ack)
		// if oka {
		// 	fmt.Println("[ReceiveM] Me llego un ack", ack)
		// 	chanAc <- ack
		// }

		// msm, okm := pack.(f.Message)
		// if okm {
		// 	fmt.Println("[ReceiveM] Me llego un message", msm)
		// 	chanMes <- msm
		// }

		// fmt.Println("[Main] Soy Message", d)
		// fmt.Println("[Main] Soy Message", ve, ok)

		switch packNew := pack.(type) {
		case f.Message:
			chanMes <- packNew
			log.Println("[ReceiveM] ===> MESSAGE ", packNew, " DE ", src)
		case f.Ack:
			chanAc <- packNew
			log.Println("[ReceiveM] ----> ACK ", packNew, " DE ", src)
		}

	}

	log.Println("[ReceiveM] |||||| FOR ReceiveM timeout |||| ")

	return err
}
