package chandylamport

import (
	"log"
	f "practice1/functions"
	v "practice1/vclock"
	"sort"
	"time"
)

// ReceiveGroup SLMA
func ReceiveGroup(chanMessage chan f.Message, chanMarker chan f.Marker, connect *f.Conn) error {
	var err error
	var arrayMsms []f.Message
	n := connect.GetAccept()
	vector := connect.GetVector()
	id := connect.GetId()

	go Receive(chanMarker, chanMessage, connect.GetPort())

	for i := 0; i < n; i++ {
		msm, ok := <-chanMessage
		if ok {
			vector.Tick(id)
			connect.SetClock(vector)
			vector.Merge(msm.GetVector())
			connect.SetClock(vector)

			// 	receiveChannel:
			// case <-time.After(20 * time.Second) :
			// 	break receiveChannel
			// log.Println("[RG] IF RG >>>: ", id, " TO: ", msm.GetTo())

			if id == msm.GetTarg() {
				n = n - 1
				log.Println("[RG] Soy el target llamo a SG ")
				go SendGroup(chanMessage, chanMarker, connect)
			}

			// Guardo el msm en un array de msm
			arrayMsms = append(arrayMsms, msm)

		} else {
			log.Println("[RG] Estoy ELSE ")
			break
		}
	}

	// Ordeno el arreglo de msm
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	<-time.After(time.Second * 6)
	for _, m := range arrayMsms {
		// log.Println("[Message] --> To: ", m.GetTo(), " From: ", m.GetFrom(), " inf: ", m.GetData())
		if m.GetTarg() != "" {
			log.Println("[Message] -->", m.GetFrom(), m.GetData(), m.GetTarg())
		} else {
			log.Println("[Message] -->", m.GetFrom(), m.GetData())
		}
	}

	return err

}
