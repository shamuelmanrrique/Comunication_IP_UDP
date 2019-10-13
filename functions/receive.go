package functions

import (
	"encoding/gob"
	"fmt"
	"net"
)

// func Receive(data interface{}) error {
func Receive(data struct{}) error {
	var connection net.Conn
	var err error
	var decoder *gob.Decoder
	var listener net.Listener

	fmt.Println("Stay in receive 1")
	listener, err = net.Listen("tcp", ":5008")
	if err != nil {
		panic("Server listen error")
	}
	
	connection, err = listener.Accept()
	if err != nil {
		panic("Server accept connection error")
	}

	decoder = gob.NewDecoder(connection)
	err = decoder.Decode(data)
	connection.Close()

	fmt.Println("Estoy en receive")
	connection.Close()
	return err

}
