package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "sd_paxos/src/functions"
	"strings"
	"time"
)

/*
-----------------------------------------------------------------
METODO: SendGroup
RECIBE: channel "f.Ack", pointer to connection "f.Conn"
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function TO send message to ip multicast and wait ack
-----------------------------------------------------------------
*/
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

	println("SendGroupM ===========")
	// Update vClock and make a copy to send that in message
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

	// Creating red connection to MulticastAddress
	red, err = net.ResolveUDPAddr("udp", f.MulticastAddress)
	f.Error(err, "SendGroupM error ResolveUDPAddr connection \n")

	// Making dial connection to ip address
	connection, err = net.DialUDP("udp", nil, red)
	f.Error(err, "SendGroupM error DialUDP connection \n")
	defer connection.Close()

	// Send message to MulticastAddress
	go func() {
		// Send the same message three times
		for i := 0; i < 3; i++ {
			encoder = gob.NewEncoder(&buffer)
			err = encoder.Encode(&msm)
			f.Error(err, "SendGroupM encoder error \n")
			_, err = connection.Write(buffer.Bytes())
			// Sleep between every delivery
			println("SendGroupM =======message", i, msm.GetTo())
			time.Sleep(200 * time.Millisecond)
		}
	}()

	println(strings.Join(ids[:], "\n\n"))
	// Delete it IP from ids to doesn't wait for its ack
	ackWait := f.Remove(ids, id)
	println(strings.Join(ackWait[:], "\n\n"))
	aux := true

readAck:
	// Tag to stay waiting for ACK messages until all arrive
	for {
		select {
		case pack := <-chanAck:
			println("readAck -------------", pack.GetOrigen)
			// Adding ACK to ACK array
			if id != pack.GetOrigen() {
				bufferAck, ok = f.AddAcks(bufferAck, pack)
			}
			// If already have all ack messsage break loop
			if len(bufferAck) == n {
				break readAck
			}

		// After timeout break loop
		case <-time.After(10 * time.Second):
			println("case <-time.After(10 * time.Second)")
			break readAck
		}
	}

	// Checking if it has all the ack messages
	ackWait, ok = f.CheckAcks(ackWait, bufferAck)
	println(ackWait, ok)

	// If lost at least one ack then send direct message to addres
	if !ok && aux {
		go func() {
			for i := 0; i < 3; i++ {
				// Send the same message three times
				for _, v := range ackWait {
					println("SEND MESSAGE 1/1")
					go SendM(msm, v)
				}
				time.Sleep(200 * time.Millisecond)
			}
		}()

		if aux {
			aux = false
			// Goto receive missing ACK
			goto readAck
		}
	}

	// Communication error at least one messsage whithout confirmation
	if !ok {
		log.Println("[SendGroupM] Communication error messsage whithout confirmation program finished ")
		return err
	}

	return err
}
