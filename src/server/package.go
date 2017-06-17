package server

import (
	"net"
	"time"
	"fmt"
)

type Packet struct {
	client_addr *net.UDPAddr
	n int
	buff []byte
}

const(
	REGIST_PACKET = 0x01
	HEARTBEAT_PACKET = 0x02
	CONNECT_PACKET = 0x03
)

func handlePackage(packet *Packet) {
	switch packet.buff[0] {
		case REGIST_PACKET:
			packet.regist()
		case HEARTBEAT_PACKET:
			packet.heartbeat()
		case CONNECT_PACKET:
			packet.connect()
	}
}


func (packet *Packet)regist() {
	uuid := string(packet.buff[1:])
	smodeclients[uuid] = &SModeClient{uuid,packet.client_addr,0,time.Now()}

	fmt.Println("regist",uuid)
}

func (packet *Packet)heartbeat() {
	uuid := string(packet.buff[1:])
	fmt.Println("heartbeat",uuid)
}

func (packet *Packet)connect() {
	uuid := string(packet.buff[1:])
	fmt.Println("connect",uuid)
}