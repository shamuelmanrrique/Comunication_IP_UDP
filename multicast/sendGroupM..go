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
	var ok bool
	var err error
	var bufferAck []f.Ack

	target := ""
	delay, _ := time.ParseDuration("0s")
	inf := "Me mataron"
	id := connect.GetId()

	// Update vClock
	vector := connect.GetVector()
	vector.Tick(id)
	connect.SetClock(vector)

	// Create copy of vector to send into the message
	copyVector := vector.Copy()

	// Check if it has a target
	log.Println("[SendGroupM] Check if it has a target")
	if len(connect.GetKill()) > 0 && len(connect.GetDelays()) > 0 {
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
	go func() {
		log.Println("[SendGroupM] FOR Send msm to multicast three times ")
		for i := 0; i < 3; i++ {
			log.Println("[SendGroupM] ENVIO Numero ", i)
			encoder = gob.NewEncoder(&buffer)
			err = encoder.Encode(msm)
			f.Error(err, "SendGroupM encoder error \n")
			_, err = connection.Write(buffer.Bytes())
			f.Error(err, "Error al enviar el msm")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	log.Println("[SendGroupM] 77 VOY A RECIBIR ACK")
	// dictAck := make(map[string]f.Ack)
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		log.Println("[SendGroupM] 81  FOR RECEIVE ACK for 3 seconds ")
		pack := <-chanAck
		if connect.GetId() != pack.GetOrigen() {
			// dictAck[pack.GetOrigen()] = pack
			bufferAck, ok = f.AddAcks(bufferAck, pack)
		}
	}

	log.Println("[SendGroupM] Salgo del FOR de recibir ACKS ")
	pendCheck, chec := f.CheckAcks(bufferAck, connect)

	// TODO Call Receive
	log.Println("[SendGroupM] IMPRIMO A VER SI FALTAN ACK ", chec, " Y LOS ACKS ", pendCheck)
	if !chec {
		//Necesito enviar tres veces
		log.Println("[SendGroupM] me faltan ACK los envio rirectamente")
		go func() {
			for i := 0; i < 3; i++ {
				for _, v := range pendCheck {
					log.Println("[SendGroupM] envio a ", v)
					go SendM(msm, v)
				}
				time.Sleep(200 * time.Millisecond)
			}
		}()

		deadline2 := time.Now().Add(3 * time.Second)
		for time.Now().Before(deadline2) {
			// for i := 0; i < nless; i++ {
			log.Println("[SendGroupM] FOR RECEIVE ACK for 3 seconds ")
			pack := <-chanAck
			if connect.GetId() != pack.GetOrigen() {
				// dictAck[pack.GetOrigen()] = pack
				bufferAck, ok = f.AddAcks(bufferAck, pack)
			}
		}
	}

	log.Print("[SendGroupM] communication error finished program ", ok)
	// _, ok := f.CheckAcks(bufferAck, connect)

	// if !ok {
	// 	log.Print("[SendGroupM] communication error finished program ")
	// 	return err
	// }

	// TODO Sort vclock
	return err

}
