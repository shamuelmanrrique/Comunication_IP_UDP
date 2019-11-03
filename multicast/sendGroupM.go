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
	var msm f.Message
	var ok bool
	var err error
	var bufferAck []f.Ack
	ids := connect.GetIds()

	id := connect.GetId()
	n := len(connect.GetIds()) - 1

	// Update vClock and make a copy
	vector := connect.GetVector()
	vector.Tick(id)
	connect.SetClock(vector)
	copyVector := vector.Copy()

	// Check if it has a target
	if len(connect.GetKill()) > 0 && len(connect.GetDelays()) > 0 {
		msm = f.Message{
			To:     f.MulticastAddress,
			From:   id,
			Targ:   connect.GetTarget(0),
			Data:   "kill",
			Vector: copyVector,
			Delay:  connect.GetDelay(0),
		}
	} else {
		delay, _ := time.ParseDuration("0s")
		msm = f.Message{
			To:     f.MulticastAddress,
			From:   id,
			Targ:   "",
			Data:   "am dead",
			Vector: copyVector,
			Delay:  delay,
		}
	}

	// Creating red connection
	red, err = net.ResolveUDPAddr("udp", f.MulticastAddress)
	f.Error(err, "SendGroupM error ResolveUDPAddr connection \n")

	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "SendGroupM error DialUDP connection \n")
	defer connection.Close()

	go func() {
		for i := 0; i < 3; i++ {
			encoder = gob.NewEncoder(&buffer)
			err = encoder.Encode(&msm)
			f.Error(err, "SendGroupM encoder error \n")
			_, err = connection.Write(buffer.Bytes())
			time.Sleep(200 * time.Millisecond)
		}
		log.Println("[SendGroupM] ENVIO MULTICAST", msm)
	}()

	ackWait := f.Remove(ids, id)
	aux := true
readAck:
	for {
		select {
		case pack := <-chanAck:
			log.Println("[SendGroupM] Adding ACK ", pack)
			if id != pack.GetOrigen() {
				bufferAck, ok = f.AddAcks(bufferAck, pack)
			}
			if len(bufferAck) == n {
				log.Println("[SendGroupM] Salgo me llegaron todos los ACK ")
				break readAck
			}
		case <-time.After(4 * time.Second):
			log.Println("[SendGroupM] TIMEOUT readAck")
			break readAck
		}
	}

	ackWait, ok = f.CheckAcks(ackWait, bufferAck)
	log.Println("[SendGroupM] IMPRIMO LOS VALORES DE CHECKACKS ",  ackWait, ok)
	if !ok && aux {
		go func() {
			for i := 0; i < 3; i++ {
				for _, v := range ackWait {
					log.Println("[SendGroupM] MSM UNICAST TO ", v)
					go SendM(msm, v)
				}
				time.Sleep(200 * time.Millisecond)
			}
		}()

		if aux {
			aux = false
			goto readAck
		}

		// readAck2:
		// 	for {
		// 		select {
		// 		case pack := <-chanAck:
		// 			log.Println("[SendGroupM] Adding ACK 222", pack)
		// 			if id != pack.GetOrigen() {
		// 				bufferAck, ok = f.AddAcks(bufferAck, pack)
		// 			}
		// 			if len(bufferAck) == n {
		// 				break readAck2
		// 			}
		// 		case <-time.After(6 * time.Second):
		// 			log.Println("[SendGroupM] TIMEOUT 22")
		// 			break readAck2
		// 		}
		// 	}

		// ackWait, ok = f.CheckAcks(ackWait, bufferAck)
	}

	if !ok {
		log.Println("[SendGroupM] Communication error messsage whitout confirmation program finished ")
		return err
	}

	log.Println("[SendGroupM] |||||| Fin send Group |||| ")

	return err

}
