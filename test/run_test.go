package main

import (
	"io"
	"log"
	"os"
	f "practice1/functions"
	"strings"
	"testing"
)

func TestTCP(t *testing.T) {
	rsa := "/home/shamuel/.ssh/id_rsa"
	ips := []string{"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400"}
	valuesRun := []string{
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.199\" -t \"155.210.154.209:1400\" -d \"5s\"  -m=true   -e=\"tcp\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.208\" -e=\"tcp\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.209\" -e=\"tcp\"",
	}

	for i, k := range ips {
		ip := strings.Split(k, ":")
		session, err := f.InitSSH("a802400", ip[0], rsa)
		defer session.Close()
		if err != nil {
			log.Fatal(err.Error())
		}

		sessionOut, err := session.StdoutPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stdout, sessionOut)
		sessionError, err := session.StderrPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stderr, sessionError)
		run := f.GetString(i, valuesRun)
		err = session.Run(run)
		if err != nil {
			panic(err)
		}
	}

	for {
	}
}

func TestTCP2(t *testing.T) {
	rsa := "/home/shamuel/.ssh/id_rsa"
	ips := []string{"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400"}
	valuesRun := []string{
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -t \"155.210.154.209:1400\" -d \"5s\"  -m=true   -e=\"tcp\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -e=\"tcp\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -e=\"tcp\"",
	}

	for i, k := range ips {
		ip := strings.Split(k, ":")
		session, err := f.InitSSH("a802400", ip[0], rsa)
		defer session.Close()
		if err != nil {
			log.Fatal(err.Error())
		}

		sessionOut, err := session.StdoutPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stdout, sessionOut)
		sessionError, err := session.StderrPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stderr, sessionError)
		run := f.GetString(i, valuesRun)
		run = run + " -i=" + ip[0]
		err = session.Run(run)
		if err != nil {
			panic(err)
		}
	}

	for {
	}
}

func TestUDP(t *testing.T) {
	rsa := "/home/shamuel/.ssh/id_rsa"
	ips := []string{"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400"}
	valuesRun := []string{
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.199\" -t \"155.210.154.209:1400\" -d \"5s\"  -m=true   -e=\"udp\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.208\" -e=\"udp\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.209\" -e=\"udp\"",
	}

	for i, k := range ips {
		ip := strings.Split(k, ":")
		session, err := f.InitSSH("a802400", ip[0], rsa)
		defer session.Close()
		if err != nil {
			log.Fatal(err.Error())
		}

		sessionOut, err := session.StdoutPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stdout, sessionOut)
		sessionError, err := session.StderrPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stderr, sessionError)
		run := f.GetString(i, valuesRun)
		err = session.Run(run)
		if err != nil {
			panic(err)
		}
	}

	for {
	}
}

func TestChandyLamport(t *testing.T) {
	rsa := "/home/shamuel/.ssh/id_rsa"
	ips := []string{"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400"}
	valuesRun := []string{
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.199\" -t \"155.210.154.209:1400\" -d \"5s\"  -m=true   -e=\"chandy\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.208\" -e=\"chandy\"",
		"/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.209\" -e=\"chandy\"",
	}

	for i, k := range ips {
		ip := strings.Split(k, ":")
		session, err := f.InitSSH("a802400", ip[0], rsa)
		defer session.Close()
		if err != nil {
			log.Fatal(err.Error())
		}

		sessionOut, err := session.StdoutPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stdout, sessionOut)
		sessionError, err := session.StderrPipe()
		if err != nil {
			panic(err)
		}

		go io.Copy(os.Stderr, sessionError)
		run := f.GetString(i, valuesRun)
		err = session.Run(run)
		if err != nil {
			panic(err)
		}
	}

	for {
	}
}
