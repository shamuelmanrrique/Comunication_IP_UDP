package functions

import (
	"time"
)

func SendGroup(connect Connection, canal chan Message) error {
	var err error
	id := connect.GetId()

	for i, v := range connect.GetIds() {
		if v != id {

			m := "[SEND] => " + connect.GetId() + " He disparado a " + connect.GetKill()

			var msm Message = Message{
				To:   connect.GetId(),
				From: connect.GetEnv(i),
				Data: m,
			}

			delay := time.Duration(connect.GetDelay(i)) * time.Second
			time.Sleep(delay)

			go Send(msm)
		}

	}
	return err

}
