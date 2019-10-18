package functions

import (
	"encoding/gob"
	"fmt"
	"net"
)

func R(conn Connection, msm Message) error {
	var connection net.Conn
	var err error
	var decoder *gob.Decoder
	var listener net.Listener

	listener, err = net.Listen("tcp", conn.GetPort())
	if err != nil {
		panic("Server listen error")
	}
	
	connection, err = listener.Accept()
	if err != nil {
		panic("Server accept connection error")
	}

	decoder = gob.NewDecoder(connection)
	err = decoder.Decode(conn)

	fmt.Println("Estoy en receive")
	connection.Close()
	return err

}