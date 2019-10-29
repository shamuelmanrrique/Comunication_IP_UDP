package multicast

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	f "practice1/functions"
	"time"
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
	// for i := 0; i < n; i++ {

	listener.SetReadBuffer(f.MaxBufferSize)

	buffer := make([]byte, f.MaxBufferSize)
	nRead, _, _ := listener.ReadFromUDP(buffer)
	fmt.Println("[ReceiveGroupM] Dentro del for i: ", nRead)

	dataBuffer := bytes.NewBuffer(buffer)
	decode = gob.NewDecoder(dataBuffer)
	err = decode.Decode(&msm)
	f.Error(err, "Receive error  Decode\n")

	// RECIBO y sumo 1 al vector
	vector.Tick(id)
	// SEt la nueva actualizacion de recepcion
	connect.SetClock(vector)
	// Uno los relojes
	vector.Merge(msm.GetVector())
	// connect.GetVector().Merge(vector)
	// Seteo nuevamente el reloj
	connect.SetClock(vector)

	fmt.Println("[ReceiveGroupM]  REcibido de: ", msm.GetFrom(), " Yo soy ", id)
	if msm.GetFrom() != id {

		fmt.Println("[ReceiveGroupM]  Target: ", msm.GetTarg(), " Recibio ", id)
		if msm.GetTarg() == id {
			SendGroupM(&msm, connect)
		}

		//Aplico delay receive
		delay := msm.GetDelay()
		time.Sleep(delay * time.Second)

		// 	fmt.Println("[ReceiveGroupM] que yo no envie el msm", msm.GetFrom(), "comparo con ", connect.GetId())

		// 	// Numero de msm a recibir
		// 	ackID := &f.Ack{Code: id + "," + msm.GetFrom()}

		// 	chanelAck := make(chan f.Pack)
		// 	defer close(chanelAck)

		// 	pack := &f.Pack{ConfACK: *ackID}

		// 	// for i := 0; i < n; i++ {
		// 	SendPack(pack, msm.GetFrom())
		// 	// go SendPack(pack, msm.GetFrom())

		// 	// }

	}
	// }

	return err

}
