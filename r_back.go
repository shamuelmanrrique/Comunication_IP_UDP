package functions

import (
	"fmt"
	"net"
)

// ReceiveGroup snks
func ReceiveGroups(connect Connection) error {
	var err error
	// var arrayMsms []Message
	var listener net.Listener
	var n = len(connect.GetIds())
	// vector := connect.GetVector()
	// id := connect.GetId()
	bufferMsm := make(chan Message, n)
	defer close(bufferMsm)

	listener, err = net.Listen("tcp", connect.GetPort())
	Error(err, "Listen Error")
	defer listener.Close()

	// var wg sync.WaitGroup
	// defer

	for i := 0; i < n; i++ {
		// go Receive(connect, bufferMsm, listener)
		// wg.Add(1)
		// go Receive(bufferMsm, listener, &wg)
		fmt.Printf("FOR:       %d", i)
		// go Receive(bufferMsm, listener)
	}

	// for range bufferMsm {
	for range bufferMsm {
		//
		// fmt.Printf("ESPERANDO BUFFER")
		msm, ok := <-bufferMsm

		fmt.Printf("recibi del canal ", msm)
		fmt.Println(ok)
		// if !ok {
		// 	break
		// } else {
		// 	fmt.Printf("FOR:       %d")
		// 	vector.Merge(msm.GetVector())

		// 	// Guardo el msm en un array de msm
		// 	arrayMsms = append(arrayMsms, msm)

		// 	// fmt.Printf("ELSE \n")

		// 	// fmt.Println(msm.GetFrom())
		// 	// fmt.Println(id)

		// 	// fmt.Println(msm.GetFrom() == id)
		// 	if msm.GetFrom() == id {
		// 		fmt.Println("SEND")
		// 		go SendGroup(connect)
		// 	}

		// 	// //TODO  -> Organizar elementos que me llegaron
		// 	// vector.Merge(msm.GetVector())

		// 	// // Guardo el msm en un array de msm
		// 	// arrayMsms = append(arrayMsms, msm)

		// 	// Ordeno el arreglo de msm
		// 	// sort.SliceStable(arrayMsms, func(i, j int) bool {
		// 	// 	return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
		// 	// })

		// 	// // Meto el vector act
		// 	// connect.GetVector().Merge(vector)
		// 	// if msm.GetFrom() == id {
		// 	// 	go SendGroup(connect)
		// 	// }

		// }

	}

	// // Ordeno el arreglo de msm
	// sort.SliceStable(arrayMsms, func(i, j int) bool {
	// 	return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	// })

	// Meto el vector act
	// connect.GetVector().Merge(vector)

	// fmt.Println("****************************")
	// for _, v := range arrayMsms {
	// 	m := v.GetVector()
	// 	fmt.Println(m.PrintVC)
	// }

	return err

}
