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
	vector.Tick(id)
	// fmt.Println(count)
	// vector = vector.Set(id, count+1)
	fmt.Println(vector)

	for i, v := range connect.GetIds() {
		if v != id {

			m := "[SEND] => " + id + " He disparado a " + connect.GetKill()

			var msm Message = Message{
				To:     id,
				From:   connect.GetEnv(i), //TODO PROBLEMAS
				Data:   m,
				Vector: vector,
			}

			delay := time.Duration(connect.GetDelay(i)) * time.Second
			time.Sleep(delay)

			go Send(msm)
		}

	}
	return err

}
