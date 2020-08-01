package communication

import (
	f "sd_paxos/src/functions"
	"time"
)

/*
-----------------------------------------------------------------
METODO: SendGroup
RECIBE: ip address "ip",message "f.Msm", sender "caller"
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function to send group message one to one using TCP 
-----------------------------------------------------------------
*/
func SendGroup(connect *f.Conn) error {
	var err error
	target := ""
	delay, _ := time.ParseDuration("0s")
	inf := "am dead"
	id := connect.GetId()

	// Update Clock
	vector := connect.GetVector()

	// Getting target and delaies
	if len(connect.GetKill()) > 0 && len(connect.GetDelays()) > 0 {
		target = connect.GetTarget(0)
		delay = connect.GetDelay(0)
		inf = "kill"
		connect.SetKill()
		connect.SetDelay()
	}

	// Increase clock
	vector.Tick(id)
	connect.SetClock(vector)
	copyVector := vector.Copy()

	// Send message to everyone
	for _, v := range connect.GetIds() {
		if v != id {

			msm := &f.Message{
				To:     v,
				From:   id,
				Targ:   target,
				Data:   inf,
				Vector: copyVector,
			}

			// Set delay if it isn't target
			if v != target {
				time.Sleep(delay)
			}

			go Send(v, msm, id)
		}
	}

	return err

}
