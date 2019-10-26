package comunicacionCausal

import (
	"fmt"
	"net"
	f "practice1/functions"
)

// ReceiveGroup SLMA
func ReceiveGroup(connect *f.Conn) error {
	var err error
	var listener net.Listener
	var n = len(connect.GetIds())
	vector := connect.GetVector()
	id := connect.GetId()

	bufferMsm := make(chan f.Message)
	defer close(bufferMsm)

	listener, err = net.Listen("tcp", connect.GetPort())
	f.Error(err, "Listen Error")
	defer listener.Close()

	// fmt.Printf("EL VALOR N:       %d \n", n)
	for i := 0; i < n; i++ {
		// fmt.Printf("FOR:       %d \n", i)
		go Receive(bufferMsm, listener, id)

		msm, _ := <-bufferMsm

		// fmt.Println("RELOJ RG:  ", vector)
		vector.Tick(id)
		// fmt.Println("RELOJ RG and TICK:  ", vector)
		vector.Merge(msm.GetVector())
		// fmt.Println("RELOJ MERGE:  ", vector)
		connect.SetClock(vector)
		// fmt.Println("targets en receive group:  ", connect.GetTarget(0))

		fmt.Println("***********Me quede en el if: ", id, " TO: ", msm.GetTo())
		if id == msm.GetTo() {
			// if len(connect.GetKill()) > 0 || {
			// fmt.Println("Contenido del mensaje recibido:", msm)
			go SendGroup(connect)
			// }
		}

		fmt.Println("++>>>RECIBIDO MSM DE : ", msm.GetFrom(), "Vectror: ", msm.GetVector())

	}

	// // // Guardo el msm en un array de msm
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

	fmt.Println("--------------------------")
	return err

}
