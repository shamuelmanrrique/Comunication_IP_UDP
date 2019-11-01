package multicast

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	f "practice1/functions"
	"time"
)

// SendGroupM send message to ip multicast and wait ack
func SendGroupM(msm *f.Message, connect *f.Conn) error {
	var red *net.UDPAddr
	var connection *net.UDPConn
	var encoder *gob.Encoder
	var buffer bytes.Buffer
	var err error

	target := ""
	delay, _ := time.ParseDuration("0s")
	inf := "Me mataron"
	id := connect.GetId()

	fmt.Println("[SendGroupM] Updating vclock and copy it")
	// Update vClock
	vector := connect.GetVector()
	vector.Tick(id)
	connect.SetClock(vector)

	// Create copy of vector to send into the message
	copyVector := vector.Copy()

	// Check if I have a target
	if len(connect.GetKill()) > 0 && len(connect.GetDelays()) > 0 {
		fmt.Println("[SendGroupM] I got a target")
		target = connect.GetTarget(0)
		delay = connect.GetDelay(0)
		inf = "He disparado"
		connect.SetKill()
		connect.SetDelay()
	}

	// Created message to send
	msm = &f.Message{
		To:     f.MulticastAddress,
		From:   id,
		Targ:   target,
		Data:   inf,
		Vector: copyVector,
		Delay:  delay,
	}

	fmt.Println("[SendGroupM] Send message from:  ", connect.GetId(), " to: ", f.MulticastAddress)
	// Creating red connection
	red, err = net.ResolveUDPAddr("udp", f.MulticastAddress)
	f.Error(err, "SendGroupM error ResolveUDPAddr connection \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "SendGroupM error DialUDP connection \n")
	defer connection.Close()

	// Create  chan type interface
	fmt.Println("[SendGroupM] Create chan and call ReceiveM  ")
	// packChan := make(chan interface{},10)
	var bufferAck []f.Ack

	// Send msm to ip multicast 3 times
	for i := 0; i < 3; i++ {
		log.Println("[RGM]    Dentro en el for: ", i)
		encoder = gob.NewEncoder(&buffer)
		err = encoder.Encode(msm)
		f.Error(err, "SendGroupM encoder error \n")
		_, err = connection.Write(buffer.Bytes())
		f.Error(err, "Error al recibir el msm")

		time.Sleep(500 * time.Millisecond)
	}

	// for _, v := range connect.GetIds() {
	// 	go ReceiveM(packChan, connect.GetId())
	// 	pack := <-packChan
	// 	bufferAck = append(bufferAck, pack.(f.Ack))
	// 	log.Println("[RGM]    Dentro en el for: ", v)
	// }

	_, pendCheck := f.CheckAcks(bufferAck, connect)

	// TODO Call Receive
	if len(pendCheck) > 0 {
		for _, v := range pendCheck {
			go SendM(msm, v)
		}
	}

	// TODO Sort vclock
	return err

}
