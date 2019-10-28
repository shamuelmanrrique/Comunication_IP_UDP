package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
	"time"
)

func ReceiveM(canalAck chan f.Ack, canal chan f.Message, addr string) error {
	var ac f.Ack
	var err error
	var listener *net.UDPConn
	const maxBufferSize = 4046

	log.Println("[RM]             ReceiveM ")

	//Duda solo debo pasar el puerto o el ip completo por ser UDP
	red, _ := net.ResolveUDPAddr("udp", addr)
	log.Println("[RM]             localhostAddress ", red)

	// printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	listener, err = net.ListenUDP("udp", red)
	f.Error(err, "[RM] ListenUDP Error")
	defer listener.Close()

	buffer := make([]byte, maxBufferSize)
	log.Println("[RM] buffer: ")
	nRead, src, err := listener.ReadFrom(buffer)
	f.Error(err, "Error en RM")
	// log.Println("[RM] listener: ")
	dataBuffer := bytes.NewBuffer(buffer[:nRead])
	log.Println("[RM] databuffer: ")
	decode := gob.NewDecoder(dataBuffer)

	err = decode.Decode(&ac)
	f.Error(err, "Receive error  \n")

	log.Println("[RM] -------RECIDO: ", ac, src)

	// SetDeadline(time.Now().Add(4 * time.Second)
	// err = listener.SetWriteDeadline(deadline)
	deadline := time.Now().Add(2 * time.Second)
	err = listener.SetDeadline(deadline)
	if err != nil {
		log.Println("se me acabo el tiempo ")
	}

	log.Println(nRead, addr)

	return err

}
