package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":5000")
	if err != nil {
		fmt.Println("Wrong Address")
		return
	}
	updLn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
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
			err := exec.Command("osascript", "-s", "h", "-e", `display notification "ドアが開いたよ" sound name "Submarine.aiff"`).Run()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
}
