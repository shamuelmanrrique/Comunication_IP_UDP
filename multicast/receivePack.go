package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
)

// Receive TODO ELIMINAR CALLER
func ReceivePack(canal chan<- f.Pack, listener *net.UDPConn, caller string) error {
	var pack = f.Pack{}
	var err error

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
