package main

import (
	"fmt"
	f "practice1/functions"
	"strconv"
	"strings"
	"time"
)

const (
	ip = "127.0.0.1"   //En este caso se define local
	port = ":5000"
)


func main() {
	// Determinamos el numero de procesos n
	var n int = 1
	// var ids []string = f.IdProcess(n, "remote")
	var ids []string = f.IdProcess(n, "local")
	fmt.Println(ids)
	for _, v := range ids {
		var reception string

		go f.S("XXX")
	
		go f.R(&reception)
	
		time.Sleep(3*time.Second)
	
		fmt.Println(n)

		// fmt.Println(len(os.Args), os.Args)

		// var connect f.Data
		elemts := strings.Split(v, ":")
		// fmt.Println(elemts)
		n, err := strconv.Atoi(elemts[0])
		f.Error(err, "Error Main")
		var connect f.Data = f.Data{
			Inf:  elemts[0],
			Id:   elemts[1] + elemts[2],
			Ip:   elemts[1],
			Num:  n,
			Host: elemts[2],
		}
		proof(connect)
		// Llamamos send
		// f.Send(connect)
		// f.Se(1)
		// f.Receive(connect)

	}
	// fmt.Println(ids)
	time.Sleep(10 * time.Second)
}

func proof(c f.Connection) {
	fmt.Println(c.GetInf())
	fmt.Println(c.GetNum())
	fmt.Println(c.GetHost())
	fmt.Println(c.GetId())
}

// func main() {
// 	var n int = 4
// 	var m Message

// 	// m.Ip = 0.5

// 	// var ips []string:= functions.IdProcess(n,"local")

// }
