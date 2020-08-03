package multicast

import (
	"bytes"
	"encoding/gob"
	"net"
	f "sd_paxos/src/functions"
)

/*
-----------------------------------------------------------------
METODO: ReceiveGroupM
RECIBE: canal de tipo f.Message, un listener net.Listener
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function that receive message using UDP connection  and IP multicast
-----------------------------------------------------------------
*/
func ReceiveGroupMr() error {
	var err error
	var decode *gob.Decoder
	var listener *net.UDPConn

	// Set up and address to receive message
	addr, _ := net.ResolveUDPAddr("udp", "229.0.40.000:9999")
	f.Error(err, "11111111111 ReceiveGroupMr error ResolveUDPAddr connection \n")

	listener, err = net.ListenMulticastUDP("udp", nil, addr)
	f.Error(err, "22222222222 ReceiveGroupM error ListenPacket")
	defer listener.Close()

	println("++++++++++++++++++> ReceiveGroupM DIRECCION", "229.0.40.000:9999")

	// Gorutine to receive UNICAST AND MULTICAST message

	// Defining timeout to wait MULTICAST menssage
	// deadline := time.Now().Add(40 * time.Second)

	for {
		println("++++++++++++++++++> ReceiveGroupM for por por 40s")
		// var msm f.Message
		var msm string
		// Set up to read message using buffer

		println("+++++++ Set up to read message using buffer")
		// listener.SetReadBuffer(8192)
		// buffer := make([]byte, 8192)
		// println(111111111111111)
		// listener.ReadFromUDP(buffer)
		// println(222222222222222)

		gob.NewDecoder(red)
		println("+++++++ Reding from buffer and decoder message")
		// Reding from buffer and decoder message
		dataBuffer := bytes.NewBuffer(buffer)
		decode = gob.NewDecoder(dataBuffer)
		err = decode.Decode(&msm)
		println("|||||||||||||||||||||||||||", &msm)
		if err != nil {
			f.Error(err, "Receive error  \n")
			break
		}

		println("++++++> ReceiveGroupM llego sms de ", msm)

		println("NO SOY YO EL ORIGEN ", msm)

	}

	println("++++++> receiveChannel ")

	return err

}
