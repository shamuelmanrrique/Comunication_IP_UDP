package main

import (
	"io"
	"log"
	"os"
	f "practice1/functions"
	"strings"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestTCP(t *testing.T) {
	rsa := "/home/shamuel/.ssh/id_rsa"
	ips := []string{"155.210.154.207:1400,,155.210.154.208:1400,155.210.154.209:1400"}
	valuesRun := []string{
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.207:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.207\" -e=\"tcp\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.207:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.208\" -e=\"tcp\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.207:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.209\" -t \"155.210.154.208:1400\" -d \"5s\"  -m=true   -e=\"tcp\"",
	}

	for i, k := range ips {
		var err error
		var sessionion ssh.Session
		ip := strings.Split(k, ":")
		sessionion, err = f.InitSSH("a802400", ip[0], rsa)
		defer sessionion.Close()
		if err != nil {
			log.Fatal(err.Error())
		}

		sessionionOut, err := sessionion.StdoutPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stdout, sessionionOut)
		sessionionError, err := sessionion.StderrPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stderr, sessionionError)
		// err = sessionion.Run("export PATH=$PATH:/usr/local/go/bin;export GOPATH=/home/a802400/go;export GOROOT=/usr/local/go;")
		run := f.GetString(i, valuesRun)
		go sessionion.Run(run)
		// err = sessionion.Run("ls;pwd")
		if err != nil {
			panic(err)
		}
	}

	for {
	}
}
