package functions

import (
	// "time"
	"encoding/gob"
	"fmt"
	"net"
)

func Receive(connect Connection, canal chan Message) error {
	var msm Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder
	var listener net.Listener

	id := connect.GetId()

	listener, err = net.Listen("tcp", connect.GetPort())
	Error(err, "Listen Error")

	for i := 0; i < len(connect.GetIds())+1; i++ {
		red, err = listener.Accept()
		Error(err, "Server accept red error")

		decoder = gob.NewDecoder(red)
		err = decoder.Decode(&msm)
		Error(err, "Receive error "+id+" \n")

		if msm.GetFrom() == id {
			go SendGroup(connect)
		}

		canal <- msm
		fmt.Printf("RECEIVE => To: %s From: %s \n", id, msm.GetTo())
	}

	red.Close() //lo tenia dentro del for
	listener.Close()
	close(canal)
	return err

}
