package communication

import (
	"log"
	f "practice1/functions"
	"time"
)

// SendGroup dda
func SendGroup(connect *f.Conn) error {
	var err error
	target := ""
	delay, _ := time.ParseDuration("0s")
	inf := "Me mataron"
	id := connect.GetId()
	// t := len(connect.GetKill())

	// Actualizo el reloj
	vector := connect.GetVector()

	if len(connect.GetKill()) > 0 && len(connect.GetDelays()) > 0 {
		target = connect.GetTarget(0)
		delay = connect.GetDelay(0)
		inf = "He disparado"
		connect.SetKill()
		connect.SetDelay()
	}

	// Incremento el reloj
	vector.Tick(id)
	connect.SetClock(vector)

	// TODO CREATE SNAPSHOP RELOJ []VCLOCK
	// Copio el vector
	copyVector := vector.Copy()

	// Envio el msm a todos
	for _, v := range connect.GetIds() {
		if v != id {

			var msm f.Message = f.Message{
				To:     v,
				From:   id,
				Targ:   target,
				Data:   inf,
				Vector: copyVector,
			}

			if v != target {
				// Aplico delay y envio
				log.Println("[SG] Ejecutando delay de ", delay, "milisegundos")
				time.Sleep(delay)
			}

			log.Println("[SG] LLAMO A SEND", msm)
			go Send(v, msm, id)
			// log.Println("[SG] Envio MSM ", id, "+++++>>", v)

		}
	}

	return err

}
