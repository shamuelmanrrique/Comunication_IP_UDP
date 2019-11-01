package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"time"

	c "practice1/communication"
	f "practice1/functions"
	v "practice1/vclock"
)

var flags f.Coordinates

func init() {
	flag.IntVar(&flags.Process, "n", 4, "numero de procesos que vas a crear")
	flag.StringVar(&flags.Run, "r", "local", "Se va correr local o remote")
	flag.StringVar(&flags.Port, "p", ":1400", "puerto que usara el proceso :XXXX")
	flag.BoolVar(&flags.Master, "m", false, "pppo")
	flag.Var(&flags.TimeDelay, "d", "Lista de flags separados por coma")
	flag.Var(&flags.Target, "t", "listas de ip objectivos")
}

func main() {
	flag.Parse()
	// Comentados para pruebas con UDP
	var val bool = len(flags.TimeDelay) != len(flags.Target)
	if val {
		panic("El tamaño del arreglo Targets debe ser igual al de Delays")
		os.Exit(1)
	}

	ip := f.IpAddress()
	port := flags.GetPort()
	n := flags.GetProcess()
	f.DistMsm(ip + port)

	var ids []string = f.IdProcess(n, flags.GetRun())

	// Inicializo todos el reloj del proceso
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

	if true {

		// var s []ssh.Session
		// for _, v := range ids {

		v, _ := f.InitSSH("a802400", "195.210.154.210", "/home/shamuel/.ssh/id_rsa")
		// s = append(s, aux)

		// }

		// // v, _ := f.InitSSH("shamuel", "localhost", "/home/shamuel/.ssh/id_rsa")
		var b bytes.Buffer
		v.Stdout = &b

		// // Finally, run the command
		// // v.Run("go run go/src/game/main.go -proc " + strconv.Itoa(i+1) + " -n 4 >> log" + middleware.Addresses[i+1] + ".txt")
		v.Run("bash; ls ; pwd")
		// v.Run("cd \"/home/shamuel/go/src/practice1/app\" ;ls ; pwd;go run main.go -r \"local\" -t \"127.0.1.1:5002\" -d \"5s\" -n 3 -m=true -p=\":5001\" > ho.txt")
		fmt.Println(b.String())

	}
	fmt.Println("ESTOY FUERA")
	go c.ReceiveGroup(connect)
	if flags.Master {
		// fmt.Println("Llamo sendGroup MAIN", *connect)
		time.Sleep(time.Second * 1)
		go c.SendGroup(connect)
	}

	<-time.After(time.Second * 30)

}
