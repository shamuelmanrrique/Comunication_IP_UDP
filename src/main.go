package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/ini.v1"

	l "sd_paxos/src/chandylamport"
	c "sd_paxos/src/communication"
	f "sd_paxos/src/functions"
	u "sd_paxos/src/multicast"
	v "sd_paxos/src/vclock"
)

var flags f.Coordinates

func init() {
	flag.StringVar(&flags.Machine, "m", "machine1", "Insert name like machine# (# is a number 1-3) ")
	flag.StringVar(&flags.Mode, "e", "tcp", "Mode to execute [tcp, udp, chandy]")
}

func main() {
	// Loading configuration file
	cfg, err := ini.Load("./config/go.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Declaring variables
	var ssh bool
	var jobs int
	var ip string
	var port string
	var role string
	var mode string
	var delays []time.Duration = make([]time.Duration, 0)
	var targets []string = make([]string, 0)
	var environment string
	var machineName string
	var machinesID []string

	// Parcing flags
	flag.Parse()
	gob.Register(f.Message{})
	gob.Register(f.Marker{})
	gob.Register(f.Ack{})

	environment = cfg.Section("general").Key("environment").String()
	jobs, err = cfg.Section("general").Key("jobs").Int()
	ssh, err = cfg.Section("general").Key("ssh").Bool()
	machineName = flags.GetMachine()
	mode = flags.GetMode()

	ip = cfg.Section(environment + " " + machineName).Key("ip").String()
	port = cfg.Section(environment + " " + machineName).Key("port").String()
	role = cfg.Section(environment + " " + machineName).Key("role").String()

	target := cfg.Section(environment + " " + machineName).Key("targets").String()
	if target != "" {
		targets = strings.Split(target, ",")
	}

	machinesID = strings.Split(cfg.Section(environment).Key("machinesID").String(), ",")

	durat := cfg.Section(environment + " " + machineName).Key("delays").String()
	if durat != "" {
		for _, v := range strings.Split(durat, ",") {
			duration, _ := time.ParseDuration(v)
			delays = append(delays, duration)
		}
	}

	// Inicializo todos el reloj del proceso
	var vector = v.New()
	for _, v := range machinesID {
		vector[v] = 0
	}

	println(ssh, jobs, mode, ip, port, role, machinesID)

	// Calculating how many message it will receive
	msmreceive := len(machinesID) - len(targets) - 1

	connect := &f.Conn{
		Id:     ip + port,
		Ip:     ip,
		Port:   port,
		Ids:    machinesID,
		Kill:   targets,
		Delays: delays,
		Accept: msmreceive,
		Vector: vector,
	}

	<-time.After(time.Second * 5)

	// TCP
	if mode == "tcp" {
		f.DistMsm("TCP " + ip + port)
		go c.ReceiveGroup(connect)
		if role == "master" {
			time.Sleep(time.Second * 2)
			go c.SendGroup(connect)
		}

	}

	// UDP
	if mode == "udp" {
		f.DistMsm("UDP " + ip + port)

		chanAck := make(chan f.Ack, len(connect.GetIds())-1)
		defer close(chanAck)
		chanMessage := make(chan f.Message, len(connect.GetIds()))
		defer close(chanMessage)

		go u.ReceiveM(chanAck, chanMessage, connect.GetPort())

		go u.ReceiveGroupM(chanMessage, chanAck, connect)
		if role == "master" {
			time.Sleep(time.Second * 2)
			go u.SendGroupM(chanAck, connect)
		}
	}

	// ChandyLamport
	if mode == "chandy" {
		f.DistMsm("ChandyLamport " + ip + port)
		chanMarker := make(chan f.Marker, jobs)
		defer close(chanMarker)
		chanMessage := make(chan f.Message, jobs)
		defer close(chanMessage)
		chanPoint := make(chan string, jobs)
		defer close(chanPoint)

		go l.ReceiveGroupC(chanPoint, chanMessage, chanMarker, connect)
		if role == "master" {
			time.Sleep(time.Second * 2)
			go l.SendGroupC(chanPoint, chanMessage, chanMarker, connect)
		}

		marker := f.Marker{
			Counter: len(machinesID),
			Recoder: false,
		}

		// Init Snapshot
		if role == "master" {
			time.Sleep(time.Second * 4)
			cap := connect.GetEnv(0)
			go l.SendC(marker, cap)
		}

	}

	<-time.After(time.Second * 60)

}
