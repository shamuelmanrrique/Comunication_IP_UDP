package multicast

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	f "practice1/functions"
)

// ReceiveGroupM  das
func ReceiveGroupM(connect *f.Conn) error {
	var err error
	var msm f.Message
	var decode *gob.Decoder
	var listener *net.UDPConn

	// Parse the string address
	addr, _ := net.ResolveUDPAddr("udp", f.MulticastAddress)

	// Open up a connection
	listener, err = net.ListenMulticastUDP("udp", nil, addr)
	// listener, err = net.ListenPacket("udp", f.MulticastAddress)
	f.Error(err, "ReceiveGroupM error ListenPacket")
	defer listener.Close()

	listener.SetReadBuffer(f.MaxBufferSize)

	// Loop forever reading from the socket

	var arrayMsms []f.Message
	n := connect.GetAccept()
	vector := connect.GetVector()
	id := connect.GetId()
	fmt.Println("[ReceiveGroupM] ", arrayMsms, n, vector, id)
	// for {

	//TODO EL n NO ME QUEDA TAN CLARO
	fmt.Println("[ReceiveGroupM]  Entre en el for con n: ", n)
	for i := 0; i < n; i++ {
		fmt.Println("[ReceiveGroupM] DEntro del for i: ", i)
		// Recibo de multicast un numero de veces (aun no se )

		// Si recibe por multicast envio un ack de confirmación

		buffer := make([]byte, f.MaxBufferSize)
		nRead, _, _ := listener.ReadFromUDP(buffer)
		// fmt.Println("[ReceiveGroupM] print BUFFER: ", src)

		dataBuffer := bytes.NewBuffer(buffer)
		decode = gob.NewDecoder(dataBuffer)
		// fmt.Println("[ReceiveGroupM]  DECODE: ", decode)
		err = decode.Decode(&msm)
		// f.Error(err, "Receive error  Decode\n")

		// deadline := time.Now().Add(2)
		// err = listener.SetWriteDeadline(deadline)
		// f.Error(err, "ReceiveGroupM Error SetWriteDeadline ")
		// fmt.Println(nRead, addr)

		//Recibo el msm y envio el ack

		fmt.Println("[ReceiveGroupM] IF: PRINT MSM", msm, nRead)

		ackID := &f.Ack{Code: msm.GetTo() + "," + msm.GetFrom()}

		// Send confirmacion ack
		go SendM(ackID, connect.GetId())

		// Numero de msm a recibir
		// n := len(connect.GetIds())

		// Creo un buffer de Ack
		bufferAck := make(chan f.Ack)
		defer close(bufferAck)

		bufferMessage := make(chan f.Message)
		defer close(bufferMessage)

		// for i := 0; i < n; i++ {
		// 	fmt.Println(ackID)
		// 	// COmo limite que voy a escuchar un ACK O MSM
		// 	// go ReceiveM(bufferMessage, connect)

		// }

		if msm.GetTo() == connect.GetId() {
			break
		}

	}

	return err

}

// func SendAck(id string, addr string)  {

// }
