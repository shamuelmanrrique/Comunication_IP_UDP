package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
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
var checklog bool

func init() {
	// Register all interface to use
	gob.Register(f.Message{})
	gob.Register(f.Pack{})
	gob.Register(f.Marker{})
	gob.Register(f.Ack{})
	
	// Reading flags from terminal
	flag.StringVar(&machineName, "name", "machine1", "Insert name like machine# (# is a number 1-3) ")
	flag.StringVar(&mode, "mode", "tcp", "Mode to execute [tcp, udp, chandy]")
	flag.BoolVar(&checklog, "log", true, "Send output to file true otherwise false")
}

func main() {
	// Parcing flags
	flag.Parse()

	// Loading configuration file
	cfg, err := ini.Load("./config/go.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Getting configuration values from .ini
	jobs, err = cfg.Section("general").Key("jobs").Int()
	ssh, err = cfg.Section("general").Key("ssh").Bool()
	environment = cfg.Section("general").Key("environment").String()
	machinesID = strings.Split(cfg.Section(environment).Key("machinesID").String(), ",")
	port = cfg.Section(environment + " " + machineName).Key("port").String()
	role = cfg.Section(environment + " " + machineName).Key("role").String()
	ip = cfg.Section(environment + " " + machineName).Key("ip").String()

	target := cfg.Section(environment + " " + machineName).Key("targets").String()
	if target != "" {
		targets = strings.Split(target, ",")
	}

	durat := cfg.Section(environment + " " + machineName).Key("delays").String()
	if durat != "" {
		for _, v := range strings.Split(durat, ",") {
			duration, _ := time.ParseDuration(v)
			delays = append(delays, duration)
		}
	}

	// Writting output in log if checklog is true
	if checklog {
		file, err := os.OpenFile("logs/["+mode+"-"+ip+port+"]-"+machineName+".log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	// Calculating how many message machine will receive
	msmreceive := len(machinesID) - len(targets) - 1

	// Init every machine clock
	var vector = v.New()
	for _, v := range machinesID {
		vector[v] = 0
	}

	// Declaring connection variable
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

	<-time.After(time.Second * 2)

	// Executing in TCP  mode
	if mode == "tcp" {
		f.DistMsm("TCP " + ip + port)

		// Calling gorutine to receive direct message
		go c.ReceiveGroup(connect)

		// Master Node send first message
		if role == "master" {
			time.Sleep(time.Second * 5)
			go c.SendGroup(connect)
		}

	}

	// Executing in  UDP mode
	if mode == "udp" {
		f.DistMsm("UDP " + ip + port)

		// Creating channel to ACK
		chanAck := make(chan f.Ack, len(connect.GetIds())-1)
		defer close(chanAck)

		// Creating channel to ACK
		chanMessage := make(chan f.Message, len(connect.GetIds()))
		defer close(chanMessage)

		// Calling gorutine to receive direct message
		go u.ReceiveM(chanAck, chanMessage, connect.GetPort())

		// Calling gorutine to receive broadcast message
		go u.ReceiveGroupM(chanMessage, chanAck, connect)
		// go u.ReceiveGroupMr()

		// Master Node send first message
		if role == "master" {
			time.Sleep(time.Second * 2)
			go u.SendGroupM(chanAck, connect)
			// go u.Send()
		}
	}

	// Executin in ChandyLamport mode
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
