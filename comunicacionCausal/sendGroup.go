package comunicacionCausal

import (
	"fmt"
	f "practice1/functions"
	"time"
)

// SendGroup dda
func SendGroup(connect *f.Conn) error {
	var err error
	id := connect.GetId()
	// t := len(connect.GetKill())

	// Actualizo el reloj
	vector := connect.GetVector()
	// fmt.Println("Target ANTES del cambio  ", connect.GetKill())

	// if t > 0 {
	target := connect.GetTarget(0)
	delay := connect.GetDelay(0)
	connect.SetKill()
	// fmt.Println("Target despues del cambio  ", connect.GetKill())

	// Incremento el reloj
	vector.Tick(id)
	connect.SetClock(vector)
	// TODO CREATE SNAPSHOP RELOJ []VCLOCK

	// Copio el vector
	// copyVector := vector.Copy()

	var msm f.Message = f.Message{
		From: id, //TODO PROBLEMAS
		To:   target,
		// Data:   " -- disparo --> ",
		Vector: vector,
	}
	// fmt.Println("target en send group:  ", target)

	// Envio el msm a todos
	for _, v := range connect.GetIds() {
		if v != id {
			// msm.To = ""
			// msm.Data = " -- recibido --> "
			if v != target {
				// Aplico delay y envio
				// fmt.Println("Ejecutando delay de ", delay, "milisegundos")
				time.Sleep(delay)
			}
			go Send(v, msm, id)
			fmt.Println("Envio MSM ", id, "+++++>>", v)

		}
	}

	// }
	return err

}
