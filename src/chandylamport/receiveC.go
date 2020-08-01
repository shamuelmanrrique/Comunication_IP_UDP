package chandylamport

import (
	"encoding/gob"
	"log"
	"net"
	f "sd_paxos/src/functions"
)

/*
-----------------------------------------------------------------
METODO: ReceiveC
RECIBE: canal de tipo string, canal de tipo f.Message, IP address addr
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function receive message using TCP connection.
			It can receive three type of message:
				Marker 	--> Marker chandy lamport
				Message --> receive new message
				String  --> ACK confirmation
-----------------------------------------------------------------
*/
func ReceiveC(chanPoint chan<- string, chanMar chan<- f.Marker, chanMes chan<- f.Message, addr string) error {
	var listener net.Listener
	var decoder *gob.Decoder
	var pack interface{}
	var red net.Conn
	var err error

	// log.Println("[Receive] Port:  ", addr)

	listener, err = net.Listen("tcp", addr)
	f.Error(err, "Listen Error")
	defer listener.Close()

	for {
		// Set up and address to receive message
		red, err = listener.Accept()
		f.Error(err, "Server accept red error")

		// Reading incomming message and decode its
		decoder = gob.NewDecoder(red)
		err = decoder.Decode(&pack)
		f.Error(err, "Receive error  \n")

		// Case to filter incoming message using type
		switch packNew := pack.(type) {
		case f.Message:
			chanMes <- packNew
			log.Println(" RECEIVE -->: from ", packNew.GetFrom(), " to ", packNew.GetTo(), "  || OBJ: ", packNew.GetTarg(),
				"\n                     Vector: ", packNew.GetVector())
		case f.Marker:
			chanMar <- packNew
			log.Println(" RECEIVE -->: Init Marker:", packNew)
		case string:
			chanPoint <- packNew
			log.Println(" RECEIVE --> ACK from: ", packNew)

		}

		red.Close()

	}

	return err
}
