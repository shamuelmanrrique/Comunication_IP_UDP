package functions

import (
	"time"
)

func SendGroup(connect Connection) error {
	var err error
	id := connect.GetId()

	// Actualizo reloj
	vector := connect.GetVector()
	vector.Tick(id)
	// fmt.Println(id)

	// Defino msm a enviar
	m := "[ " + id + " -> SEND] => He disparado a " + connect.GetKill()
	var msm Message = Message{
		To:     id,
		From:   connect.GetKill(), //TODO PROBLEMAS
		Data:   m,
		Vector: vector,
	}

	for i, v := range connect.GetIds() {
		// TODO PROBLEMAS CUANDO ENVIO MSM DE VUELTA NUEVAMENTE (CONDICION DE PARADA)
		// println(v != id)
		// println(i)
		if v != id {
			// Aplico delay en el envio
			msm.Ignor = v
			// fmt.Println(msm)
			delay := time.Duration(connect.GetDelay(i))
			time.Sleep(delay * time.Millisecond)
			// fmt.Printf("_____________Estoy SendGroup_______________ %s \n", connect.GetId())
			// fmt.Println(msm)
			go Send(msm)
		}

	}
	return err

}
