package functions

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func Sendidi(data int) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder

	fmt.Println("Stay in send 1")
	time.Sleep(2 * time.Second)
	connection, err = net.Dial("tcp", "127.0.0.1:5008")
	if err != nil {
		panic("Client connection error")
	}

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(data)

	fmt.Println("Estoy en send")

	connection.Close()
	return err

}
