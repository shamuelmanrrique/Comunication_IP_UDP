package functions

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

// func Send(data interface{}) error {
func Send(conect Conection) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder

	fmt.Println("Stay in send 1")
	time.Sleep(2 * time.Second)
	connection, err = net.Dial("tcp", "conect.getIp()" )
	Error(err)

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode("data")

	fmt.Println("Estoy en send")

	connection.Close()
	return err

}
