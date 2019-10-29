package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
)

// Receive TODO ELIMINAR CALLER
func ReceivePack(canal chan<- interface{}, caller string) error {
	var pack interface{}
	var err error

	red, _ := net.ResolveUDPAddr("udp", caller)
	log.Println("[RM]             localhostAddress ", red)

	// printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	listener, err := net.ListenUDP("udp", red)
	f.Error(err, "[RM] ListenUDP Error")
	defer listener.Close()

	buffer := make([]byte, f.MaxBufferSize)
	log.Println("[RM] buffer: ")
	nRead, src, err := listener.ReadFrom(buffer)
	f.Error(err, "Error en RM")
	// log.Println("[RM] listener: ")
	dataBuffer := bytes.NewBuffer(buffer[:nRead])
	log.Println("[RM] databuffer: ")
	decode := gob.NewDecoder(dataBuffer)

	err = decode.Decode(&pack)
	f.Error(err, "Receive error  \n")

	log.Println("[RM] -------RECIDO: ", pack, src)
	canal <- pack

	return err
}
