package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	l "practice1/chandylamport"
	c "practice1/communication"
	f "practice1/functions"
	u "practice1/multicast"
	v "practice1/vclock"

	"golang.org/x/crypto/ssh"
)

var flags f.Coordinates

func init() {
	flag.IntVar(&flags.Process, "n", 3, "numero de procesos que vas a crear")
	flag.StringVar(&flags.Run, "r", "local", "Se va correr local o remote")
	flag.StringVar(&flags.Port, "p", ":1400", "puerto que usara el proceso :XXXX")
	flag.StringVar(&flags.IPuse, "i", "127.0.1.1", "puerto que usara el proceso :XXXX")
	flag.StringVar(&flags.IPsRem, "c", "127.0.1.1:5001,127.0.1.1:5002,127.0.1.1:5003", "IPs a connect")
	flag.BoolVar(&flags.Master, "m", false, "pppo")
	flag.BoolVar(&flags.SshExc, "ssh", false, "pppo")
	flag.Var(&flags.TimeDelay, "d", "Lista de flags separados por coma")
	flag.Var(&flags.Target, "t", "listas de ip objectivos")
	flag.StringVar(&flags.Exec, "e", "tcp", "Execution mode")
}

func main() {
	flag.Parse()
	gob.Register(f.Message{})
	gob.Register(f.Marker{})
	gob.Register(f.Ack{})

	// Comentados para pruebas con UDP
	var val bool = len(flags.TimeDelay) != len(flags.Target)
	if val {
		panic("El tamaño del arreglo Targets debe ser igual al de Delays")
		// os.Exit(1)
	}

	var err error
	var ip = flags.GetIPuse()
	port := flags.GetPort()
	n := flags.GetProcess()

	// var ids []string = f.IdProcess(n, flags.GetRun())
	var ids []string = flags.GetIPsRem()
	log.Println(ids)
	var com = f.NewCommand(ids, flags.GetRun())

	// // Inicializo todos el reloj del proceso
	var vector = v.New()
	for _, v := range ids {
		vector[v] = 0
	}

	msmreceive := len(ids) - len(flags.GetTarget()) - 1

	connect := &f.Conn{
		Id:     ip + port,
		Ip:     ip,
		Port:   port,
		Ids:    ids,
		Delays: flags.GetTimeDelay(),
		Kill:   flags.GetTarget(),
		Accept: msmreceive,
		Vector: vector,
	}

	log.Println("Connections :", connect)
	// Ssh connection
	if flags.GetSshExc() {
		go func() {
			name := flags.GetRun()
			for _, k := range ids {
				var session ssh.Session
				if name == "local" {
					s := strings.Split(k, ":")
					ipu, _ := s[0], s[1]

					session, err = f.InitSSH("shamuel", ipu, "/home/shamuel/.ssh/id_rsa")
					if err != nil {
						log.Fatal(err.Error())
					}
					var b bytes.Buffer
					session.Stdout = &b

					// session.Run("bash; ls ; pwd")
					fmt.Println("go run /home/shamuel/go/src/practice1/app/main.go", f.FlagsExec(com, k))
					run := "go run /home/shamuel/go/src/practice1/app/main.go " + f.FlagsExec(com, k)
					err = session.Run(run)
					if err != nil {
						log.Fatal(err.Error())
					}

				} else if name == "proof" {

					session, _ = f.InitSSH("a802400", k, "/home/shamuel/.ssh/id_rsa")
					// log.SetOutput(session.Output())
					// log.SetOutput(os.Stderr)

					var b bytes.Buffer
					session.Stdout = &b

					runCom := "/usr/local/go/bin/go run /home/a802400/go/practice1/app/main.go " + f.FlagsExec(com, k)
					session.Run(runCom)
					fmt.Println(runCom)
				}
			}
		}()

		// Execution Modules
	} else {
		<-time.After(time.Second * 3)

		// TCP
		if flags.GetExec() == "tcp" {
			f.DistMsm("TCP " + ip + port)
			go c.ReceiveGroup(connect)
			if flags.Master {
				time.Sleep(time.Second * 2)
				go c.SendGroup(connect)
			}

		}

		// UDP
		if flags.GetExec() == "udp" {
			f.DistMsm("UDP " + ip + port)

			chanAck := make(chan f.Ack, len(connect.GetIds())-1)
			defer close(chanAck)
			chanMessage := make(chan f.Message, len(connect.GetIds()))
			defer close(chanMessage)

			go u.ReceiveM(chanAck, chanMessage, connect.GetPort())

			go u.ReceiveGroupM(chanMessage, chanAck, connect)
			if flags.GetMaster() {
				time.Sleep(time.Second * 2)
				go u.SendGroupM(chanAck, connect)
			}
		}

		// ChandyLamport
		if flags.GetExec() == "chandy" {
			f.DistMsm("ChandyLamport " + ip + port)
			chanMarker := make(chan f.Marker, n)
			defer close(chanMarker)
			chanMessage := make(chan f.Message, n)
			defer close(chanMessage)
			chanPoint := make(chan string, n)
			defer close(chanPoint)

			// var marker = &f.Marker{}
			ids = nil

			go l.ReceiveGroupC(chanPoint, chanMessage, chanMarker, connect)
			if flags.Master {
				time.Sleep(time.Second * 2)
				go l.SendGroupC(chanPoint, chanMessage, chanMarker, connect)
			}

			marker := f.Marker{
				Counter: len(connect.GetIds()),
				Recoder: false,
			}

			// Init Snapshot
			if flags.Master {
				time.Sleep(time.Second * 4)
				cap := connect.GetEnv(0)
				go l.SendC(marker, cap)
			}

		}
	}

	<-time.After(time.Second * 40)
}
