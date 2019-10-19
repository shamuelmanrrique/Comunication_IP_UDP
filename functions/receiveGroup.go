package functions

import "fmt"

func ReceiveGroup(connect Connection, n int) error {

	bufferMsm := make(chan Message, n)
	var err error
	id := connect.GetId()

	for i := range bufferMsm {
		fmt.Println(id)
		fmt.Println(i)
		// Send(connect, bufferMsm)

	}

	return err

}
