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

// SendGroupM function
func SendGroupM(msm *f.Message, connect *f.Conn) error {
	var red *net.UDPAddr
	var connection *net.UDPConn
	var encoder *gob.Encoder
	var buffer bytes.Buffer
	var err error

	fmt.Println("[SendGroupM] Inicio ", connect.GetId())
	red, err = net.ResolveUDPAddr("udp", f.MulticastAddress)
	f.Error(err, "SendGroupM error ResolveUDPAddr connection \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "SendGroupM error DialUDP connection \n")
	defer connection.Close()

	encoder = gob.NewEncoder(&buffer)
	err = encoder.Encode(msm)
	f.Error(err, "SendGroupM encoder error \n")
	_, err = connection.Write(buffer.Bytes())
	f.Error(err, "Error al recibir el msm")

	var bufferAcks []f.Ack
	var bufferMessage []f.Message

	// var canalAcks chan f.Ack
	// var canalMessage chan f.Message
	n := len(connect.Ids)
	for i := 0; i < n-1; i++ {
		// go ReceiveM(canalAcks, canalMessage, connect.GetId())

		fmt.Println(bufferAcks, bufferMessage)
	}

	deadline := time.Now().Add(2 * time.Second)
	err = connection.SetDeadline(deadline)
	if err != nil {
		log.Println("se me acabo el tiempo ")
	}

	return err

}
