package functions

import (
	// "time"
	"encoding/gob"
	"fmt"
	"net"
)

func Receive(connect Connection, canal chan Message, liste net.Listener) error {
	var msm Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder
	id := connect.GetId()

	red, err = liste.Accept()
	Error(err, "Server accept red error")
	defer red.Close()

	decoder = gob.NewDecoder(red)
	err = decoder.Decode(&msm)
	Error(err, "Receive error "+id+" \n")

	if msm.GetTo() != id {
		fmt.Printf("[NEWS] %s --> %s \n", msm.GetTo(), msm.GetFrom())

	}

	// fmt.Println(msm.GetFrom())
	// fmt.Println(id)
	// if msm.GetFrom() == id {
	// 	fmt.Printf("[DEAD] => %s -- %s -> %s \n", msm.GetTo(), msm.GetData(), id)
	// 	go SendGroup(connect)
	// }

	canal <- msm
	// close(canal)
	return err

}
