package multicast

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
	"sort"

	v "practice1/vclock"
	"time"
)

// ReceiveGroupM  das
func ReceiveGroupM(chanMess chan f.Message, chanAck chan f.Ack, connect *f.Conn) error {
	var err error
	var msm f.Message
	var ok bool
	var decode *gob.Decoder
	var listener *net.UDPConn
	var arrayMsms []f.Message
	n := len(connect.GetIds()) - 1
	vector := connect.GetVector()
	id := connect.GetId()

	// Parse the string address
	addr, _ := net.ResolveUDPAddr("udp", f.MulticastAddress)

	// Open up a connection
	listener, err = net.ListenMulticastUDP("udp", nil, addr)
	f.Error(err, "ReceiveGroupM error ListenPacket")
	defer listener.Close()

	var msmMult []f.Message

	func() {
		for i := 0; i < n; i++ {
			log.Println("[ReceiveGroupM]  ++++ INICIO FOR ReceiveGroupM i: ", i, " N ", n)

			listener.SetReadBuffer(f.MaxBufferSize)

			buffer := make([]byte, f.MaxBufferSize)
			// nRead, _, _ := listener.ReadFromUDP(buffer)
			listener.ReadFromUDP(buffer)

			dataBuffer := bytes.NewBuffer(buffer)
			decode = gob.NewDecoder(dataBuffer)
			err = decode.Decode(&msm)
			f.Error(err, "Receive error  Decode\n")

			msmMult, ok, _ = f.CheckMsm(msmMult, msm)
			log.Println("[ReceiveGroupM]  RESULTADO DE msmMult ", msmMult, " el valor de OK :", ok)
			if ok {
				i--
			}
		}
	}()

	// Chequeamos si el msm recibido esta en el array de msm
	from := msm.GetFrom()
	if from != id {
		msmMult, ok, msm = f.CheckMsm(msmMult, msm)
		log.Println("[ReceiveGroupM]  RESULTADO DE msmMult ", msmMult, " el valor de OK :", ok)
		if ok {
			i--
		} else {
			log.Println("[ReceiveGroupM]  ELSE  CON MSM DE ", msm.GetFrom())
			ackID := &f.Ack{Code: connect.GetId() + "," + from}
			go SendM(ackID, from)

			// RECIBO y sumo 1 al vector
			vector.Tick(id)
			// SEt la nueva actualizacion de recepcion
			connect.SetClock(vector)
			// Uno los relojes
			vector.Merge(msm.GetVector())
			// connect.GetVector().Merge(vector)
			// Seteo nuevamente el reloj
			connect.SetClock(vector)

			log.Println("[ReceiveGroupM]  Recibido de: ", msm.GetFrom(), " Yo soy ", id, "Target: ", msm.GetTarg())
			if msm.GetTarg() == id {
				log.Println("[ReceiveGroupM] SOY TARGET entre en el IF")
				n--
				//Aplico delay receive
				delay := msm.GetDelay()
				log.Println("[ReceiveGroupM] Aplico Delay")
				time.Sleep(delay * time.Second)

				log.Println("[ReceiveGroupM]  Target: ", msm.GetTarg(), " Recibio ", id)
				// if msm.GetTarg() == id {
				go SendGroupM(chanAck, connect)
				// }

			}
		}

		log.Println("[ReceiveGroupM]  ------ FIN FOR ReceiveGroupM i: ", i, " N ", n)
	}

	// TODO debo recibir paquetes directamente

	log.Println("[ReceiveGroupM] SALDO DEL FOR")
	// Ordeno el arreglo de msm
	sort.SliceStable(arrayMsms, func(i, j int) bool {
		return arrayMsms[i].Vector.Compare(arrayMsms[j].Vector, v.Descendant)
	})

	log.Println("|||||||||||||| FIN ReceiveGroupM ||||||||||||||||||||")
	<-time.After(time.Second * 2)
	for _, m := range arrayMsms {
		log.Println("[Message] --> To: ", m.GetTo(), " From: ", m.GetFrom(), " inf: ", m.GetData())
	}

	return err
}
