package functions

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func Sended(connect Connection) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder
	message := 4
	// host := connect.GetHost()
	// fmt.Println(" address value: %s", address)
	// fmt.Println(host)
	fmt.Println("Stay in send 1")
	time.Sleep(2 * time.Second)
	// connection, err = net.Dial("tcp", host)
	connection, err = net.Dial("tcp", "127.0.0.1:5008")
	Error(err, "Error iniciando send")

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(message)

	fmt.Println("Estoy en send")

	connection.Close()
	return err

}
