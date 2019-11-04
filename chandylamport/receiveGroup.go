package chandylamport

import (
	"fmt"
	"log"
	f "practice1/functions"
	v "practice1/vclock"
	"sort"
	"time"
)

// ReceiveGroup SLMA
func ReceiveGroup(chanPoint chan string, chanMessage chan f.Message, chanMarker chan f.Marker, connect *f.Conn) error {
	var err error
	var marker = &f.Marker{}
	var arrayMsms []f.Message
	var recordMsms []f.Message
	n := len(connect.GetIds())
	vector := connect.GetVector()
	id := connect.GetId()
	go Receive(chanPoint, chanMarker, chanMessage, connect.GetPort())

receiveChannel:
	for {
		select {
		case msm, ok := <-chanMessage:
			log.Println("[ReceiveGroup] REcibi un msm ")
			if ok {
				vector.Tick(id)
				connect.SetClock(vector)
				vector.Merge(msm.GetVector())
				connect.SetClock(vector)

				if id == msm.GetTarg() {
					log.Println("[ReceiveGroup] Soy el target llamo a SG ")
					go SendGroup(chanPoint, chanMessage, chanMarker, connect)
				}

				// Guardo el msm en un array de msm
				arrayMsms = append(arrayMsms, msm)
				recordMsms = append(recordMsms, msm)

			} else {
				log.Println("[RG] Estoy ELSE ")
				break receiveChannel
			}

		// Init Snapshot
		case (*marker) = <-chanMarker:
			log.Println("[ReceiveGroup] PRINT ARRAY ", arrayMsms)
			log.Println("[ReceiveGroup]____________ RECIBI INIT MARKER _______ ")
			marker.SetRecoder(true)
			marker.SetHeader(arrayMsms)
			sendPoint(id, connect.GetIds())
			marker.SetCounter(n - 1)
			recordMsms = []f.Message{}
			log.Println("[ReceiveGroup]  MARKER INIT: ", marker)

		// Init Recibo CheckPoint
		case checkPoint := <-chanPoint:
			log.Println("[ReceiveGroup] CHECKPOINT _______ ", checkPoint, "VALUE OF MASTER ", marker)
			if !marker.GetRecoder() && marker.GetCounter() == 0 {
				log.Println("[ReceiveGroup] INIT CHECKPOINT")
				marker.SetRecoder(true)
				marker.SetCounter(n - 1)
				marker.SetHeader(recordMsms)
				marker.SetCheckPoints(checkPoint)
				sendPoint(id, connect.GetIds())
				marker.SetCounter(n - 1)
				log.Println("[ReceiveGroup] IF CHECKPOINT", marker)
				recordMsms = []f.Message{}

			} else {
				if marker.GetCounter() == 0 {
					log.Println("[ReceiveGroup] IF RECEIVE ALL VALUES")
					marker.SetRecoder(false)
					marker.SetCheckPoints(checkPoint)

					// marker.SetChanString(checkPoint)
					marker.SetCounter(marker.GetCounter() - 1)
					// Termino ejecucion imprimiendo mis estados
				} else {
					log.Println("[ReceiveGroup] ELSE RECEIVE ALL VALUES")
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
	// Ordeno el arreglo de msm
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})
	
	fmt.Println("|||||||||||||||||||| END |||||||||||||||||||||||")
	marker.PrintMarker(id)

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
			// time.Sleep(time.Millisecond * 130)
			point := id + "," + v
			log.Println("[ReceiveGroup] -->", id, " pointcheck ", v)
			go Send(point, v)
		}
	}

}
