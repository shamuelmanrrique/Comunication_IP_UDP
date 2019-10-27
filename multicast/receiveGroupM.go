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
func ReceiveGroupM(canal chan f.Message, connect *f.Conn) error {
	var err error
	// var red net.Conn
	var msm f.Message
	var decode *gob.Decoder
	var listener *net.UDPConn
	// var nRead int
	// var listener net.PacketConn

	// Parse the string address
	addr, _ := net.ResolveUDPAddr("udp", f.MulticastAddress)

	// Open up a connection
	listener, err = net.ListenMulticastUDP("udp", nil, addr)
	// listener, err = net.ListenPacket("udp", f.MulticastAddress)
	f.Error(err, "ReceiveGroupM error ListenPacket")
	defer listener.Close()

	listener.SetReadBuffer(f.MaxBufferSize)

	// Loop forever reading from the socket
	for {

		fmt.Println("[ReceiveGroupM]  Entre en el for: ", connect.GetId())
		buffer := make([]byte, f.MaxBufferSize)
		nRead, src, _ := listener.ReadFromUDP(buffer)
		fmt.Println("[ReceiveGroupM] print BUFFER: ", src)

		// nRead, addr, err := listener.ReadFrom(buffer)
		dataBuffer := bytes.NewBuffer(buffer)
		decode = gob.NewDecoder(dataBuffer)
		fmt.Println("[ReceiveGroupM]  DECODE: ", decode)
		err = decode.Decode(&msm)
		f.Error(err, "Receive error  Decode\n")

		deadline := time.Now().Add(2)
		err = listener.SetWriteDeadline(deadline)
		f.Error(err, "ReceiveGroupM Error SetWriteDeadline ")
		fmt.Println(nRead, addr)

		fmt.Println("[ReceiveGroupM] IF: PRINT MSM", msm)

		if msm.GetTo() == connect.GetId() {
			break
		}

	}

	return err

}
