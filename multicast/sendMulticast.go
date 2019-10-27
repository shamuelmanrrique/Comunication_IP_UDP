package multicast

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	f "practice1/functions"
	"time"
)

// SendMulticast functionkl;
func SendMulticast(msm *f.Message, conn *f.Conn) error {
	// func SendMulticast(ip string, msm *f.Msm, reader io.Reader) error {
	var connection net.Conn
	var red *net.UDPAddr
	var buffer bytes.Buffer
	var err error

	red, err = net.ResolveUDPAddr("udp", conn.GetId())
	f.Error(err, "Send connection error \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "Send connection error \n")
	defer connection.Close()

	// log.Println("Estoy en SM")
	// var buffer bytes.Buffer

	for {
		fmt.Println("[SM]  Entre en el for: ", conn.GetId())

		encoder := gob.NewEncoder(&buffer)
		// for  i := 0; i < 5000; i++{
		err = encoder.Encode(msm)
		f.Error(err, "Error en broacast: ")
		_, err = connection.Write(buffer.Bytes())
		f.Error(err, "Error al recibir el msm")
		buffer.Reset()

		time.Sleep(2 * time.Second)
	}

	return err
}
