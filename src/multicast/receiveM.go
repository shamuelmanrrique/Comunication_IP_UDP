package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "sd_paxos/src/functions"
	"time"
)

/*
-----------------------------------------------------------------
METODO: ReceiveM
RECIBE: canal de tipo f.Ack, canal de tipo f.Message, un listener net.Listener
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function that receive message using UDP connection
-----------------------------------------------------------------
*/
func ReceiveM(chanAc chan<- f.Ack, chanMes chan<- f.Message, caller string) error {
	var pack interface{}
	var err error

	// Creting upd connection
	red, _ := net.ResolveUDPAddr("udp", caller)

	// Set address to listener message
	listener, err := net.ListenUDP("udp", red)
	f.Error(err, "[ReceiveM] ListenUDP Error")
	defer listener.Close()

	// Defining time to receive menssage
	timeoutDuration := 38 * time.Second
	listener.SetReadDeadline(time.Now().Add(timeoutDuration))

	// Infinite loop to receive message,
	// finished when timeout duration is off
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

		// Case to filter incoming message using type
		switch packNew := pack.(type) {
		case f.Message:
			chanMes <- packNew
			log.Println(" RECEIVE MESSAGE --> from ", packNew.GetFrom(), " to ", packNew.GetTo(), "  || OBJ: ", packNew.GetTarg(),
				"\n                     Vector: ", packNew.GetVector())
		case f.Ack:
			chanAc <- packNew
			log.Println(" RECEIVE ACK --> from: ", packNew.GetOrigen(), "code:", packNew)
		}
	}

	return err
}
