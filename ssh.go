package main

import (
	"bytes"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

// ssh a802400@155.210.154.210s
// pwd : sMDJMA-21
// escribir > bash

// ssh a802400@155.210.154.195

// "155.210.154.210"

func main(user string, addr string, idRsa string) (ssh.Session, error) {

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", addr+":22", config)

	config := &ssh.ClientConfig{
		User: shamuel,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort("localhost", "22"), config)
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
