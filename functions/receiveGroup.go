package functions

import (
	v "practice1/vclock"
	"sort"
)

// ReceiveGroup snks
func ReceiveGroup(connect Connection, n int) error {
	var arrayMsms []Message
	bufferMsm := make(chan Message, len(connect.GetIds()))
	var err error
	id := connect.GetId()
	vector := connect.GetVector()

	go Receive(connect, bufferMsm)

	for {
		msm, ok := <-bufferMsm
		// fmt.Println(msm.GetData())

		//TODO  -> Organizar elementos que me llegaron
		vector.Merge(msm.GetVector())

		// Guardo el msm en un array de msm
		arrayMsms = append(arrayMsms, msm)

		// Ordeno el arreglo de msm
		sort.SliceStable(arrayMsms, func(i, j int) bool {
			return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
		})

		// Meto el vector act
		connect.GetVector().Merge(vector)
		if msm.GetFrom() == id {
			go SendGroup(connect)
		}

		if ok == false {
			break
		}
	}

	return err

}
