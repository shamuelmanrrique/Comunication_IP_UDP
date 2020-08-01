package chandylamport

import (
	"encoding/gob"
	"log"
	"net"
	f "sd_paxos/src/functions"
)

/*
-----------------------------------------------------------------
METODO: SendC
RECIBE: pack interface{}, IPAddress addr
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function that let you send any message using
			TCP connection. It can send two type of message:
				Marker 	--> Marker chandy lamport
				Message --> Send new message
				String  --> ACK confirmation
-----------------------------------------------------------------
*/
func SendC(pack interface{}, addr string) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder

	// Making dial connection to ip address
	connection, err = net.Dial("tcp", addr)
	defer connection.Close()

	// Encoder and send message
	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(&pack)

	// Getting pack tipe message to print this
	switch packNew := pack.(type) {
	case *f.Message:
		log.Println(" ++> SEND MSM: from ", packNew.GetFrom(), " to ", packNew.GetTo(), "  || OBJ: ", packNew.GetTarg(),
			"\n                     Vector: ", packNew.GetVector())
	case *f.Marker:
		log.Println(" ++> SEND Marker: Init Marker", packNew)
	case *string:
		log.Println(" ++> SEND Count: ", packNew)

	}

	return err

}
