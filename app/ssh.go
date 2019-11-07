package main

import (
	"bytes"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

// ssh a802400@155.210.154.210s
// pwd : sMDJMA-21
// escribir > bash

func main() {

	config := &ssh.ClientConfig{
		User:            "a802400",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password("sMDJMA-21"),
		},
	}
	// ssh a802400@155.210.154.195

	client, err := ssh.Dial("tcp", net.JoinHostPort("155.210.154.210", "22"), config)
	if err != nil {
		panic(err.Error())
	}

	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		panic(err.Error())
	}

	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b

	// Finally, run the command
	err = session.Run("bash; ls")
	fmt.Println(b.String())

}
