package chandylamport

import (
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
)

// Send function
func SendC(pack interface{}, addr string) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder

	// log.Println("[Send] chandy Envio a ", addr)
	connection, err = net.Dial("tcp", addr)
	// f.Error(err, "Send connection error \n")
	defer connection.Close()

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(&pack)

	switch packNew := pack.(type) {
	case *f.Message:
		log.Println(" ++> SEND : from ", packNew.GetFrom(), " to ", packNew.GetTo(), "  || OBJ: ", packNew.GetTarg(),
			"\n                     Vector: ", packNew.GetVector())
	case *f.Marker:
		// log.Println("[ReceiveM] ----> Marker ", packNew, " DE ", packNew.GetCounter())
		log.Println(" ++> SEND : Init Marker", packNew)
	case *string:
		// log.Println("[ReceiveM] ----> checkpoint ", packNew )
		log.Println(" ++> SEND : Count ", packNew)

	}

	return err

}
