package communication

import (
	"log"
	"net"
	f "practice1/functions"
	v "practice1/vclock"
	"sort"
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

	// a := n - len(connect.GetKill())
	// switch connect.GetPort() {
	// case ":5001":
	// 	n = 1
	// case ":5002":
	// 	n = 1
	// case ":5003":
	// 	n = 2
	// }

	for i := 0; i < n; i++ {
		log.Println("[RG] EL VALOR N:       ", n, " El valor de i :", i)
		// i := 0
		// for {
		log.Println("[RG] FOR RECEIVE GROUP:      ", i)

		log.Println("[RG] LLAMO A Receive")
		go Receive(bufferMsm, listener, id)

		log.Println("[RG]________________________")
		msm, ok := <-bufferMsm
		log.Println("[RG]+++++++++++++++++++++++++")

		log.Println("[RG] VALOR DE OK: ", ok)
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

			log.Println("[RG] IF RG >>>: ", id, " TO: ", msm.GetTo())
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

	log.Println(" [RG] Estoy fuera del FOR ")
	// // Ordeno el arreglo de msm
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	// Meto el vector act
	log.Println("|||||||||||||||-----------------------")
	return err

}
