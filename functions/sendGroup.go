package functions

import (
	"fmt"
)

func SendGroup(connect Connection) error {
	var err error
	id := connect.GetId()

	// Actualizo reloj
	vector := connect.GetVector()
	u, _ := vector.FindTicks(id)
	h := int(u)
	vector.Tick(id)
	// copiar el vector y enviar una copia

	var msm Message = Message{
		To:     id,
		From:   connect.GetTarget(h), //TODO PROBLEMAS
		Data:   " -- disparo --> ",
		Vector: vector,
	}

	// m := " -- disparo --> "
	// m := id + " -- disparo --> " + connect.GetTarget(h)
	// m := "[SEND => " + id + " mato a " + connect.GetTarget(h) + "]"

	// fmt.Printf("FOR  ------ %s ", connect.GetIds())
	for _, v := range connect.GetIds() {
		// fmt.Println("\n ++++++++++++++++++++")
		// fmt.Printf("valores   de env %d:", v)
		if v != id {
			// Aplico delay en el envio
			msm.Ignor = v
			// delay := time.Duration(connect.GetDelay(i))
			// time.Sleep(delay * time.Millisecond)
			fmt.Printf("[GOSSIP] %s --> %s \n", id, v)
			go Send(msm)
		}

	}
	return err

}
