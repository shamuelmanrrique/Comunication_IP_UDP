package functions

import (
	"fmt"
)

// buscar elementos o que te retornes los elementos
// de la session

// change
// chsh /bin/bash

func SendGroup(connect Connection) error {
	var err error
	id := connect.GetId()

	targets := connect.GetKill()

	t := len(targets)
	if t > 0 {

		fmt.Println("000000000000000000000000  \n")
		fmt.Println(targets)
		temp := targets[:t-1]
		fmt.Printf("1111111111111111111111111 ")
		connect.Conn(temp)
		fmt.Println(connect.GetKill())
		// Actualizo reloj
		vector := connect.GetVector()
		// u, _ := vector.FindTicks(id)
		// h := int(u)
		vector.Tick(id)

		// copiar el vector y enviar una copia
		copyVector := vector.Copy()

		var msm Message = Message{
			To:   id,
			From: connect.GetTarget(1), //TODO PROBLEMAS
			Data: " -- disparo --> ",
			// Vector: vector,
			Vector: copyVector,
		}

		// m := " -- disparo --> "
		// m := id + " -- disparo --> " + connect.GetTarget(h)
		// m := "[SEND => " + id + " mato a " + connect.GetTarget(h) + "]"

		// fmt.Printf("FOR  ------ %s ", connect.GetIds())

		for _, v := range connect.GetIds() {
			// fmt.Println("\n ++++++++++++++++++++")
			// fmt.Printf("valores   de env %d:", v)
			// if v != id {
			// Aplico delay en el envio
			msm.Ignor = v
			// delay := time.Duration(connect.GetDelay(i))
			// time.Sleep(delay * time.Millisecond)

			fmt.Printf("[GOSSIP] %s --> %s \n", id, v)
			go Send(msm)
			// }

		}
	}
	return err

}
