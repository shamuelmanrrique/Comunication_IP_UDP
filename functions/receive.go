package functions

import (
	// "time"
	"encoding/gob"
	"fmt"
	"net"
)

func Receive(conn Connection, canal chan<- Message) error {
	var msm Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder
	var listener net.Listener

	id := conn.GetId()

	fmt.Printf("#------------ RECEIVE %s ----------------# \n", conn.GetPort())
	listener, err = net.Listen("tcp", conn.GetPort())
	Error(err, "Listen Error")

	for i := 0; i < len(conn.GetIds())+1; i++ {

		red, err = listener.Accept()
		Error(err, "Server accept red error")

		decoder = gob.NewDecoder(red)
		err = decoder.Decode(&msm)
		canal <- msm

		Error(err, "Receive error "+id+" \n")

		// fmt.Println(msm)
	}
	close(canal)
	// m := "[RECEIVE] => " + conn.GetKill() + "me ha disparado " + conn.GetId()
	// fmt.Println(m)
	red.Close()
	return err

}
