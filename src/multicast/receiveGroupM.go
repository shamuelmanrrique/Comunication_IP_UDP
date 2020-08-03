package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "sd_paxos/src/functions"
	"sort"

	v "sd_paxos/src/vclock"
	"time"
)

/*
-----------------------------------------------------------------
METODO: ReceiveGroupM
RECIBE: canal de tipo f.Message, un listener net.Listener
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function that receive message using UDP connection  and IP multicast
-----------------------------------------------------------------
*/
func ReceiveGroupM(chanMess chan f.Message, chanAck chan f.Ack, connect *f.Conn) error {
	var err error
	var ok bool
	var decode *gob.Decoder
	var listener *net.UDPConn
	var arrayMsms []f.Message
	n := len(connect.GetIds()) - 1
	vector := connect.GetVector()
	id := connect.GetId()

	// Set up and address to receive message
	addr, _ := net.ResolveUDPAddr("udp", f.MulticastAddress)
	listener, err = net.ListenMulticastUDP("udp", nil, addr)
	f.Error(err, "ReceiveGroupM error ListenPacket")
	defer listener.Close()

	// Creating message array to incoming message
	var msmMult []f.Message
	// Creating channel message
	m := make(chan f.Message)
	defer close(m)

	println("++++++++++++++++++> ReceiveGroupM DIRECCION", f.MulticastAddress)

	// Gorutine to receive UNICAST AND MULTICAST message
	go func() {
		// Defining timeout to wait MULTICAST menssage
		deadline := time.Now().Add(40 * time.Second)

		for time.Now().Before(deadline) {
			println("++++++++++++++++++> ReceiveGroupM for por por 40s")
			var msm f.Message
			// Set up to read message using buffer
			listener.SetReadBuffer(12000)
			buffer := make([]byte, 12000)
			listener.ReadFromUDP(buffer)

			// Reding from buffer and decoder message
			dataBuffer := bytes.NewBuffer(buffer)
			decode = gob.NewDecoder(dataBuffer)
			err = decode.Decode(&msm)
			println("|||||||||||||||||||||||||||", &msm)
			if err != nil {
				f.Error(err, "Receive error  \n")
				break
			}

			println("++++++> ReceiveGroupM llego sms de ", id, msm.GetFrom())

			// Checking message from other IP
			if msm.GetFrom() != id {
				println("NO SOY YO EL ORIGEN ", msm.GetFrom())
				// Validate message to doesn't add duplicates
				msmMult, ok, _ = f.CheckMsm(msmMult, msm)
				if !ok {
					log.Println(" RECEIVE MULTICAST-->: from ", msm.GetFrom(), " to ", msm.GetTo(), "  || OBJ: ", msm.GetTarg(),
						"\n                                 Vector: ", msm.GetVector())
					// Sending message to channel
					m <- msm
				}
			}
		}

		println("++++++> receiveChannel ")
	receiveChannel:
		// Tag to stay waiting for UNICAST messages until all arrive
		for {
			select {
			case msm, ok := <-chanMess:
				if msm.GetFrom() != id {
					// Validate message to doesn't add duplicates
					msmMult, ok, _ = f.CheckMsm(msmMult, msm)
					if !ok {
						log.Println(" RECEIVE UNICAST-->: from ", msm.GetFrom(), " to ", msm.GetTo(), "  || OBJ: ", msm.GetTarg(),
							"\n                                 Vector: ", msm.GetVector())
						// Sending message to channel
						m <- msm
					}
				}
			// After timeout break loop
			case <-time.After(time.Second * 40):
				break receiveChannel
			}
		}

	}()

readMessage:
	// Tag to stay waiting for messages
	for {
		var messag f.Message
		select {

		// Getting incoming message from channel
		case messag, ok = <-m:
			go func() {
				println("Recibi un mmsmsmssms")
				ackID := &f.Ack{
					Origen: connect.GetId(),
					Code:   connect.GetId() + "," + messag.GetFrom(),
				}

				// Appling manual delay to receive message
				if messag.GetTarg() != id {
					delay := messag.GetDelay()
					time.Sleep(delay)
				}

				// Send ACK confirmation
				SendM(ackID, messag.GetFrom())

				// Updating vclock
				vector.Tick(id)
				connect.SetClock(vector)
				vector.Merge(messag.GetVector())
				connect.SetClock(vector)

				// Checking if its the target
				if messag.GetTarg() == id {
					n--
					// Sending multicast message
					go SendGroupM(chanAck, connect)
				}

				// Adding message to array
				arrayMsms = append(arrayMsms, messag)
			}()

		// After timeout break loop
		case <-time.After(40 * time.Second):
			break readMessage
		}
	}

	<-time.After(time.Second * 7)

	// Sort message array by vector clock
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	f.DistUnic("Output Message")
	// Print message in order
	for _, men := range arrayMsms {
		if men.GetTarg() != "" {
			log.Println("[Message] -->", men.GetFrom(), men.GetData(), men.GetTarg())
		} else {
			log.Println("[Message] -->", men.GetFrom(), men.GetData())
		}
	}

	return err
}
