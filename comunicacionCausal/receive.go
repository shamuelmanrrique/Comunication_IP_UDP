package comunicacionCausal

import (
	"encoding/gob"
	"fmt"
	"net"
	f "practice1/functions"
)

func Receive(canal chan f.Message, liste net.Listener, caller string) error {
	var msm f.Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder

	fmt.Println(caller, "is going to wait for a message.")
	red, err = liste.Accept()
	fmt.Println(caller, "has received a message.")
	f.Error(err, "Server accept red error")
	defer red.Close()

	decoder = gob.NewDecoder(red)
	err = decoder.Decode(&msm)

	fmt.Println("Mensaje recibido desde", msm.GetFrom())
	f.Error(err, "Receive error  \n")

	canal <- msm
	return err
}
