package functions

import (
	// "time"
	"encoding/gob"
	"fmt"
	"net"
)

func Receive(conn Connection, canal chan Message) error {
	var msm Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder
	var listener net.Listener

	id := conn.GetId()

	fmt.Printf("#------------ RECEIVE %s ----------------# \n", id)

	listener, err = net.Listen("tcp", conn.GetPort())
	Error(err, "Listen Error")

	red, err = listener.Accept()
	Error(err, "Server accept red error")

	decoder = gob.NewDecoder(red)
	err = decoder.Decode(&msm)

	// m := "[Receive] => " + connect.GetId() + " He disparado a " + connect.GetKill()

	canal <- msm

	Error(err, "Receive error "+id+" \n")
	fmt.Println(msm.GetData())

	red.Close()
	return err

}
