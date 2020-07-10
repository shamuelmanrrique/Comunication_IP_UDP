package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "sd_paxos/functions"
	"sort"

	v "sd_paxos/vclock"
	"time"
)

// ReceiveGroupM  das
func ReceiveGroupM(chanMess chan f.Message, chanAck chan f.Ack, connect *f.Conn) error {
	var err error
	var ok bool
	var decode *gob.Decoder
	var listener *net.UDPConn
	var arrayMsms []f.Message
	n := len(connect.GetIds()) - 1
	vector := connect.GetVector()
	id := connect.GetId()

	// Open up a connection
	addr, _ := net.ResolveUDPAddr("udp", f.MulticastAddress)
	listener, err = net.ListenMulticastUDP("udp", nil, addr)
	f.Error(err, "ReceiveGroupM error ListenPacket")
	defer listener.Close()

	var msmMult []f.Message
	m := make(chan f.Message)
	defer close(m)

	go func() {
		deadline := time.Now().Add(40 * time.Second)
		for time.Now().Before(deadline) {
			var msm f.Message
			listener.SetReadBuffer(f.MaxBufferSize)
			buffer := make([]byte, f.MaxBufferSize)
			listener.ReadFromUDP(buffer)
			dataBuffer := bytes.NewBuffer(buffer)
			decode = gob.NewDecoder(dataBuffer)
			err = decode.Decode(&msm)
			if err != nil {
				break
			}

			if msm.GetFrom() != id {
				msmMult, ok, _ = f.CheckMsm(msmMult, msm)
				if !ok {
					log.Println(" RECEIVE MULTICAST-->: from ", msm.GetFrom(), " to ", msm.GetTo(), "  || OBJ: ", msm.GetTarg(),
						"\n                                 Vector: ", msm.GetVector())
					m <- msm
				}
			}
		}

	receiveChannel:
		for {
			select {
			case msm, ok := <-chanMess:
				if msm.GetFrom() != id {
					msmMult, ok, _ = f.CheckMsm(msmMult, msm)
					if !ok {
						log.Println(" RECEIVE UNICAST-->: from ", msm.GetFrom(), " to ", msm.GetTo(), "  || OBJ: ", msm.GetTarg(),
							"\n                                 Vector: ", msm.GetVector())
						m <- msm
					}
				}
			case <-time.After(time.Second * 40):
				break receiveChannel
			}
		}

	}()

readMessage:
	for {
		var messag f.Message
		select {
		case messag, ok = <-m:
			go func() {
				ackID := &f.Ack{
					Origen: connect.GetId(),
					Code:   connect.GetId() + "," + messag.GetFrom(),
				}

				// Aplico Delay
				if messag.GetTarg() != id {
					delay := messag.GetDelay()
					// log.Println("Delay: ", delay)
					time.Sleep(delay)
				}

				SendM(ackID, messag.GetFrom())
				vector.Tick(id)
				connect.SetClock(vector)
				vector.Merge(messag.GetVector())
				connect.SetClock(vector)

				if messag.GetTarg() == id {
					n--
					go SendGroupM(chanAck, connect)
				}
				arrayMsms = append(arrayMsms, messag)
			}()
		case <-time.After(25 * time.Second):
			break readMessage
		}
	}

	<-time.After(time.Second * 7)

	// Sort vector Array
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	f.DistUnic("Output Message")
	for _, men := range arrayMsms {
		if men.GetTarg() != "" {
			log.Println("[Message] -->", men.GetFrom(), men.GetData(), men.GetTarg())
		} else {
			log.Println("[Message] -->", men.GetFrom(), men.GetData())
		}
	}

	return err
}
