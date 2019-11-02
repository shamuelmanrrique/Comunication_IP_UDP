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
	m := make(chan f.Message)
	defer close(m)

	go func() {
		// for i := 0; i < n; i++ {
		deadline := time.Now().Add(10 * time.Second)
		for time.Now().Before(deadline) {
			// log.Println("[ReceiveGroupM]  ++++ INICIO FOR ReceiveGroupM i: ", i, " N ", n)

			listener.SetReadBuffer(f.MaxBufferSize)

			buffer := make([]byte, f.MaxBufferSize)
			// nRead, _, _ := listener.ReadFromUDP(buffer)
			listener.ReadFromUDP(buffer)

			dataBuffer := bytes.NewBuffer(buffer)
			decode = gob.NewDecoder(dataBuffer)
			err = decode.Decode(&msm)
			f.Error(err, "Receive error  Decode\n")

			msmMult, ok, _ = f.CheckMsm(msmMult, msm)
			if !ok {
				log.Println("[ReceiveGroupM]  ENVIO UN MSM POR EL CANAL++++++")
				m <- msm
			}
		}
	}()

	var msm1 f.Message
	// Chequeamos si el msm recibido esta en el array de msm
	log.Println("[ReceiveGroupM]  Segundo FOR")
	for {
		// for i := 0; i < n; i++ {
		// log.Println("[ReceiveGroupM]  Esperando tomar un segundo msm")
		msm1, ok = <-m
		// log.Println("[ReceiveGroupM]  ++++ INICIO FOR ReceiveGroupM i: ", i, " N ", n)
		if !ok {
			break
		}
		from := msm1.GetFrom()
		if from != id {
			log.Println("[ReceiveGroupM]  ELSE  CON MSM DE ", msm1.GetFrom())
			ackID := &f.Ack{
				Origen: connect.GetId(),
				Code:   connect.GetId() + "," + from,
			}
			go SendM(ackID, from)

			// RECIBO y sumo 1 al vector
			vector.Tick(id)
			// SEt la nueva actualizacion de recepcion
			connect.SetClock(vector)
			// Uno los relojes
			vector.Merge(msm1.GetVector())
			// connect.GetVector().Merge(vector)
			// Seteo nuevamente el reloj
			connect.SetClock(vector)

			log.Println("[ReceiveGroupM]  Recibido de: ", msm1.GetFrom(), " Yo soy ", id, "Target: ", msm1.GetTarg())
			if msm1.GetTarg() == id {
				log.Println("[ReceiveGroupM] SOY TARGET entre en el IF")
				n--
				//Aplico delay receive
				// delay := msm1.GetDelay()
				log.Println("[ReceiveGroupM] Aplico Delay")
				// time.Sleep(delay * time.Second)

				log.Println("[ReceiveGroupM]  Llamo a send group : ")
				// if msm1.GetTarg() == id {
				go SendGroupM(chanAck, connect)
				// }
			}
			// log.Println("[ReceiveGroupM]  ------ FIN FOR ReceiveGroupM i: ", i, " N ", n)
		}

	}

	// TODO debo recibir paquetes directamente
	log.Println("[ReceiveGroupM] SALGO DEL FOR")
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
