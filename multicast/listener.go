package multicast

import (
	"encoding/gob"
	"fmt"
	"net"
)

func Listener(connect Connection) error {
	var connection net.Conn
	var err error
	var decoder *gob.Decoder
	var listener net.Listener

	listener, err = net.ResolveUDPAddr("udp", ":5008")
	Error(err, "Server listen error")

	connection, err = listener.Accept()
	Error(err, "Server accept connection error")

	decoder = gob.NewDecoder(connection)
	err = decoder.Decode(0)
	connection.Close()

	fmt.Println("Estoy en receive")
	connection.Close()
	return err

}

