package functions

import (
	"encoding/gob"
	"net"
	"time"
	"fmt"
)

func S(conn Connection, msm Msm, canal chan Message) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder

	fmt.Println("------------SEND----------------")
	time.Sleep(2*time.Second)
	fmt.Println(conn.GetId())
	connection, err = net.Dial("tcp", conn.GetId() )
	Error(err, "Send connection error")

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(msm)
	
	fmt.Println("Estoy en send")
	
	
	connection.Close()
	return err

}