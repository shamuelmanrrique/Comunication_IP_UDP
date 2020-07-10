package communication

import (
	"encoding/gob"
	"log"
	"net"
	f "sd_paxos/functions"
)

// Send function
func Send(ip string, msm f.Msm, caller string) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder

	connection, err = net.Dial("tcp", ip)
	f.Error(err, "Send connection error \n")
	defer connection.Close()

	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(msm)

	log.Println(" ++> SEND : from ", caller, " to ", ip, "|| OBJ: ", msm.GetTarg(),
		"\n                     Vector: ", msm.GetVector())
	return err

}
