package multicast

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	f "practice1/functions"
	"time"
)

// SendGroupM function
func SendGroupM(msm *f.Message, connect *f.Conn) error {
	var red *net.UDPAddr
	var connection *net.UDPConn
	var encoder *gob.Encoder
	var buffer bytes.Buffer
	var err error

	red, err = net.ResolveUDPAddr("udp", f.MulticastAddress)
	f.Error(err, "SendGroupM error ResolveUDPAddr connection \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "SendGroupM error DialUDP connection \n")
	defer connection.Close()

	for {
		fmt.Println("[SendGroupM]  Entre en el for: ", connect.GetId())

		encoder = gob.NewEncoder(&buffer)
		err = encoder.Encode(msm)
		f.Error(err, "SendGroupM encoder error \n")
		_, err = connection.Write(buffer.Bytes())
		f.Error(err, "Error al recibir el msm")
		// buffer.Reset()

		time.Sleep(1 * time.Second)
	}

	return err

}
