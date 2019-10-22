package functions

import (
	"encoding/gob"
	"fmt"
	"net"
)

func Received(connect Connection) error {
	var connection net.Conn
	var err error
	var decoder *gob.Decoder
	var listener net.Listener
	// port := ":" + connect.GetHost()

	fmt.Println("Stay in receive 1")
	listener, err = net.Listen("tcp", ":5008")
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
