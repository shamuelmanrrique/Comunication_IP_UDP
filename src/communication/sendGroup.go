package communication

import (
	f "sd_paxos/src/functions"
	"time"
)

// SendGroup function to send group message one to one using TCP
func SendGroup(connect *f.Conn) error {
	var err error
	target := ""
	delay, _ := time.ParseDuration("0s")
	inf := "am dead"
	id := connect.GetId()

	// Update Clock
	vector := connect.GetVector()

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

			print(v)
			msm := &f.Message{
				To:     v,
				From:   id,
				Targ:   target,
				Data:   inf,
				Vector: copyVector,
			}

			if v != target {
				// Get delay
				time.Sleep(delay)
			}

			go Send(v, msm, id)
		}
	}

	return err

}
