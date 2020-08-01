package main

import (
	"log"
	"strings"
	"testing"
	"time"

	f "sd_paxos/src/functions"
)

func init() {
	testing.Init()
}

func TestSSH(t *testing.T) {
	for name, ip := range f.RemoteIP3 {
		addr := strings.Split(ip, ":")
		connection := f.InitSSH(addr[0])
		log.Println(connection, name, ip)
		// go f.ExcecuteSSH(f.GoMainLog+" -i="+ip+" -n="+name, connection)
		// go u.ExcecuteSSH(u.GoMainLog+" -ip="+ip+" -n="+name, connection)
		go f.ExcecuteSSH("ls", connection)
		println("val")
		// Ready
		// go u.ExcecuteSSH(u.GoTest+"TestConnect", connection)
		// go u.ExcecuteSSH(u.GoTest+"TestLog", connection)
	}
	time.Sleep(50 * time.Second)
}

// func TestConnect(t *testing.T) {
// 	cons := f.NewConnec(u.RemoteIP3s)
// 	log.Println(cons)
// }

// var LocalIP3 = map[string]string{"TestSubNetL0": "127.0.1.1:5000", "TestSubNetL1": "127.0.1.1:5001", "TestSubNetL2": "127.0.1.1:5002"}

// var conects3 = u.NewConnec(u.LocalIP3s)

// func TestLocalSSH(t *testing.T) {
// 	for testS, ip := range u.LocalIP3 {
// 		addr := strings.Split(ip, ":")
// 		log.Println(addr[0])
// 		connection := u.InitSSH(addr[0])
// 		log.Println(" -ip="+ip+" -n="+testS, connection)
// 		// go u.ExcecuteSSH(u.GoMainLog+" -ip="+ip+" -n="+testS, connection)
// 		// go u.ExcecuteSSH(u.GoLocalTest+"TestLog", connection)
// 		break
// 	}

// 	time.Sleep(300 * time.Second)

// }
