package main

import (
	// 	"time"
	"fmt"
	f "practice1/functions"
	"strconv"
	"strings"
)

func main() {
	var n int = 4
	var ids []string = f.IdProcess(n, "local")

	for _, v := range ids {
		// var connect f.Data
		elemts := strings.Split(v, ":")
		fmt.Println(elemts)
		n, err := strconv.Atoi(elemts[0])
		f.Error(err)
		var connect f.Data = f.Data{
			Inf:  elemts[0],
			Id:   elemts[1] + elemts[2],
			Ip:   elemts[1],
			Num:  n,
			Host: elemts[2],
		}
		proof(connect)

		// Llamamos send
		go functions.Send("XXX")
	}
	// fmt.Println(ids)

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

// 	// var reception string

// 	// go functions.Send("XXX")

// 	// go functions.Receive(&reception)

// 	// time.Sleep(3*time.Second)

// 	fmt.Println(n)
// }
