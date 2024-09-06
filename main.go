package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Listen on UDP port 9999
	addr, err := net.ResolveUDPAddr("udp4", ":25643")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Listening for broadcast messages on port 25643...")

	buffer := make([]byte, 1024) // Adjust buffer size as needed

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		if n > 0 {
			receivedHostname := string(buffer[0:n])
			fmt.Printf("Received hostname: %s from %s\n", receivedHostname, remoteAddr)
		}
	}
}
