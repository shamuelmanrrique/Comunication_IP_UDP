package chandylamport

import (
	"fmt"
	f "practice1/functions"
	"time"
)

// SendGroup dda
func SendGroupC(chanPoint chan string, chanMess chan f.Message, chanMarker chan f.Marker, connect *f.Conn) error {
	var err error
	target := ""
	delay, _ := time.ParseDuration("0s")
	inf := "am dead"
	id := connect.GetId()

	// Actualizo el reloj
	vector := connect.GetVector()

	if len(connect.GetKill()) > 0 && len(connect.GetDelays()) > 0 {
		target = connect.GetTarget(0)
		delay = connect.GetDelay(0)
		inf = "kill"
		connect.SetKill()
		connect.SetDelay()
	}

	// Incremento el reloj
	vector.Tick(id)
	connect.SetClock(vector)
	copyVector := vector.Copy()

	// Envio el msm a todos
	for _, v := range connect.GetIds() {
		if v != id {

			msm := &f.Message{
				To:     v,
				From:   id,
				Targ:   target,
				Data:   inf,
				Vector: copyVector,
			}

			if v != target {
				time.Sleep(delay)
			}

			// fmt.Println("[SendGroup]: Imprimo value envio a ", v)
			// f.BufferMsm[id+","+v] = *msm
			go SendC(msm, v)
			fmt.Println(msm)

		}
	}

	return err

}
