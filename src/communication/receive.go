package communication

import (
	"encoding/gob"
	"log"
	"net"
	f "sd_paxos/src/functions"
)

/*
-----------------------------------------------------------------
METODO: Receive
RECIBE: canal de tipo f.Message, un listener net.Listener
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function that receive message using TCP connection  
-----------------------------------------------------------------
*/
func Receive(canal chan f.Message, liste net.Listener) error {
	var msm f.Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder

	// Creating a generical connection
	red, err = liste.Accept()
	f.Error(err, "Server accept red error")
	defer red.Close()

	// Reading incomming message and decode its
	decoder = gob.NewDecoder(red)
	err = decoder.Decode(&msm)
	f.Error(err, "Receive error  \n")

	log.Println(" RECEIVE -->: from ", msm.GetFrom(), " to ", msm.GetTo(), "|| OBJ: ", msm.GetTarg(),
		"\n                     Vector: ", msm.GetVector())
	canal <- msm

	return err
}
