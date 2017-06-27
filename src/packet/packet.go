package packet

import (
	"net"
)

type Packet struct {
	Addr *net.UDPAddr
	buff []byte
}

type BytesPacket interface {
	Bytes() []byte
}

const(
	REGIST_PACKET = 0x01
	CONNECT_PACKET = 0x02
	ERROR_PACKET = 0x03
	OK_PACKET = 0x04
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