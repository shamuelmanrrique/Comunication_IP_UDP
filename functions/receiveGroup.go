package functions

import (
	"fmt"
)

// ReceiveGroup snks
func ReceiveGroup(connect Connection, n int) error {
	bufferMsm := make(chan Message, len(connect.GetIds()))
	var err error
	go Receive(connect, bufferMsm)

	for {
		fmt.Printf("_____________ESPERO MSM_______________ %s \n", connect.GetId())
		msm, ok := <-bufferMsm

		fmt.Printf("Mi ip: %s el msm diri: %s", connect.GetId(), msm.GetFrom())
		// Si el msm esta dirirgido a mi notifico a los demas
		if msm.GetFrom() == connect.GetId() {
			go SendGroup(connect)
		}

		//TODO  -> Organizar elementos que me llegaron
		if ok == false {
			break
		}
	}

	return err

}
