package functions

import (
	"fmt"
)

// ReceiveGroup snks
func ReceiveGroup(connect Connection, n int) error {
	// bufferMsm := make(chan Message)
	bufferMsm := make(chan Message, len(connect.GetIds()))
	var err error
	go Receive(connect, bufferMsm)

	for {
		// fmt.Println(len(bufferMsm))
		// fmt.Printf("_____________ESPERO ReceiveGroup_______________ %s \n", connect.GetId())
		// fmt.Println(connect.GetId())
		// go Receive(connect, bufferMsm)
		msm, ok := <-bufferMsm
		// k := msm.GetFrom() == connect.GetId()
		// fmt.Printfm("Mi ip: %s el msm diri: %s  \n", connect.GetId(), msm.GetFrom())

		// fmt.Println(msm.GetFrom())
		// fmt.Println(k)
		// Si el msm esta dirirgido a mi notifico a los demas
		// if msm.GetFrom() == connect.GetId() {
		// 	fmt.Println("msm.GetFrom()")
		// 	go SendGroup(connect)
		// 	// go SendGroup(connect)
		// }

		fmt.Println(msm.GetData())
		//TODO  -> Organizar elementos que me llegaron
		if ok == false {
			break
		}
	}

	return err

}
