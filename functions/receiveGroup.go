package functions

import (
	"fmt"
	"net"
)

func ReceiveGroup(connect Connection) error {
	var err error
	var listener net.Listener
	// var wg sync.WaitGroup
	var n = len(connect.GetIds())
	id := connect.GetId()
	// bufferMsm := make(chan Message, n)
	bufferMsm := make(chan Message)
	// defer close(bufferMsm)

	listener, err = net.Listen("tcp", connect.GetPort())
	Error(err, "Listen Error")
	defer listener.Close()

	fmt.Printf("EL VALOR N:       %d \n", n)
	for i := 0; i < n; i++ {
		// wg.Add(1)
		fmt.Printf("FOR:       %d \n", i)
		// go Receive(bufferMsm, listener, &wg)
		go Receive(bufferMsm, listener)

		msm, ok := <-bufferMsm
		fmt.Printf("[RECEIVE] %s --> \n", msm)
		fmt.Printf("recibi del canal \n ", msm)
		fmt.Println(ok)
		if id == msm.To {
			go SendGroup(connect)
		}
	}

	// for i:+ n {
	// 	fmt.Printf("------------->")
	// 	msm, ok := <-bufferMsm
	// 	fmt.Printf("[RECEIVE] %s -->", msm)
	// 	fmt.Printf("recibi del canal \n ", msm)
	// 	fmt.Println(ok)
	// 	if id == msm.To {
	// 		go SendGroup(connect)
	// 	}
	// }

	// wg.Wait()

	fmt.Printf("--------------------------")
	return err

}
