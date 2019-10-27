package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
	"time"
)

// Receive TODO ELIMINAR
func ReceiveMulticast(canal chan f.Message, conn *f.Conn) error {
	var msm f.Message
	// var listener net.PacketConn
	var err error
	const maxBufferSize = 4046

	// NO necesitamos llamar al ResolveUDPAddr
	// log.Println("[RM] mi ip : ", conn.GetId(), "port: ", conn.GetPort())
	// listener, err = net.ListenPacket("udp", conn.GetId())
	localhostAddress, _ := net.ResolveUDPAddr("udp", conn.GetPort())
	// printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	listener, _ := net.ListenMulticastUDP("udp", nil, localhostAddress)
	// printError("DialUDP in Broadcast localhost failed.", e)
	// defer connection.Close()

	f.Error(err, "Listen Error")
	defer listener.Close()

	for {

		log.Println("[RM] -------Estoy Escuchando FOR: ")
		buffer := make([]byte, maxBufferSize)
		log.Println("[RM] buffer: ")
		nRead, addr, err := listener.ReadFrom(buffer)
		f.Error(err, "Error en RM")
		log.Println("[RM] listener: ")
		dataBuffer := bytes.NewBuffer(buffer[:nRead])
		log.Println("[RM] databuffer: ")
		decode := gob.NewDecoder(dataBuffer)
		err = decode.Decode(&msm)
		log.Println("[RM] -------Estoy Escuchando FOR: ", msm)
		f.Error(err, "Receive error  \n")

		deadline := time.Now().Add(2)
		err = listener.SetWriteDeadline(deadline)
		f.Error(err, "Listen Error")
		log.Println(nRead, addr)

	}

	return err

}
