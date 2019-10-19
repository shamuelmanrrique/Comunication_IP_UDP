package functions

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func S(conn Connection, msm Msm, canal chan Message) error {

	var connection net.Conn
	var err error
	var encoder *gob.Encoder
	id := conn.GetId()

	fmt.Printf("#------------ SEND    %s ----------------# \n", id)
	for i, v := range conn.GetIds() {
		if v != id {
			// fmt.Printf("Envio a la clave 2**%d = %d \\n", i, v)
			delay := time.Duration(conn.GetDelay(i)) * time.Second
			time.Sleep(delay)
			connection, err = net.Dial("tcp", conn.GetId())
			Error(err, "Send connection "+v+" error")

			encoder = gob.NewEncoder(connection)
			err = encoder.Encode(msm)

		}

	}

	// for i in conn
	// time.Sleep(2 * time.Second)
	// // fmt.Println(conn.GetId())
	// connection, err = net.Dial("tcp", conn.GetId())
	// Error(err, "Send connection error")

	// encoder = gob.NewEncoder(connection)
	// err = encoder.Encode(msm)

	// fmt.Println("Estoy en send")

	connection.Close()
	return err

}
