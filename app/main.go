package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	// "net"

	f "practice1/functions"
	v "practice1/vclock"
	"time"
)

// Estas constantes pasaran como flash en la consola
const (
	n    = 2           // Determinamos el numero de procesos n
	ip   = "127.0.0.1" //En este caso se define local
	port = ":5001"
)

// Example 3: A user-defined flag type, a slice of durations.
type interval []time.Duration

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (i *interval) Set(value string) error {
	// If we wanted to allow the flag to be set multiple times,
	// accumulating values, we would delete this if statement.
	// That would permit usages such as
	//	-deltaT 10s -deltaT 15s
	// and other combinations.
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// Define a flag to accumulate durations. Because it has a special type,
// we need to use the Var function and therefore create the flag during
// init.

var intervalFlag interval

func init() {
	// Tie the command-line flag to the intervalFlag variable and
	// set a usage message.
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
	// var intervalFlag interval
	//Estoy coloacando cuatro procesos por defoult
	// var process = flag.Int("process", 4, "numero de procesos que quieres crear")

	// t := make([]time.Duration, process)

	// var intervalFlag interval
	// p := flag.Var(&intervalFlag, "p", "comma-separated list of intervals to use between events")
	flag.Parse()

	//var delay = flag.Duration("delay",, "arreglo de retardos")
	//var procesos = flag.Int("procesos", 4, "numero de procesos que quieres crear")
	// fmt.Println(.String())
	// fmt.Println("count value ", *process)
	fmt.Printf("###################### MAIN  %s ########################### \n", ip+port)
	d := []int{5, 8}
	kill := "127.0.0.1:5																																																																																	002"
	var ids []string = f.IdProcess(n, "local")

	// Inicializo todos el reloj del proceso
	var vector = v.New()
	for _, v := range ids {
		vector[v] = 0
	}

	var connect f.Conn = f.Conn{
		Id:     ip + port,
		Ip:     ip,
		Port:   port,
		Ids:    ids,
		Delay:  d,
		Kill:   kill,
		Vector: vector,
	}

	// Proceso maestro llama el send y receive de una vez
	//go f.ReceiveGroup(connect, n)
	time.Sleep(time.Second * 2)
	go f.SendGroup(connect)

	//<-time.After(time.Second * 20)
}
