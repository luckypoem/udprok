package packet

import (
	"net"
)

type Packet struct {
	addr *net.UDPAddr
	buff []byte
}

const(
	REGIST_PACKET = 0x01
	UUIDUSED_PACKET = 0x02
	REGISTED_PACKET = 0x03
	CONNECT_PACKET = 0x04
)

func NewPacket(addr *net.UDPAddr, data []byte) *Packet {
	return &Packet{addr,data}
}

func (packet *Packet)Type() byte {
	return packet.buff[0]
}

func (packet *Packet)Payload() []byte {
	return packet.buff[1:]
}