package multicast

import (
	"bytes"
	"encoding/gob"
	"fmt"
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

	red, _ := net.ResolveUDPAddr("udp", caller)
	log.Println("[ReceiveM]             localhostAddress ", red)

	// printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	listener, err := net.ListenUDP("udp", red)
	f.Error(err, "[ReceiveM] ListenUDP Error")
	defer listener.Close()

	timeoutDuration := 10 * time.Second

	for {

		listener.SetReadDeadline(time.Now().Add(timeoutDuration))
		log.Println("[ReceiveM] ***************************", err)

		log.Println("[ReceiveM] ====================", err)
		buffer := make([]byte, f.MaxBufferSize)
		nRead, src, err := listener.ReadFrom(buffer)
		if err != nil {
			// fmt.Println(err)
			return err
		}

		dataBuffer := bytes.NewBuffer(buffer[:nRead])
		decode := gob.NewDecoder(dataBuffer)
		log.Println("[ReceiveM] buffer: ", nRead)

		err = decode.Decode(&pack)
		f.Error(err, "Receive error  \n")

		log.Println("[ReceiveM] -------RECIDO: ", pack, src)

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
			fmt.Println("[ReceiveM] Me llego un message", packNew)
		case f.Ack:
			chanAc <- packNew
			fmt.Println("[ReceiveM] Me llego un ack", packNew)

		}

	}

	log.Println("[ReceiveM]            Sali por timeout ")

	return err
}
