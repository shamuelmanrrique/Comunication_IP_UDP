package communication

import (
	"log"
	"net"
	f "practice1/functions"
	v "practice1/vclock"
	"sort"
	"time"
)

// ReceiveGroup SLMA
func ReceiveGroup(connect *f.Conn) error {
	var err error
	var listener net.Listener
	var arrayMsms []f.Message
	n := connect.GetAccept()
	vector := connect.GetVector()
	id := connect.GetId()

	bufferMsm := make(chan f.Message)
	defer close(bufferMsm)

	listener, err = net.Listen("tcp", connect.GetPort())
	f.Error(err, "Listen Error")
	defer listener.Close()

	for i := 0; i < n; i++ {
		go Receive(bufferMsm, listener, id)

		msm, ok := <-bufferMsm
		if ok {
			// RECIBO y sumo 1 al vector
			vector.Tick(id)
			connect.SetClock(vector)
			vector.Merge(msm.GetVector())
			connect.SetClock(vector)

			if id == msm.GetTarg() {
				n = n - 1
				go SendGroup(connect)
			}

			// Guardo el msm en un array de msm
			arrayMsms = append(arrayMsms, msm)

		} else {
			break
		}
	}

	<-time.After(time.Second * 6)

	// Ordeno el arreglo de msm
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	f.DistUnic("Output Message")
	for _, m := range arrayMsms {
		if m.GetTarg() != "" {
			log.Println("[Message] -->", m.GetFrom(), m.GetData(), m.GetTarg(), "|||| Vector:", m.GetVector())
		} else {
			log.Println("[Message] -->", m.GetData(), m.GetFrom(), "|||| Vector:", m.GetVector())
		}
	}

	return err

}
