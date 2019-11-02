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
		// log.Println("[RG] EL VALOR N:       ", n, " El valor de i :", i)
		// i := 0
		// for {
		// log.Println("[RG] FOR RECEIVE GROUP:      ", i)

		// log.Println("[RG] LLAMO A Receive")
		go Receive(bufferMsm, listener, id)

		// log.Println("[RG]________________________")
		msm, ok := <-bufferMsm
		// log.Println("[RG]+++++++++++++++++++++++++")

		// log.Println("[RG] VALOR DE OK: ", ok)
		if ok {
			// RECIBO y sumo 1 al vector
			vector.Tick(id)
			// SEt la nueva actualizacion de recepcion
			connect.SetClock(vector)
			// Uno los relojes
			vector.Merge(msm.GetVector())
			// connect.GetVector().Merge(vector)
			// Seteo nuevamente el reloj
			connect.SetClock(vector)

			// log.Println("[RG] IF RG >>>: ", id, " TO: ", msm.GetTo())
			if id == msm.GetTarg() {
				n = n - 1
				log.Println("[RG] Soy el target llamo a SG ")
				go SendGroup(connect)
			}

			// Guardo el msm en un array de msm
			arrayMsms = append(arrayMsms, msm)

		} else {
			log.Println("[RG] Estoy ELSE ")
			break
		}
		// i = i + 1
	}

	// Ordeno el arreglo de msm
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	// log.Println("|||||||||||||||||||||||||||||||||||||||||||")
	<-time.After(time.Second * 6)
	for _, m := range arrayMsms {
		log.Println("[Message] --> To: ", m.GetTo(), " From: ", m.GetFrom(), " inf: ", m.GetData())
	}

	return err

}
