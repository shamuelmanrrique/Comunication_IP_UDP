package communication

import (
	"encoding/gob"
	"log"
	"net"
	f "sd_paxos/src/functions"
)

/*
-----------------------------------------------------------------
METODO: Send
RECIBE: ip address "ip",message "f.Msm", sender "caller"
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function that send message using TCP connection  
-----------------------------------------------------------------
*/
func Send(ip string, msm f.Msm, caller string) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder

	// Making dial connection to ip address
	connection, err = net.Dial("tcp", ip)
	f.Error(err, "Send connection error \n")
	defer connection.Close()

	// Encoder and send message
	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(msm)

	log.Println(" ++> SEND : from ", caller, " to ", ip, "|| OBJ: ", msm.GetTarg(),
		"\n                     Vector: ", msm.GetVector())
	return err

}
