package communication

import (
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
)

// Receive is a function that get message by TCP connection  
func Receive(canal chan f.Message, liste net.Listener, caller string) error {
	var msm f.Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder

	red, err = liste.Accept()
	f.Error(err, "Server accept red error")
	defer red.Close()

	decoder = gob.NewDecoder(red)
	err = decoder.Decode(&msm)
	f.Error(err, "Receive error  \n")

	log.Println(" RECEIVE -->: from ", msm.GetFrom(), " to ", msm.GetTo(), "|| OBJ: ", msm.GetTarg(),
		"\n                     Vector: ", msm.GetVector())
	canal <- msm

	return err
}
