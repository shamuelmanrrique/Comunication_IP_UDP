// func main() {

//    port := 8001
//    fmt.Println("Listening for UDP Broadcast Packet")


//    socket, err := net.ListenUDP("udp4", &net.UDPAddr{
//       IP:   net.IPv4(0, 0, 0, 0),
//       //IP:   net.IPv4( 192, 168, 1, 255 ),
//       Port: port,
//    })

//    if err != nil {
//       fmt.Println("Error listen: " , err)

//    }
//    for {
//       data := make([]byte, 4096)
//       read, remoteAddr, err := socket.ReadFromUDP(data)
//       if err != nil {
//          fmt.Println("readfromudp: ", err)
//       }
//       for i :=0; i<read; i++{
//          fmt.Println(data[i])
//       }
//       fmt.Printf("Read from: %v\n", remoteAddr)
//    }
// }
