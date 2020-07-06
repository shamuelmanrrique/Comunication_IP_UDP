package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tkanos/gonfig"
	"gopkg.in/ini.v1"
)

// Coment
type Machine struct {
	port            string
	ip              string
	role            string
	MachinesDelay   string
	machinesTargets string
}

// Coment
type Configuration struct {
	ssh           bool
	jobsNumer     int
	executionMode string
	machinesID    string
	machine       []Machine
	// machine2      Machine
	// machine3      Machine
}

func main() {
	configuration := Configuration{}
	err := gonfig.GetConf("./config/config_local.json", &configuration)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(configuration)

	cfg, err := ini.Load("./config/go.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Classic read of values, default section can be represented as empty string
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

	// Let's do some candidate value limitation
	fmt.Println("Server Protocol:",
		cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// Value read that is not in candidates will be discarded and fall back to given default value
	fmt.Println("Email Protocol:",
		cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// Try out auto-type conversion
	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// Now, make some changes and save it
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("my.ini.local")

}
