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
	var arrayMsms []f.Message
	n := len(connect.GetIds())
	vector := connect.GetVector()
	id := connect.GetId()

	fmt.Println(arrayMsms, n, vector, id)

	// Parse the string address
	addr, _ := net.ResolveUDPAddr("udp", f.MulticastAddress)

	// Open up a connection
	listener, err = net.ListenMulticastUDP("udp", nil, addr)
	f.Error(err, "ReceiveGroupM error ListenPacket")
	defer listener.Close()

	//TODO EL n NO ME QUEDA TAN CLARO
	fmt.Println("[ReceiveGroupM]  Entre en el for con n: ", n)
	for i := 0; i < n; i++ {

		listener.SetReadBuffer(f.MaxBufferSize)

		buffer := make([]byte, f.MaxBufferSize)
		nRead, _, _ := listener.ReadFromUDP(buffer)
		fmt.Println("[ReceiveGroupM] Dentro del for i: ", i, nRead)

		dataBuffer := bytes.NewBuffer(buffer)
		decode = gob.NewDecoder(dataBuffer)
		err = decode.Decode(&msm)
		f.Error(err, "Receive error  Decode\n")

		if msm.GetFrom() != id {
			fmt.Println("[ReceiveGroupM] que yo no envie el msm", msm.GetFrom(), "comparo con ", connect.GetId())

			// Numero de msm a recibir
			ackID := &f.Ack{Code: id + "," + msm.GetFrom()}

			chanelAck := make(chan f.Pack)
			defer close(chanelAck)

			pack := &f.Pack{ConfACK: *ackID}

			// for i := 0; i < n; i++ {
			SendPack(pack, msm.GetFrom())
			// go SendPack(pack, msm.GetFrom())

			// }

		}
	}

	return err

}
