package main

import (
	// 	"time"
	"fmt"
	f "practice1/functions"
	"strings"
)

func main() {
	var n int = 4
	var ids []string = f.IdProcess(n, "local")

	// var p string =
	// primes := [1]string{"2:127.0.0.1:1402"}
	// var d Data
	// d.data = primes

	primes := [1]string{"2:127.0.0.1:1402"}
	// reg := ["2:127.0.0.1:1402"]
	a := strings.Join(primes[:], " ")
	d := strings.Split(a, ":")
	fmt.Println(d[1])
	for _, v := range ids {
		var conect f.Data
		elemts := strings.Split(v, ":")
		conect.Inf = elemts
		conect.Getdata()
		// proof(conect)
	}

	fmt.Println(ids)
}

func proof(c f.Conection) {
	fmt.Println(c.GetId())
	// fmt.Println(c.num)
	// fmt.Println(c.host)
	// fmt.Println(c.id)
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
