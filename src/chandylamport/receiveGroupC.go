package chandylamport

import (
	"log"
	f "sd_paxos/src/functions"
	v "sd_paxos/src/vclock"
	"sort"
	"time"
)

/*
-----------------------------------------------------------------
METODO: ReceiveGroupC
RECIBE:  canal de tipo string, canal de tipo f.Message, canal de tipo f.Marker, connection connect
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: Management of incoming messages
-----------------------------------------------------------------
*/
func ReceiveGroupC(chanPoint chan string, chanMessage chan f.Message, chanMarker chan f.Marker, connect *f.Conn) error {
	var err error
	var marker = &f.Marker{}
	var arrayMsms []f.Message
	var recordMsms []f.Message
	n := len(connect.GetIds())
	vector := connect.GetVector()
	id := connect.GetId()

	// Enabling to receive checkpoints
	go ReceiveC(chanPoint, chanMarker, chanMessage, connect.GetId())

receiveChannel:
	// Tag to stay waiting for messages until all arrive
	for {

		// Case to filter by channel incoming message 
		select {	
		// Normal Message 
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
			// Marker message 
			marker.SetRecoder(true)
			marker.SetHeader(arrayMsms)
			sendPoint(id, connect.GetIds())
			marker.SetCounter(n - 1)
			recordMsms = []f.Message{}

		// Receive CheckPoint
		case checkPoint := <-chanPoint:
			// Message to start recording 
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

		// After timeout break loop
		case <-time.After(time.Second * 10):
			break receiveChannel
		}
	}

	<-time.After(time.Second * 5)

	// Sort message array by vector clock
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	// Print Snapshot 
	marker.PrintMarker(id)

	f.DistUnic("Output Message")
	// Print message in order
	for _, m := range arrayMsms {
		if m.GetTarg() != "" {
			log.Println("[Message] -->", m.GetFrom(), m.GetData(), m.GetTarg())
		} else {
			log.Println("[Message] -->", m.GetFrom(), m.GetData())
		}
	}

	return err
}

/*
-----------------------------------------------------------------
METODO: sendPoint
RECIBE:  id string, array ids
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: Send checkpoint
-----------------------------------------------------------------
*/
func sendPoint(id string, ids []string) {
	for _, v := range ids {
		if v != id {
			point := id + "," + v
			// Send checkpoint
			go SendC(point, v)
			log.Println(" ++> SEND COUNT: to ", v, "  |||| Count: ", point)
		}
	}

}
