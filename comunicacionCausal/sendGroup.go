package comunicacionCausal

import (
	"fmt"
	f "practice1/functions"
	"time"
)

func SendGroup(connect f.Conn) error {
	var err error
	id := connect.GetId()
	targets := connect.GetKill()

	t := len(targets)
	if t > 0 {

		// Obtengo target and delay y los elimino
		target := connect.GetTarget(1)
		delay := connect.GetDelay(1)
		connect.SetKill(connect.GetKill()[:t-1]) 
		connect.Delays = connect.GetDelays()[:t-1]
		fmt.Println("SENDGROUP  ", connect.GetKill())

		// Actualizo el reloj
		vector := connect.GetVector()
		vector.Tick(id)

		// Copyo el vector
		copyVector := vector.Copy()

		var msm f.Message = f.Message{
			To:     id,
			From:   target, //TODO PROBLEMAS
			Data:   " -- disparo --> ",
			Vector: copyVector,
		}

		// Envio el msm a todos
		for _, v := range connect.GetIds() {
			if v != target {
				// Aplico delay y envio
				time.Sleep(delay * time.Millisecond)
				fmt.Printf("[GOSSIP] %s --> %s \n", id, v)
				msm.Ignor = v
				go Send(msm)
			}
		}
	}
	return err

}
