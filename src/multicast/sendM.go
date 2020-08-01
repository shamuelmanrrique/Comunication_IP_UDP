package multicast

import (
	"bytes"
	"encoding/gob"
	"log"

	"net"
	f "sd_paxos/src/functions"
)

/*
-----------------------------------------------------------------
METODO: Send
RECIBE: message interface{}, IPAddress ip
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function that let you send any message using 
			UDP connection. It can send two type of message:
				ACK 	--> confirming receipt of message
				Message --> send new message
-----------------------------------------------------------------
*/
func SendM(i interface{}, ip string) error {
	var connection net.Conn
	var red *net.UDPAddr
	var buffer bytes.Buffer
	var err error

	// Making dial connection to ip address
	red, err = net.ResolveUDPAddr("udp", ip)
	f.Error(err, "Send connection error \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "Send connection error \n")
	defer connection.Close()

	// Encoder and send message
	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(&i)
	f.Error(err, "Error en broacast: ")
	_, err = connection.Write(buffer.Bytes())
	f.Error(err, "Error al recibir el msm")

	// Print message according to type
	switch packNew := i.(type) {
	case *f.Message:
		log.Println(" ++> SEND : from ", packNew.GetFrom(), " to ", packNew.GetTo(), "  || OBJ: ", packNew.GetTarg(),
			"\n                     Vector: ", packNew.GetVector())
	case *f.Ack:
		log.Println(" ++> SEND : ACK to: ", ip, "code:", packNew)
	}

	return err
}
