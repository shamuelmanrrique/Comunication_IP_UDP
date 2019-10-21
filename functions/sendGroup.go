package functions

import (
	"fmt"
	"time"
)

func SendGroup(connect Connection) error {
	var err error
	id := connect.GetId()

	// Actualizo reloj
	vector := connect.GetVector()
	u, _ := vector.FindTicks(id)
	h := int(u)
	vector.Tick(id)

	m := "[SEND => " + id + " mato a " + connect.GetTarget(h) + "]"
	var msm Message = Message{
		To:     id,
		From:   connect.GetTarget(h), //TODO PROBLEMAS
		Data:   m,
		Vector: vector,
	}

	for i, v := range connect.GetIds() {
		if v != id {
			// Aplico delay en el envio
			msm.Ignor = v
			delay := time.Duration(connect.GetDelay(i))
			time.Sleep(delay * time.Millisecond)
			fmt.Printf("[ENVIO] %s --> %s \n", id, v)
			go Send(msm)
		}

	}
	return err

}
