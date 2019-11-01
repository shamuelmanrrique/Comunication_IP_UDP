fmt.Println("[MAIN]  --> ", ackID, data, vr, msm)

// ve, ok := data.(f.Message)
vEe, ok1 := data.(f.Ack)
// fmt.Println("[Main] Soy Message", d)
fmt.Println("[Main] Soy Message", vEe, ok1)

ve, ok := data.(f.Message)
// fmt.Println("[Main] Soy Message", d)
fmt.Println("[Main] Soy Message", ve, ok)
switch packNew := data.(type) {
case f.Message:
	fmt.Println("[Main] Soy Message 11", data, packNew)
case f.Ack:
	fmt.Println("[Main] Soy ACK", data, packNew)
}

if false {

	// inicio ReceiveGroupM
	go u.ReceiveGroupM(connectM)
	time.Sleep(time.Second * 2)

	// Si soy master llamo SendGroupM msm
	if flags.Master {

		target := ""
		delay, _ := time.ParseDuration("0s")
		inf := "Me mataron"
		id := connectM.GetId()

		// Actualizo el reloj
		vector := connectM.GetVector()

		if len(connectM.GetKill()) > 0 && len(connectM.GetDelays()) > 0 {
			target = connectM.GetTarget(0)
			delay = connectM.GetDelay(0)
			inf = "He disparado"
			connectM.SetKill()
			connectM.SetDelay()
		}

		// Incremento el reloj
		vector.Tick(id)
		connectM.SetClock(vector)

		// TODO CREATE SNAPSHOP RELOJ []VCLOCK
		// Copio el vector
		copyVector := vector.Copy()

		// IMprimo TODO
		// fmt.Println("[Main] ", copyVector, target, delay, inf)

		// En este caso tomo el target para enviar el delay
		var msm f.Message = f.Message{
			To:     f.MulticastAddress,
			From:   id,
			Targ:   target,
			Data:   inf,
			Vector: copyVector,
			Delay:  delay,
		}

		fmt.Println("Llamo sendGroup MAIN", *connectM)
		time.Sleep(time.Second * 2)
		go u.SendGroupM(&msm, connectM)
	}

	
	
	
	
	// // Define connection to udp
	// redUDP, err = net.ResolveUDPAddr("udp", connect.GetId())
	// f.Error(err, "Send connection error \n")

	// connectionUDP, err = net.ListenUDP("udp", redUDP)
	// f.Error(err, "Send connection error \n")
	// defer connectionUDP.Close()

	// // red, _ := net.ResolveUDPAddr("udp", connect.GetId())
	// // log.Println("[RM]             localhostAddress ", red)

	// // // printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	// // listener, err := net.ListenUDP("udp", red)
	// // f.Error(err, "[RM] ListenUDP Error")
	// // defer listener.Close()

	// var bufferPacks []f.Pack
	// canalPacks := make(chan f.Pack)

	// for i := 0; i < n-1; i++ {
	// 	go ReceivePack(canalPacks, connectionUDP, connect.GetId())

	// 	pt, _ := <-canalPacks
	// 	fmt.Println("[SendGroupM] recibo del canal: ", pt)
	// 	bufferPacks = append(bufferPacks, pt)

	// }

	// deadline := time.Now().Add(2 * time.Second)
	// err = connection.SetDeadline(deadline)