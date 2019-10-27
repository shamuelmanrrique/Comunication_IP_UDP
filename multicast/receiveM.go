package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
)

// Receive TODO ELIMINAR
// func ReceiveM(canalAck chan f.Ack, canal chan f.Message, conn *f.Conn) error {
func ReceiveM(canal chan f.Message, addr string) error {
	// func ReceiveM(canal chan f.Message, conn *f.Conn) error {
	// var msm f.Message
	var ac f.Ack
	// var listener net.PacketConn
	var err error
	const maxBufferSize = 4046

	log.Println("[RM]             ReceiveM ")

	// NO necesitamos llamar al ResolveUDPAddr
	// log.Println("[RM] mi ip : ", conn.GetId(), "port: ", conn.GetPort())
	// listener, err = net.ListenPacket("udp", conn.GetId())

	//Duda solo debo pasar el puerto o el ip completo por ser UDP
	localhostAddress, _ := net.ResolveUDPAddr("udp", addr)
	log.Println("[RM]             localhostAddress ", localhostAddress)
	// localhostAddress, _ := net.ResolveUDPAddr("udp", conn.GetPort())

	// printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	listener, _ := net.ListenUDP("udp", localhostAddress)
	log.Println("[RM]             listener ", listener)
	// printError("DialUDP in Broadcast localhost failed.", e)
	// defer connection.Close()

	f.Error(err, "Listen Error")
	defer listener.Close()

	for {

		// log.Println("[RM] -------Estoy Escuchando FOR: ")
		buffer := make([]byte, maxBufferSize)
		// log.Println("[RM] buffer: ")
		nRead, addr, err := listener.ReadFrom(buffer)
		// f.Error(err, "Error en RM")
		// log.Println("[RM] listener: ")
		dataBuffer := bytes.NewBuffer(buffer[:nRead])
		log.Println("[RM] databuffer: ")
		decode := gob.NewDecoder(dataBuffer)

		err = decode.Decode(&ac)
		// err = decode.Decode(&msm)
		// if err != nil {
		// 	err = decode.Decode(&ac)
		// 	fmt.Println("[RM] .....recibi un msm", ac)

		// }
		// err = decode.Decode(&msm)
		log.Println("[RM] -------RECIDO: ", ac)
		f.Error(err, "Receive error  \n")

		// deadline := time.Now().Add(2)
		// err = listener.SetWriteDeadline(deadline)
		// f.Error(err, "Listen Error")

		log.Println(nRead, addr)

		// ackID := &f.Ack{Code: "GABO GAY"}
		// GO u.SendM(ackID, "127.0.1.1:1400")

	}

	return err

}
