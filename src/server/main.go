package server

import (
	"flag"
	"loger"
	"net"
	"strconv"
	"time"
	"errors"
	"fmt"
)

type CModeClient struct {
	uuid string
	addr net.UDPAddr
	status byte
	updatetime time.Time
}

type Packet struct {
	client_addr net.UDPAddr
	n int
	buff []byte
}

const(
	REGIST_PACKET = 0x01

)

func handlePackage(packet Packet) {
	switch packet.buff[0] {
		case REGIST_PACKET:
			handleRegistPacket(&packet)
	}
}

func handleRegistPacket(packet *Packet) {
	uuid := string(packet.buff[1:])
	fmt.Println(uuid)
}

var conn net.UDPConn

func Main() {

	port := flag.Int("p", 7450, "The port to listen UDP")
	addr := flag.String("l", "0.0.0.0", "The address to listen UDP")
	flag.Parse()

	if *port < 1 || *port > 65535 {
		loger.Error(loger.PORT_NOT_IN_RANGE)
	}

	server_addr, err := net.ResolveUDPAddr("udp", *addr + ":" + strconv.Itoa(*port))
	loger.CheckError(err)

	conn, err := net.ListenUDP("udp", server_addr)
	loger.CheckError(err)

	var buf [1024]byte
	for {
		n, client_addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			loger.LogError(err)
			continue
		}
		if n <=0 {
			loger.LogError(errors.New("read udp error."))
			continue
		}
		go handlePackage(Packet{*client_addr,n,buf[0:n]})
	}

}