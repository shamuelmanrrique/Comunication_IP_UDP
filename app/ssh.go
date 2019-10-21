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

// const (
// 	n    = 2           //numero de procesos
// 	ip   = "127.0.0.1" //En este caso se define local
// 	port = f.GetIp()
// )

func main() {

	config := &ssh.ClientConfig{
		User:            "a802400",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password("sMDJMA-21"),
		},
	}

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
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	// you can also pass what gets input to the stdin, allowing you to pipe
	// content from client to server
	//      session.Stdin = bytes.NewBufferString("My input")

	// Finally, run the command
	err = session.Run("bash; ls")
	fmt.Println(b.String())
	// return b.String(), err

	// // Determinamos el numero de procesos n
	// var ids []string = f.IdProcess(n, "local") // var ids []string = f.IdProcess(n, "remote")
	// fmt.Println(ids)

	// var connect f.Conn = f.Conn{
	// 	Id:   ip + port,
	// 	Ip:   ip,
	// 	Port: port,
	// 	Ids:  ids,
	// }

	// fmt.Println(connect)

	// Necesito escuchar uso r

	//TODO ARRAY de delay

	//LLAMO A FUNCION QUE VA A TENER SEND RECEIVE

}
