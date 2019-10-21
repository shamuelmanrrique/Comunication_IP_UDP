package functions

import (
	"fmt"
)

// ReceiveGroup snks
func ReceiveGroup(connect Connection, n int) error {
	// var sortMsms []Message
	bufferMsm := make(chan Message, len(connect.GetIds()))
	var err error
	vector := connect.GetVector()
	id := connect.GetId()

	go Receive(connect, bufferMsm)

	for {
		msm, ok := <-bufferMsm
		fmt.Println(msm.GetData())

		//TODO  -> Organizar elementos que me llegaron
		vector.Merge(msm.GetVector())

		if msm.GetFrom() == id {
			go SendGroup(connect)
		}

		if ok == false {
			break
		}

	}

	return err

}
