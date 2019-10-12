package main

import (
	// 	"time"
	"fmt"
	"practice1/functions"
)

// type Message interface {
// 	Ip solo usa procedure
// }

type Message struct {
	ip, data, vector string
	// el msm debe ser una variable local
}

func main() {
	var n int = 4
	var m Message

	m.data = "jh"
	var ids []string = functions.IdProcess(n, "local")

	for i:= 0; i <= n; i++ {
		
	} 
	fmt.Println(ids)
}

// func main() {
// 	var n int = 4
// 	var m Message

// 	// m.Ip = 0.5

// 	// var ips []string:= functions.IdProcess(n,"local")

// 	// var reception string

// 	// go functions.Send("XXX")

// 	// go functions.Receive(&reception)

// 	// time.Sleep(3*time.Second)

// 	fmt.Println(n)
// }
