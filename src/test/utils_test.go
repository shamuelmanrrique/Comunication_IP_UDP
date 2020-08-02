package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"gopkg.in/ini.v1"

	f "sd_paxos/src/functions"
)

func init() {
	testing.Init()
}

func TestSimpleSSH(t *testing.T) {

	connection := f.InitSSH("127.0.0.1")
	println(connection, "localhost", "ip")
	go f.ExcecuteSSH("ls", connection)
	time.Sleep(5 * time.Second)
}

func TestSSH(t *testing.T) {
	// Change values of
	// 		mode -> whatever you want tcp,udp,chandy
	//		log  -> true, false
	//      enviromment -> develomment, production

	// Declaring variables
	var environment string
	var logMode string
	var mode string
	var path string
	var machinesID []string
	var machinesName []string

	// Loading configuration file
	// cfg, err := ini.Load("~/go/src/sd_paxos/src/config/go.ini")
	cfg, err := ini.Load("../config/go.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Getting configuration values from .ini
	environment = cfg.Section("general").Key("environment").String()
	path = cfg.Section(environment).Key("mainPath").String()
	logMode = cfg.Section("general").Key("log").String()
	mode = cfg.Section("general").Key("mode").String()
	machinesID = strings.Split(cfg.Section(environment).Key("machinesID").String(), ",")
	machinesName = strings.Split(cfg.Section(environment).Key("machinesName").String(), ",")

	for i, ip := range machinesID {
		addr := strings.Split(ip, ":")
		connection := f.InitSSH(addr[0])
		println(path+" -name="+machinesName[i]+" -mode="+mode+" -log="+logMode, ip)

		go f.ExcecuteSSH("cd ~/go/src/sd_paxos/src ; go run main.go -name="+machinesName[i]+" -mode="+mode+" -log="+logMode, connection)
	}

	time.Sleep(50 * time.Second)
}
