package functions

import (
	"encoding/gob"
	"fmt"
	"net"
)

func Receive(canal chan Message, liste net.Listener) error {
	var msm Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder

	id := msm.GetTo()
	// defer wg.Done()

	red, err = liste.Accept()
	Error(err, "Server accept red error")
	defer red.Close()

	decoder = gob.NewDecoder(red)
	err = decoder.Decode(&msm)
	Error(err, "Receive error "+id+" \n")

	// if msm.GetTo() != id {
	// 	fmt.Printf("[NEWS] %s --> %s \n", id, msm.GetFrom())
	// }

	fmt.Printf("Antes del canal \n")
	canal <- msm
	fmt.Printf("Despuess del canal \n")
	return err
}
