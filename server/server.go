package main

import (
	"log"
	"net"
	"os"

	"github.com/gen2brain/beeep"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":5000")
	if err != nil {
		log.Fatalln("Wrong Address")
		return
	}
	updLn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatalln(err)
		return
	}

	buf := make([]byte, 1024)
	log.Println("Starting udp server...")

	for {
		n, addr, err := updLn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		go func() {
			log.Printf("Reciving data: %s from %s", string(buf[:n]), addr.String())
			err := beeep.Notify("DoorChecker", "ドアが開いたよ", "assets/warning.png")
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
}
