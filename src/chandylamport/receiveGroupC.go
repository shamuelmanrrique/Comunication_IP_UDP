package chandylamport

import (
	"log"
	f "sd_paxos/src/functions"
	v "sd_paxos/src/vclock"
	"sort"
	"time"
)

// ReceiveGroup SLMA
func ReceiveGroupC(chanPoint chan string, chanMessage chan f.Message, chanMarker chan f.Marker, connect *f.Conn) error {
	var err error
	var marker = &f.Marker{}
	var arrayMsms []f.Message
	var recordMsms []f.Message
	n := len(connect.GetIds())
	vector := connect.GetVector()
	id := connect.GetId()
	go ReceiveC(chanPoint, chanMarker, chanMessage, connect.GetId())

receiveChannel:
	for {
		select {
		case msm, ok := <-chanMessage:
			if ok {
				vector.Tick(id)
				connect.SetClock(vector)
				vector.Merge(msm.GetVector())
				connect.SetClock(vector)

				if id == msm.GetTarg() {
					go SendGroupC(chanPoint, chanMessage, chanMarker, connect)
				}

				arrayMsms = append(arrayMsms, msm)
				recordMsms = append(recordMsms, msm)

			} else {
				break receiveChannel
			}

		// Init Snapshot
		case (*marker) = <-chanMarker:
			marker.SetRecoder(true)
			marker.SetHeader(arrayMsms)
			sendPoint(id, connect.GetIds())
			marker.SetCounter(n - 1)
			recordMsms = []f.Message{}

		// Init Recibo CheckPoint
		case checkPoint := <-chanPoint:
			if !marker.GetRecoder() && marker.GetCounter() == 0 {
				marker.SetRecoder(true)
				marker.SetCounter(n - 1)
				marker.SetHeader(recordMsms)
				marker.SetCheckPoints(checkPoint)
				sendPoint(id, connect.GetIds())
				marker.SetCounter(n - 1)
				recordMsms = []f.Message{}

			} else {
				if marker.GetCounter() == 0 {
					marker.SetRecoder(false)
					marker.SetCheckPoints(checkPoint)

					marker.SetCounter(marker.GetCounter() - 1)
				} else {
					marker.SetChannel(recordMsms)
					marker.SetCheckPoints(checkPoint)
					marker.SetCounter(marker.GetCounter() - 1)
					recordMsms = []f.Message{}
				}
			}
		case <-time.After(time.Second * 10):
			break receiveChannel
		}
	}

	<-time.After(time.Second * 5)
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	marker.PrintMarker(id)
	f.DistUnic("Output Message")
	for _, m := range arrayMsms {
		if m.GetTarg() != "" {
			log.Println("[Message] -->", m.GetFrom(), m.GetData(), m.GetTarg())
		} else {
			log.Println("[Message] -->", m.GetFrom(), m.GetData())
		}
	}

	return err
}

func sendPoint(id string, ids []string) {
	for _, v := range ids {
		if v != id {
			point := id + "," + v
			go SendC(point, v)
			log.Println(" ++> SEND COUNT: to ", v, "  |||| Count: ", point)
		}
	}

}
