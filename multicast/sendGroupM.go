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

	id := connect.GetId()

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

	// Send msm to ip multicast 3 times
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

	log.Println("[SendGroupM] 77 VOY A RECIBIR ACK")
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		select {
		case pack, _ := <-chanAck:
			log.Println("[SendGroupM] bufferAck ", <-chanAck)

			if connect.GetId() != pack.GetOrigen() {
				bufferAck, ok = f.AddAcks(bufferAck, pack)
			}
			log.Println("[SendGroupM] buffer:", len(bufferAck), "lenght", len(connect.GetIds())-1)
			if len(bufferAck) == len(connect.GetIds())-1 {
				break
			}
		default:
			if len(bufferAck) == len(connect.GetIds())-1 {
				break
			}
		}

	}

	log.Println("[SendGroupM] CHEQUEOS LOS ACKS ")
	pendCheck, chec := f.CheckAcks(bufferAck, connect)

	// TODO Call Receive
	if !chec {
		go func() {
			for i := 0; i < 3; i++ {
				for _, v := range pendCheck {
					log.Println("[SendGroupM] MSM UNICAST TO ", v)
					go SendM(msm, v)
				}
				time.Sleep(200 * time.Millisecond)
			}
		}()

		log.Println("[SendGroupM] WAITING ACK LOST")

	readChannel:
		for {
			select {
			case pack := <-chanAck:
				log.Println("[SendGroupM] Adding ACK ", pack)
				if connect.GetId() != pack.GetOrigen() {
					bufferAck, ok = f.AddAcks(bufferAck, pack)
				}
			case <-time.After(3 * time.Second):
				log.Println("[SendGroupM] TIMEOUT ")
				break readChannel
			}
		}
	}

	pendCheck, chec = f.CheckAcks(bufferAck, connect)
	log.Println("[SendGroupM] CHEC ", chec, "array", pendCheck, "faltan", chec)
	if chec {
		log.Println("[SendGroupM] Communication error messsage whitout confirmation program finished ")
		return err
	}

	log.Println("[SendGroupM] |||||| Fin send Group |||| ", ok)

	return err

}
