package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
)

// ReceiveM akma
func ReceiveM(canal chan *interface{}, addr string) error {
	var pack interface{}
	var err error
	var listener *net.UDPConn

	//Duda solo debo pasar el puerto o el ip completo por ser UDP
	red, _ := net.ResolveUDPAddr("udp", addr)
	log.Println("[RM]             localhostAddress ", red)

	// printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	listener, err = net.ListenUDP("udp", red)
	f.Error(err, "[RM] ListenUDP Error")
	defer listener.Close()

	buffer := make([]byte, f.MaxBufferSize)
	log.Println("[RM] buffer: ")
	// nRead, src, err := listener.ReadFrom(buffer)
	f.Error(err, "Error en RM")
	// log.Println("[RM] li2048stener: ")
	// dataBuffer := bytes.NewBuffer(buffer[:nRead])
	dataBuffer := bytes.NewBuffer(buffer[:f.MaxBufferSize])
	log.Println("[RM] databuffer: ")
	decode := gob.NewDecoder(dataBuffer)

	err = decode.Decode(pack)
	f.Error(err, "Receive error  \n")

	log.Println("[RM] -------RECIDO: ", pack)
	// canal <- *interf

	// SetDeadline(time.Now().Add(4 * time.Second)
	// err = listener.SetWriteDeadline(deadline)
	// deadline := time.Now().Add(2 * time.Second)
	// err = listener.SetDeadline(deadline)
	// if err != nil {
	// 	log.Println("se me acabo el tiempo ")
	// }

	// log.Println(nRead, addr)

	return err

}
