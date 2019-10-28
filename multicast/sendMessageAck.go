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
func SendMessageAck( c chan f.Ack, i interface{}, ip string) error {
	// func SendMulticast(ip string, msm *f.Msm, reader io.Reader) error {
	var connection net.Conn
	var red *net.UDPAddr
	var buffer bytes.Buffer
	var err error


	fmt.Println("[SM] Inicio ", ip)
	red, err = net.ResolveUDPAddr("udp", ip)
	f.Error(err, "Send connection error \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "Send connection error \n")
	defer connection.Close()

	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(i)
	f.Error(err, "Error en broacast: ")
	_, err = connection.Write(buffer.Bytes())
	f.Error(err, "Error al recibir el msm")

	time.Sleep(1 * time.Second)

	return err
}
