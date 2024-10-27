package networking

import (
	"log"
	"net"
	"os"
)

type SignalSender struct {
}

func (sg SignalSender) SendPairConfirmationSignal(adress *net.UDPAddr) {
	broadcastAddr, err := net.ResolveUDPAddr("udp4", adress.String())
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp4", nil, broadcastAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte{1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pairing confirmation byte sent!")
}

func (sg SignalSender) CreateConnection() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp4", ":25643")
	if err != nil {
		log.Println("Error resolving address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	log.Println("Listening for broadcast messages on port 25643...")
	return conn
}

func (sg SignalSender) readBuffer(conn *net.UDPConn) (*net.UDPAddr, string) {
	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error reading from UDP:", remoteAddr, err)
		}

		if n > 0 {
			receivedHostname := string(buffer[0:n])
			log.Printf("Received hostname: %s from %s\n", receivedHostname, remoteAddr)
			return remoteAddr, receivedHostname
		}
	}
	return nil, "null"
}

func (sg SignalSender) ListenAndHandle() (*net.UDPAddr, string) {
	conn := sg.CreateConnection()
	return sg.readBuffer(conn)
}
