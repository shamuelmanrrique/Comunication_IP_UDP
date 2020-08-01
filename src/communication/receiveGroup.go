package communication

import (
	"log"
	"net"
	f "sd_paxos/src/functions"
	v "sd_paxos/src/vclock"
	"sort"
	"time"
)

/*
-----------------------------------------------------------------
METODO: ReceiveGroup
RECIBE:  Conn struct
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: Management of incoming messages
-----------------------------------------------------------------
*/
func ReceiveGroup(connect *f.Conn) error {
	// defining variables
	var err error
	var listener net.Listener
	var arrayMsms []f.Message
	n := connect.GetAccept()
	vector := connect.GetVector()
	id := connect.GetId()

	// Creating buffer of message type
	bufferMsm := make(chan f.Message)
	defer close(bufferMsm)

	// Creating connection TCP with especific port. 
	listener, err = net.Listen("tcp", connect.GetPort())
	f.Error(err, "Listen Error")
	defer listener.Close()

	// Loop to receive the determined n messages
	for i := 0; i < n; i++ {
		go Receive(bufferMsm, listener)

		// Getting incoming message from channel 
		msm, ok := <-bufferMsm
		if ok {
			// Procesing message and update vector add one event
			vector.Tick(id)
			connect.SetClock(vector)
			vector.Merge(msm.GetVector())
			connect.SetClock(vector)

			// If this machine is the target then send message to all 
			if id == msm.GetTarg() {
				n = n - 1
				go SendGroup(connect)
			}

			// Save message in array of incoming message 
			arrayMsms = append(arrayMsms, msm)
		
		// Id somenthing was wrong break loop
		} else {
			break
		}
	}

	<-time.After(time.Second * 25)

	// Sort message array by vector clock
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	f.DistUnic("Output Message")
	// Print message in order 
	for _, m := range arrayMsms {
		if m.GetTarg() != "" {
			log.Println("[Message] -->", m.GetFrom(), m.GetData(), m.GetTarg(), "|||| Vector:", m.GetVector())
		} else {
			log.Println("[Message] -->", m.GetFrom(), m.GetData(), "|||| Vector:", m.GetVector())
		}
	}

	return err

}
