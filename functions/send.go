package functions

import (
	"encoding/gob"
	"fmt"
	"net"
)

func Send(msm Msm) error {

	var connection net.Conn
	var err error
	var encoder *gob.Encoder
	id := msm.GetTo()

	fmt.Printf("#------------ SEND    %s ----------------# \n", id)

	connection, err = net.Dial("tcp", id)
	Error(err, "Send connection error \n")

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(msm)

	connection.Close()
	return err

}
