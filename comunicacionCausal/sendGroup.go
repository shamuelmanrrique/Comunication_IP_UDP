package comunicacionCausal

import (
	"fmt"
	f "practice1/functions"
)

// SendGroup
func SendGroup(connect *f.Conn) error {
	var err error
	id := connect.GetId()

	fmt.Println("Target ANTES del cambio  ", connect.GetKill())
	t := len(connect.GetKill())
	if t > 0 {
		target := connect.GetTarget(0)
		connect.SetKill()
		fmt.Println("Target despues del cambio  ", connect.GetKill())

		// Obtengo target and delay y los elimino
		// target := connect.GetTarget(0)
		// delay := connect.GetDelay(0)
		// connect.SetKill(connect.GetKill()[:t-1])
		// connect.Delays = connect.GetDelays()[:t-1]
		// fmt.Println("SENDGROUP  ", connect.GetKill())
		// fmt.Println("Kill restantes", connect.GetKill())
		// fmt.Println("Delays restantes", connect.GetDelays())

		// Actualizo el reloj
		vector := connect.GetVector()
		vector.Tick(id)

		// Copyo el vector
		copyVector := vector.Copy()

		var msm f.Message = f.Message{
			From:   id, //TODO PROBLEMAS
			To:     target,
			Data:   " -- disparo --> ",
			Vector: copyVector,
		}
		fmt.Println("target en send group:  ", target)

		// Envio el msm a todos
		// fmt.Println("Ids a enviar", connect.GetIds())
		for _, v := range connect.GetIds() {
			if v != id {
				// Aplico delay y envio
				// fmt.Println("Ejecutando delay de ", delay, "milisegundos")
				// time.Sleep(delay * time.Millisecond)
				msm.To = ""
				msm.Data = " -- recibido --> "
				go Send(v, msm, id)
			}
			// fmt.Printf("[GOSSIP] %s --> %s \n", id, v)

		}
	}
	return err

}
