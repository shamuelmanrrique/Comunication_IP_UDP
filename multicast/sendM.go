package multicast

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	f "practice1/functions"
	"time"
)

// SendM functionkl;
func SendM(i interface{}, ip string) error {
	var connection net.Conn
	var red *net.UDPAddr
	var buffer bytes.Buffer
	var err error

	fmt.Println("[SM] estoy en ResolveUDPAddr  ", ip)
	red, err = net.ResolveUDPAddr("udp", ip)
	f.Error(err, "Send connection error \n")

	fmt.Println("[SM] estoy en DialUDP  ", ip)
	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "Send connection error \n")
	defer connection.Close()

	fmt.Println("[SM]  Entre en el for: ", ip)
	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(&i)
	f.Error(err, "Error en broacast: ")
	_, err = connection.Write(buffer.Bytes())

	f.Error(err, "Error al recibir el msm")

	time.Sleep(1 * time.Second)

	return err
}
