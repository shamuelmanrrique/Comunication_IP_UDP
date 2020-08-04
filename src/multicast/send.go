package multicast

import (
	f "sd_paxos/src/functions"
	"time"

	"github.com/dmichael/go-multicast/multicast"
)

/*
-----------------------------------------------------------------
METODO: ReceiveGroupM
RECIBE: canal de tipo f.Message, un listener net.Listener
DEVUELVE: OK si todo va bien o ERROR en caso contrario
PROPOSITO: It's a function that receive message using UDP connection  and IP multicast
-----------------------------------------------------------------
*/
func Send() error {
	// var red *net.UDPAddr
	// var connection *net.UDPConn
	// var encoder *gob.Encoder
	// var buffer bytes.Buffer
	var err error

	// // Creating red connection to MulticastAddress
	// red, err = net.ResolveUDPAddr("udp", "229.0.40.000:9999")
	// f.Error(err, "SendGroupM error ResolveUDPAddr connection \n")
	// println(red)

	// // Making dial connection to ip address
	// connection, err = net.DialUDP("udp", nil, red)
	// f.Error(err, "SendGroupM error DialUDP connection \n")
	// defer connection.Close()

	// var msm string = "prueba 1"
	// println("connection", connection)
	conn, err := multicast.NewBroadcaster("229.0.40.000:9999")
	if err != nil {
		f.Error(err, "SendGroupM error DialUDP connection \n")
	}

	msm := "Ja Wueno \n"
	// Send message to MulticastAddress
	go func() {
		// Send the same message three times
		for i := 0; i < 3; i++ {
			// encoder = gob.NewEncoder(&buffer)
			// err = encoder.Encode(msm)
			// f.Error(err, "SendGroupM encoder error \n")
			conn.Write([]byte(msm))

			// _, err = connection.Write(buffer.Bytes())
			// f.Error(err, "SendGroupM error ResolveUDPAddr connection \n")

			// Sleep between every delivery
			println("SendGroupM =======message", i, msm)
			println("SendGroupM =======message", i, &msm)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return err
}
