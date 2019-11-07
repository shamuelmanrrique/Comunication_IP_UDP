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

	connection, err = net.Dial("tcp", addr)
	defer connection.Close()

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(&pack)

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
