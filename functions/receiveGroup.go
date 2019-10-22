package functions

import (
	"fmt"
	"net"
	v "practice1/vclock"
	"sort"
)

// ReceiveGroup snks
func ReceiveGroup(connect Connection, n int) error {
	var listener net.Listener
	var arrayMsms []Message
	bufferMsm := make(chan Message)
	var err error
	id := connect.GetId()
	vector := connect.GetVector()

	listener, err = net.Listen("tcp", connect.GetPort())
	Error(err, "Listen Error")
	defer listener.Close()

	// connect.GetListe =

	go Receive(connect, bufferMsm, listener)

	for {
		// fmt.Printf("ESPERANDO BUFFER")
		msm, ok := <-bufferMsm
		// fmt.Printf("recibi del canal %")
		// fmt.Println(ok)
		if !ok {
			break
		} else {
			// fmt.Printf("ELSE \n")
			// fmt.Println(vector)
			vector.Merge(msm.GetVector())
			// fmt.Println(msm.GetVector())

			// Guardo el msm en un array de msm
			arrayMsms = append(arrayMsms, msm)
			// fmt.Println(msm.GetFrom())
			// fmt.Println(id)

			// fmt.Println(msm.GetFrom() == id)
			if msm.GetFrom() == id {
				fmt.Println("SEND")


				go SendGroup(connect)
			}

			// //TODO  -> Organizar elementos que me llegaron
			// vector.Merge(msm.GetVector())

			// // Guardo el msm en un array de msm
			// arrayMsms = append(arrayMsms, msm)

			// // Ordeno el arreglo de msm
			// sort.SliceStable(arrayMsms, func(i, j int) bool {
			// 	return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
			// })

			// // Meto el vector act
			// connect.GetVector().Merge(vector)
			// if msm.GetFrom() == id {
			// 	go SendGroup(connect)
			// }

		}

	}

	close(bufferMsm)
	// Ordeno el arreglo de msm
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	// Meto el vector act
	connect.GetVector().Merge(vector)

	for _, v := range arrayMsms {
		fmt.Println(v)
	}

	return err

}
