
	red, _ := net.ResolveUDPAddr("udp", connectM.GetId())
	log.Println("[RM]             localhostAddress ", red)

	// printError("ResolvingUDPAddr in Broadcast localhost failed.", er)
	listener, err := net.ListenUDP("udp", red)
	f.Error(err, "[RM] ListenUDP Error")
	defer listener.Close()

	msm := &f.Message{
		To:   f.MulticastAddress,
		From: connectM.GetId(),
		Targ: connectM.GetId(),
		Data: "inf",
	}
	ackID := &f.Ack{Code: connectM.GetId() + "," + msm.GetFrom()}

	pack := &f.Pack{
		Mes:     *msm,
		ConfACK: *ackID,
	}

	canalPacks := make(chan f.Pack)
	defer close(canalPacks)

	v := "pack"

	go m.ReceivePack(canalPacks, listener, connectM.GetPort())
	time.Sleep(time.Second * 1)
	fmt.Println("[Main] before send ")
	go m.SendPack(pack, connectM.GetId())

	fmt.Println("[Main] espero en el canal  ")
	recpack, ok := <-canalPacks

	fmt.Println("[Main] recibi por el canal: ", recpack, v, ok)