package multicast

import (
	// "bytes"
	// "encoding/gob"
	// "net"
	// f "sd_paxos/src/functions"
	"encoding/hex"
	"log"
	"net"

	"github.com/dmichael/go-multicast/multicast"
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
	// var decode *gob.Decoder
	// var listener *net.UDPConn

	// // Set up and address to receive message
	// addr, _ := net.ResolveUDPAddr("udp", "229.0.40.000:9999")
	// f.Error(err, "11111111111 ReceiveGroupMr error ResolveUDPAddr connection \n")

	// listener, err = net.ListenMulticastUDP("udp", nil, addr)
	// f.Error(err, "22222222222 ReceiveGroupM error ListenPacket")
	// defer listener.Close()

	// println("++++++++++++++++++> ReceiveGroupM DIRECCION", "229.0.40.000:9999")

	// // Gorutine to receive UNICAST AND MULTICAST message

	// // Defining timeout to wait MULTICAST menssage
	// // deadline := time.Now().Add(40 * time.Second)

	// for {
	// 	println("++++++++++++++++++> ReceiveGroupM for por por 40s")
	// 	var msm string
	// 	// Set up to read message using buffer
	// 	listener.SetReadBuffer(12000)
	// 	buffer := make([]byte, 12000)
	// 	listener.ReadFromUDP(buffer)

	// 	// Reding from buffer and decoder message
	// 	dataBuffer := bytes.NewBuffer(buffer)
	// 	decode = gob.NewDecoder(dataBuffer)
	// 	err = decode.Decode(&msm)

	// 	println("|||||||||||||||||||||||||||", &msm)
	// 	if err != nil {
	// 		f.Error(err, "Receive error  \n")
	// 		break
	// 	}

	// 	println("++++++> ReceiveGroupM llego sms de ", msm)

	// 	println("NO SOY YO EL ORIGEN ", msm)

	// }
	multicast.Listen("229.0.40.000:9999", msgHandler)
	println("++++++> receiveChannel ")

	return err

}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "AAAAAAAA from", src)
	log.Println(hex.Dump(b[:n]))
}
