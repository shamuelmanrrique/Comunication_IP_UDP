package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
	"time"
)

// SendGroupM send message to ip multicast and wait ack
func SendGroupM(chanAck chan f.Ack, connect *f.Conn) error {
	var red *net.UDPAddr
	var connection *net.UDPConn
	var encoder *gob.Encoder
	var buffer bytes.Buffer
	var err error
	var bufferAck []f.Ack

	target := ""
	delay, _ := time.ParseDuration("0s")
	inf := "Me mataron"
	id := connect.GetId()

	// log.Println("[SendGroupM] Updating vclock and copy it")
	// Update vClock
	vector := connect.GetVector()
	vector.Tick(id)
	connect.SetClock(vector)

	// Create copy of vector to send into the message
	copyVector := vector.Copy()

	// Check if I have a target
	if len(connect.GetKill()) > 0 && len(connect.GetDelays()) > 0 {
		// log.Println("[SendGroupM] I got a target")
		target = connect.GetTarget(0)
		delay = connect.GetDelay(0)
		inf = "He disparado"
		connect.SetKill()
		connect.SetDelay()
	}

	// Created message to send
	msm := &f.Message{
		To:     f.MulticastAddress,
		From:   id,
		Targ:   target,
		Data:   inf,
		Vector: copyVector,
		Delay:  delay,
	}

	// Creating red connection
	red, err = net.ResolveUDPAddr("udp", f.MulticastAddress)
	f.Error(err, "SendGroupM error ResolveUDPAddr connection \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "SendGroupM error DialUDP connection \n")
	defer connection.Close()

	// Send msm to ip multicast 3 times
	// log.Println("[SendGroupM] FOR Send msm to multicast three times ")
	go func() {
		for i := 0; i < 3; i++ {
			log.Println("[SendGroupM] Dentro del FOR multicast ", i)
			encoder = gob.NewEncoder(&buffer)
			err = encoder.Encode(msm)
			f.Error(err, "SendGroupM encoder error \n")
			_, err = connection.Write(buffer.Bytes())
			f.Error(err, "Error al enviar el msm")
			time.Sleep(50 * time.Millisecond)
		}
	}()

	log.Println("[SendGroupM] FOR RECEIVE ACK ")
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		// select {
		// case pack := <-chanAck:
		pack := <-chanAck
		bufferAck = append(bufferAck, pack)
		log.Println("[SendGroupM] REcibo ACK ", pack)
		// case <-time.After(3 * time.Second):
		// 	log.Println("[SendGroupM] time out and break for ")
		// 	break
		// }

	}

	log.Println("[SendGroupM]=============== ")
	// log.Println("[SendGroupM] check corregtitud of Ack ")
	pendCheck, _ := f.CheckAcks(bufferAck, connect)

	// TODO Call Receive
	nless := len(pendCheck)
	if nless > 0 {
		log.Println("[SendGroupM] me faltan ACK los envio rirectamente")
		for _, v := range pendCheck {
			log.Println("[SendGroupM] me faltan ACK los envio rirectamente")
			go SendM(msm, v)
		}
	}

	for i := 0; i < nless; i++ {
		pack := <-chanAck
		bufferAck = append(bufferAck, pack)
		log.Println("[SendGroupM] recibo el canal y lo meto en el arreglo de acks ")

	}
	_, ok := f.CheckAcks(bufferAck, connect)

	if !ok {
		log.Print("[SendGroupM] communication error finished program ")
		return err
	}

	// TODO Sort vclock
	return err

}
