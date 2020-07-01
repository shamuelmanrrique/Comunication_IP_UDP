package main

import (
	"fmt"
	"log"

	"github.com/tkanos/gonfig"
)

// Coment
type Configuration struct {
	Port              int
	Static_Variable   string
	Connection_String string
}

func main() {
	configuration := Configuration{}
	err := gonfig.GetConf("./config/config_local.json", &configuration)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(configuration.Port)
}
