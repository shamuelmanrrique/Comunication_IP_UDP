package functions

import (
	// "time"
	"encoding/gob"
	"fmt"
	"net"
)

func R( conn Connection, canal chan Message) error {
	var msm Message
	var red net.Conn
	var err error
	var decoder *gob.Decoder
	var listener net.Listener

	fmt.Println( conn.GetPort())
	listener, err = net.Listen("tcp", conn.GetPort())
	Error(err, "Listen Error")
	
	fmt.Println("------------REceive----------------")
	red, err = listener.Accept()
	Error(err, "Server accept red error")

	decoder = gob.NewDecoder(red)
	err = decoder.Decode(&msm)

	if err != nil {
		panic("lloro"+ err.Error())
	}
	fmt.Println("Estoy en receive----")
	fmt.Println(msm)
	fmt.Println("----------Estoy en receive")

	red.Close()
	return err

}