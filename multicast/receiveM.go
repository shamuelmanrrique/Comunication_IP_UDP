package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
	"time"
)

// Receive function to management communication
func ReceiveM(chanAc chan<- f.Ack, chanMes chan<- f.Message, caller string) error {
	var pack interface{}
	var err error

	red, _ := net.ResolveUDPAddr("udp", caller)

	listener, err := net.ListenUDP("udp", red)
	f.Error(err, "[ReceiveM] ListenUDP Error")
	defer listener.Close()

	timeoutDuration := 40 * time.Second
	listener.SetReadDeadline(time.Now().Add(timeoutDuration))

	for {
		buffer := make([]byte, f.MaxBufferSize)
		nRead, _, err := listener.ReadFrom(buffer)
		if err != nil {
			return err
		}

		dataBuffer := bytes.NewBuffer(buffer[:nRead])
		decode := gob.NewDecoder(dataBuffer)

		err = decode.Decode(&pack)
		f.Error(err, "Receive error  \n")

		switch packNew := pack.(type) {
		case f.Message:
			chanMes <- packNew
			log.Println(" RECEIVE -->: from ", packNew.GetFrom(), " to ", packNew.GetTo(), "  || OBJ: ", packNew.GetTarg(),
				"\n                     Vector: ", packNew.GetVector())
		case f.Ack:
			chanAc <- packNew
			log.Println(" RECEIVE --> ACK from: ", packNew.GetOrigen(), "code:", packNew)
		}
	}

	return err
}
