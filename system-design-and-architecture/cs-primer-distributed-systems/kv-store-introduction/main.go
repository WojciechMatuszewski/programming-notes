package main

import (
	"fmt"
	"net"
)

const PORT = 8000
const IP = "127.0.0.1"

func main() {
	address := &net.UDPAddr{IP: net.ParseIP(IP), Port: PORT}

	conn, err := net.ListenUDP("udp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("received:%v\n", string(buffer[:n]))
	}
}
