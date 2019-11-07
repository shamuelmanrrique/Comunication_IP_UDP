package multicast

import (
	"bytes"
	"encoding/gob"
	"log"

	"net"
	f "practice1/functions"
)

// SendM function
func SendM(i interface{}, ip string) error {
	var connection net.Conn
	var red *net.UDPAddr
	var buffer bytes.Buffer
	var err error

	red, err = net.ResolveUDPAddr("udp", ip)
	f.Error(err, "Send connection error \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "Send connection error \n")
	defer connection.Close()

	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(&i)
	f.Error(err, "Error en broacast: ")
	_, err = connection.Write(buffer.Bytes())
	f.Error(err, "Error al recibir el msm")

	switch packNew := i.(type) {
	case *f.Message:
		log.Println(" ++> SEND : from ", packNew.GetFrom(), " to ", packNew.GetTo(), "  || OBJ: ", packNew.GetTarg(),
			"\n                     Vector: ", packNew.GetVector())
	case *f.Ack:
		log.Println(" ++> SEND : ACK to: ", ip, "code:", packNew)
	}

	return err
}
