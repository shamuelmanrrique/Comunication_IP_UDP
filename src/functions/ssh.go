package functions

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"gopkg.in/ini.v1"
)

func InitSSH(addr string) *ssh.Client {
	var environment string
	var IDRsa string
	var user string

	// Loading configuration file
	cfg, err := ini.Load("../config/go.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Getting configuration values from .ini
	environment = cfg.Section("general").Key("environment").String()
	IDRsa = cfg.Section("general").Key("IDRsa").String()
	user = cfg.Section(environment).Key("user").String()

	println(environment, IDRsa, user)

	key, err := ioutil.ReadFile(IDRsa)
	if err != nil {
		panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
	if err != nil {
		panic(err.Error())
	}

	return client

}

func ExcecuteSSH(cmd string, conn *ssh.Client) {
	sess, err := conn.NewSession()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	sessStdOut, err := sess.StdoutPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stdout, sessStdOut)
	sessStderr, err := sess.StderrPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stderr, sessStderr)
	log.Println(cmd)
	err = sess.Run(cmd)
	if err != nil {
		panic(err)
	}
}
