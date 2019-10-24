package comunicacionCausal

import (
	"encoding/gob"
	"fmt"
	"net"
	f "practice1/functions"
)

// Send function
func Send(ip string, msm f.Msm, caller string) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder
	fmt.Println(caller, "is trying to send to", ip)
	connection, err = net.Dial("tcp", ip)

	f.Error(err, "Send connection error \n")
	defer connection.Close()

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(msm)
	fmt.Println(caller, "has sent to", ip)

	// if msm.GetFrom() == sendAddress {
	// 	fmt.Printf("[KILL] => To: %s From: %s \n", msm.GetTo(), sendAddress)
	// }
	return err

}
