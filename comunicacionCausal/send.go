package comunicacionCausal

import (
	"encoding/gob"
	"fmt"
	"net"
	f "practice1/functions"
)

func Send(msm f.Msm) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder
	sendAddress := msm.GetIgnor()

	connection, err = net.Dial("tcp", sendAddress)
	f.Error(err, "Send connection error \n")
	defer connection.Close()

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(msm)

	if msm.GetFrom() == sendAddress {
		fmt.Printf("[KILL] => To: %s From: %s \n", msm.GetTo(), sendAddress)
	}
	return err

}
