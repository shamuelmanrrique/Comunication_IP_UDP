package comunicacionCausal

import (
	"encoding/gob"
	"net"
	f "practice1/functions"
)

func Receive(canal chan f.Message, liste net.Listener) error {
	var msm f.Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder

	id := msm.GetTo()
	// defer wg.Done()

	red, err = liste.Accept()
	f.Error(err, "Server accept red error")
	defer red.Close()

	decoder = gob.NewDecoder(red)
	err = decoder.Decode(&msm)
	f.Error(err, "Receive error "+id+" \n")

	// if msm.GetTo() != id {
	// 	fmt.Printf("[NEWS] %s --> %s \n", id, msm.GetFrom())
	// }
	canal <- msm
	return err
}
