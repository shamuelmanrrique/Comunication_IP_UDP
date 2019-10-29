package multicast

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	f "practice1/functions"
)

// SendGroupM function
func SendGroupM(msm *f.Message, connect *f.Conn) error {
	// var red, redUDP *net.UDPAddr
	// var connection, connectionUDP *net.UDPConn
	var red *net.UDPAddr
	var connection *net.UDPConn
	var encoder *gob.Encoder
	var buffer bytes.Buffer
	// n := len(connect.GetIds())
	var err error

	fmt.Println("[SendGroupM] Inicio ", connect.GetId())
	red, err = net.ResolveUDPAddr("udp", f.MulticastAddress)
	f.Error(err, "SendGroupM error ResolveUDPAddr connection \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "SendGroupM error DialUDP connection \n")
	defer connection.Close()

	// Get msm from buffer
	encoder = gob.NewEncoder(&buffer)
	err = encoder.Encode(msm)
	f.Error(err, "SendGroupM encoder error \n")
	_, err = connection.Write(buffer.Bytes())
	f.Error(err, "Error al recibir el msm")

	// // Define connection to udp
	// redUDP, err = net.ResolveUDPAddr("udp", connect.GetId())
	// f.Error(err, "Send connection error \n")

	// connectionUDP, err = net.ListenUDP("udp", redUDP)
	// f.Error(err, "Send connection error \n")
	// defer connectionUDP.Close()

	// // red, _ := net.ResolveUDPAddr("udp", connectM.GetId())
	// // log.Println("[RM]             localhostAddress ", red)

	// // // printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	// // listener, err := net.ListenUDP("udp", red)
	// // f.Error(err, "[RM] ListenUDP Error")
	// // defer listener.Close()

	// var bufferPacks []f.Pack
	// canalPacks := make(chan f.Pack)

	// for i := 0; i < n-1; i++ {
	// 	go ReceivePack(canalPacks, connectionUDP, connect.GetId())

	// 	pt, _ := <-canalPacks
	// 	fmt.Println("[SendGroupM] recibo del canal: ", pt)
	// 	bufferPacks = append(bufferPacks, pt)

	// }

	// deadline := time.Now().Add(2 * time.Second)
	// err = connection.SetDeadline(deadline)

	

	if err != nil {
		log.Println("se me acabo el tiempo ")
	}

	return err

}
