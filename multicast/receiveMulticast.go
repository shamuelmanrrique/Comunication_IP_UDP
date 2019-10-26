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
	var listener net.PacketConn
	var err error
	const maxBufferSize = 2048

	// var err error
	// localAddress, err := net.ResolveUDPAddr("udp", port)
	// connection, err := net.ListenUDP("udp", localAddress)
	// defer connection.Close()
	// var message CommData
	// NO necesitamos llamar al ResolveUDPAddr

	// fmt.Println("[RM]  El 	IP: ", conn.GetId())
	log.Println("[RM] mi ip : ", conn.GetId(), "port: ", conn.GetPort())
	listener, err = net.ListenPacket("udp", conn.GetPort())
	f.Error(err, "Listen Error")
	defer listener.Close()

	for {

		log.Println("[RM] ITS HERE")

		buffer := make([]byte, maxBufferSize)
		nRead, addr, err := listener.ReadFrom(buffer)
		dataBuffer := bytes.NewBuffer(buffer[:nRead])
		decode := gob.NewDecoder(dataBuffer)
		err = decode.Decode(&msm)
		f.Error(err, "Receive error  \n")

		deadline := time.Now().Add(2)
		err = listener.SetWriteDeadline(deadline)
		f.Error(err, "Listen Error")
		log.Println(nRead, addr)

	}

	return err

}
