package chandylamport

import (
	"encoding/gob"
	"log"
	"net"
	f "practice1/functions"
)

// Receive TODO ELIMINAR CALLER
func ReceiveC(chanPoint chan<- string, chanMar chan<- f.Marker, chanMes chan<- f.Message, addr string) error {
	var listener net.Listener
	var decoder *gob.Decoder
	var pack interface{}
	var red net.Conn
	var err error

	log.Println("[Receive] POrt:  ", addr)
	listener, err = net.Listen("tcp", addr)
	f.Error(err, "Listen Error")
	defer listener.Close()

	for {

		red, err = listener.Accept()
		f.Error(err, "Server accept red error")
		// defer red.Close()

		decoder = gob.NewDecoder(red)
		err = decoder.Decode(&pack)
		f.Error(err, "Receive error  \n")

		log.Println("[Receive] PACK", pack)
		switch packNew := pack.(type) {
		case f.Message:
			chanMes <- packNew
			log.Println("[ReceiveM] ===> MESSAGE ", packNew, " DE ", packNew.GetFrom())
		case f.Marker:
			chanMar <- packNew
			log.Println("[ReceiveM] ----> Marker ", packNew, " DE ", packNew.GetCounter())
		case string:
			chanPoint <- packNew
			log.Println("[ReceiveM] ----> checkpoint ", packNew )

		}

		red.Close()

	}

	return err
}
